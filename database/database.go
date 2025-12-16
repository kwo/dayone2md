package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type DBTX interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func New(ctx context.Context, db DBTX) (*Queries, error) {
	join, err := detectEntryTagJoin(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("detect entry/tag relationship: %w", err)
	}
	return &Queries{
		db:           db,
		entriesQuery: formatEntriesQuery(join),
	}, nil
}

type Queries struct {
	db           DBTX
	entriesQuery string
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:           tx,
		entriesQuery: q.entriesQuery,
	}
}

type entryTagJoin struct {
	table       string
	entryColumn string
	tagColumn   string
}

func detectEntryTagJoin(ctx context.Context, db DBTX) (entryTagJoin, error) {
	rows, err := db.QueryContext(ctx, `SELECT name FROM sqlite_master WHERE type='table' AND name LIKE 'Z_%TAGS'`)
	if err != nil {
		return entryTagJoin{}, err
	}
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	var tableNames []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return entryTagJoin{}, err
		}
		tableNames = append(tableNames, tableName)
	}
	if err := rows.Err(); err != nil {
		return entryTagJoin{}, err
	}
	if err := rows.Close(); err != nil {
		return entryTagJoin{}, err
	}
	rows = nil

	for _, tableName := range tableNames {
		join, err := inspectEntryTagJoin(ctx, db, tableName)
		if err != nil {
			return entryTagJoin{}, err
		}
		if join.table != "" && join.entryColumn != "" && join.tagColumn != "" {
			return join, nil
		}
	}

	return entryTagJoin{}, nil
}

func inspectEntryTagJoin(ctx context.Context, db DBTX, tableName string) (entryTagJoin, error) {
	join := entryTagJoin{}

	fkStmt := fmt.Sprintf("PRAGMA foreign_key_list(%s)", quoteIdent(tableName))
	fkRows, err := db.QueryContext(ctx, fkStmt)
	if err != nil {
		return join, err
	}
	defer fkRows.Close()

	for fkRows.Next() {
		var (
			id, seq                   int64
			parentTable, from, to     string
			onUpdate, onDelete, match string
		)
		if err := fkRows.Scan(&id, &seq, &parentTable, &from, &to, &onUpdate, &onDelete, &match); err != nil {
			return join, err
		}
		switch parentTable {
		case "ZENTRY":
			if join.entryColumn == "" {
				join.entryColumn = from
			}
		case "ZTAG":
			if join.tagColumn == "" {
				join.tagColumn = from
			}
		}
	}
	if err := fkRows.Err(); err != nil {
		return join, err
	}

	if join.entryColumn != "" && join.tagColumn != "" {
		join.table = tableName
		return join, nil
	}

	infoStmt := fmt.Sprintf("PRAGMA table_info(%s)", quoteIdent(tableName))
	infoRows, err := db.QueryContext(ctx, infoStmt)
	if err != nil {
		return join, err
	}
	defer infoRows.Close()

	for infoRows.Next() {
		var (
			cid     int64
			name    string
			colType string
			notNull int64
			dflt    sql.NullString
			pk      int64
		)
		if err := infoRows.Scan(&cid, &name, &colType, &notNull, &dflt, &pk); err != nil {
			return join, err
		}
		if join.entryColumn == "" && strings.Contains(strings.ToUpper(name), "ENTR") {
			join.entryColumn = name
		}
		if join.tagColumn == "" && strings.Contains(strings.ToUpper(name), "TAG") {
			join.tagColumn = name
		}
	}
	if err := infoRows.Err(); err != nil {
		return join, err
	}

	if join.entryColumn != "" && join.tagColumn != "" {
		join.table = tableName
		return join, nil
	}

	return entryTagJoin{}, nil
}

func quoteIdent(name string) string {
	escaped := strings.ReplaceAll(name, `"`, `""`)
	return `"` + escaped + `"`
}
