package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const getEntriesTemplate = `-- name: GetEntries :many
SELECT e.ZUUID, e.ZCREATIONDATE, e.ZMODIFIEDDATE, e.ZTIMEZONE, e.ZDURATION,
e.ZISPINNED, e.ZSTARRED, e.ZISALLDAY,
e.ZMARKDOWNTEXT, e.ZEDITINGTIME,
e.ZCREATIONDEVICE, e.ZCREATIONDEVICETYPE, e.ZCREATIONDEVICEMODEL, e.ZCREATIONOSNAME, e.ZCREATIONOSVERSION,
%[1]s AS tags,
e.ZLOCATION, l.ZUSERLABEL, l.ZLATITUDE, l.ZLONGITUDE, l.ZALTITUDE, l.ZPLACENAME, l.ZLOCALITYNAME, l.ZADMINISTRATIVEAREA, l.ZCOUNTRY,
e.ZWEATHER, w.ZCONDITIONSDESCRIPTION, w.ZMOONPHASE, w.ZMOONPHASECODE, w.ZPRESSUREMB, w.ZRELATIVEHUMIDITY, w.ZSUNRISEDATE, w.ZSUNSETDATE, w.ZTEMPERATURECELSIUS, w.ZVISIBILITYKM, w.ZWEATHERCODE, w.ZWEATHERSERVICENAME, w.ZWINDBEARING, w.ZWINDSPEEDKPH
FROM ZENTRY e
JOIN ZJOURNAL j ON (e.ZJOURNAL = j.Z_PK)
%[2]s
LEFT JOIN ZLOCATION l ON (l.Z_PK = e.ZLOCATION)
LEFT JOIN ZWEATHER w ON (w.Z_PK = e.ZWEATHER)
WHERE j.ZNAME = ?
GROUP BY e.ZUUID
`

type GetEntriesRow struct {
	Zuuid                  sql.NullString
	Zcreationdate          sql.NullString
	Zmodifieddate          sql.NullString
	Ztimezone              []byte
	Zduration              sql.NullInt64
	Zispinned              sql.NullInt64
	Zstarred               sql.NullInt64
	Zisallday              sql.NullInt64
	Zmarkdowntext          sql.NullString
	Zeditingtime           sql.NullFloat64
	Zcreationdevice        sql.NullString
	Zcreationdevicetype    sql.NullString
	Zcreationdevicemodel   sql.NullString
	Zcreationosname        sql.NullString
	Zcreationosversion     sql.NullString
	Tags                   sql.NullString
	Zlocation              sql.NullInt64
	Zuserlabel             sql.NullString
	Zlatitude              sql.NullFloat64
	Zlongitude             sql.NullFloat64
	Zaltitude              sql.NullFloat64
	Zplacename             sql.NullString
	Zlocalityname          sql.NullString
	Zadministrativearea    sql.NullString
	Zcountry               sql.NullString
	Zweather               sql.NullInt64
	Zconditionsdescription sql.NullString
	Zmoonphase             sql.NullFloat64
	Zmoonphasecode         sql.NullString
	Zpressuremb            sql.NullFloat64
	Zrelativehumidity      sql.NullFloat64
	Zsunrisedate           sql.NullString
	Zsunsetdate            sql.NullString
	Ztemperaturecelsius    sql.NullFloat64
	Zvisibilitykm          sql.NullFloat64
	Zweathercode           sql.NullString
	Zweatherservicename    sql.NullString
	Zwindbearing           sql.NullFloat64
	Zwindspeedkph          sql.NullFloat64
}

func (q *Queries) GetEntries(ctx context.Context, zname string) ([]GetEntriesRow, error) {
	rows, err := q.db.QueryContext(ctx, q.entriesQuery, toNullString(zname))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEntriesRow
	for rows.Next() {
		var i GetEntriesRow
		if err := rows.Scan(
			&i.Zuuid,
			&i.Zcreationdate,
			&i.Zmodifieddate,
			&i.Ztimezone,
			&i.Zduration,
			&i.Zispinned,
			&i.Zstarred,
			&i.Zisallday,
			&i.Zmarkdowntext,
			&i.Zeditingtime,
			&i.Zcreationdevice,
			&i.Zcreationdevicetype,
			&i.Zcreationdevicemodel,
			&i.Zcreationosname,
			&i.Zcreationosversion,
			&i.Tags,
			&i.Zlocation,
			&i.Zuserlabel,
			&i.Zlatitude,
			&i.Zlongitude,
			&i.Zaltitude,
			&i.Zplacename,
			&i.Zlocalityname,
			&i.Zadministrativearea,
			&i.Zcountry,
			&i.Zweather,
			&i.Zconditionsdescription,
			&i.Zmoonphase,
			&i.Zmoonphasecode,
			&i.Zpressuremb,
			&i.Zrelativehumidity,
			&i.Zsunrisedate,
			&i.Zsunsetdate,
			&i.Ztemperaturecelsius,
			&i.Zvisibilitykm,
			&i.Zweathercode,
			&i.Zweatherservicename,
			&i.Zwindbearing,
			&i.Zwindspeedkph,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPhotos = `-- name: GetPhotos :many
SELECT e.ZUUID, a.ZTYPE, a.ZFILENAME, a.ZIDENTIFIER, a.ZMD5
FROM ZATTACHMENT a
JOIN ZENTRY e ON (e.Z_PK = a.ZENTRY)
`

type GetPhotosRow struct {
	Zuuid       sql.NullString
	Ztype       sql.NullString
	Zfilename   sql.NullString
	Zidentifier sql.NullString
	Zmd5        sql.NullString
}

func (q *Queries) GetPhotos(ctx context.Context) ([]GetPhotosRow, error) {
	rows, err := q.db.QueryContext(ctx, getPhotos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPhotosRow
	for rows.Next() {
		var i GetPhotosRow
		if err := rows.Scan(
			&i.Zuuid,
			&i.Ztype,
			&i.Zfilename,
			&i.Zidentifier,
			&i.Zmd5,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func toNullString(value string) sql.NullString {
	valid := len(strings.TrimSpace(value)) != 0
	return sql.NullString{String: value, Valid: valid}
}

func formatEntriesQuery(join entryTagJoin) string {
	tagsExpr := "NULL"
	joinClause := ""
	if join.table != "" && join.entryColumn != "" && join.tagColumn != "" {
		tagsExpr = "GROUP_CONCAT(t.ZNAME, ',')"
		joinClause = fmt.Sprintf("LEFT JOIN %s et ON (et.%s = e.Z_PK)\nLEFT JOIN ZTAG t ON (t.Z_PK = et.%s)\n",
			quoteIdent(join.table), quoteIdent(join.entryColumn), quoteIdent(join.tagColumn))
	}
	return fmt.Sprintf(getEntriesTemplate, tagsExpr, joinClause)
}
