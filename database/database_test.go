package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

func TestDetectEntryTagJoinWithForeignKeys(t *testing.T) {
	db := newTestDB(t, "entry_tag_fk")
	t.Cleanup(func() { _ = db.Close() })

	mustExec(t, db, `CREATE TABLE ZENTRY (Z_PK INTEGER PRIMARY KEY)`)
	mustExec(t, db, `CREATE TABLE ZTAG (Z_PK INTEGER PRIMARY KEY)`)
	mustExec(t, db, `
CREATE TABLE Z_42TAGS (
	Z_42ENTRIES INTEGER,
	Z_55TAGS1 INTEGER,
	FOREIGN KEY(Z_42ENTRIES) REFERENCES ZENTRY(Z_PK),
	FOREIGN KEY(Z_55TAGS1) REFERENCES ZTAG(Z_PK)
)`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	join, err := detectEntryTagJoin(ctx, db)
	if err != nil {
		t.Fatalf("detect entry/tag join: %v", err)
	}
	if join.table != "Z_42TAGS" {
		t.Fatalf("unexpected table detected: %s", join.table)
	}
	if join.entryColumn != "Z_42ENTRIES" || join.tagColumn != "Z_55TAGS1" {
		t.Fatalf("unexpected column mapping: %+v", join)
	}
}

func TestDetectEntryTagJoinWithoutForeignKeys(t *testing.T) {
	db := newTestDB(t, "entry_tag_no_fk")
	t.Cleanup(func() { _ = db.Close() })

	mustExec(t, db, `CREATE TABLE Z_90TAGS (Z_90ENTRIES INTEGER, Z_77TAGS1 INTEGER)`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	join, err := detectEntryTagJoin(ctx, db)
	if err != nil {
		t.Fatalf("detect entry/tag join: %v", err)
	}
	if join.table != "Z_90TAGS" {
		t.Fatalf("unexpected table detected: %s", join.table)
	}
	if join.entryColumn != "Z_90ENTRIES" || join.tagColumn != "Z_77TAGS1" {
		t.Fatalf("unexpected column mapping: %+v", join)
	}
}

func TestFormatEntriesQueryWithoutJoin(t *testing.T) {
	query := formatEntriesQuery(entryTagJoin{})
	if !strings.Contains(query, "NULL AS tags") {
		t.Fatalf("expected tags to default to NULL, got: %s", query)
	}
	if strings.Contains(query, "LEFT JOIN ZTAG t") {
		t.Fatalf("did not expect tag join in query: %s", query)
	}
}

func newTestDB(t *testing.T, name string) *sql.DB {
	t.Helper()
	dsn := fmt.Sprintf("file:%s?mode=memory", name)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	db.SetMaxOpenConns(1)
	return db
}

func mustExec(t *testing.T, db *sql.DB, stmt string) {
	t.Helper()
	if _, err := db.Exec(stmt); err != nil {
		t.Fatalf("exec %q: %v", stmt, err)
	}
}
