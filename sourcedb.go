package dayone2md

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite" // import sqlite database library

	"kwo.dev/dayone2md/database"
)

const (
	dbFlavorSqlite = "sqlite"
	macosEpoch     = 978307200 * time.Second // 2001-01-01T00:00:00Z, CFAbsoluteTime
)

func newDBSource(dbFile string) *dbSource {
	return &dbSource{
		dbFile:    dbFile,
		dayoneDir: filepath.Dir(dbFile),
	}
}

type dbSource struct {
	dayoneDir string
	dbFile    string // full filename
}

func (z *dbSource) GetJournal(name string) (*Journal, error) {
	dsn := fmt.Sprintf("file:%s?mode=ro", z.dbFile)
	db, err := sql.Open(dbFlavorSqlite, dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot open database (%s): %w", z.dbFile, err)
	}

	queries := database.New(db)

	entries, err := loadEntries(queries, name)
	if err != nil {
		return nil, err
	}

	if err := loadPhotos(queries, entries); err != nil {
		return nil, err
	}

	var e []Entry
	for _, entry := range entries {
		e = append(e, *entry)
	}

	j := &Journal{
		Entries: e,
	}
	return j, nil
}

func (z *dbSource) GetPhoto(filename string) ([]byte, error) {
	fullname := filepath.Join(z.dayoneDir, "DayOnePhotos", filename)
	return os.ReadFile(fullname)
}

// loadEntries return the entries mapped by UUID
func loadEntries(q *database.Queries, journalName string) (map[string]*Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rs, err := q.GetEntries(ctx, toNullString(journalName))
	if err != nil {
		return nil, fmt.Errorf("cannot query entries: %w", err)
	}

	entries := make(map[string]*Entry)
	for _, row := range rs {
		e := toEntry(row)
		entries[e.UUID] = e
	}

	return entries, nil
}

func loadPhotos(q *database.Queries, entries map[string]*Entry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rs, err := q.GetPhotos(ctx)
	if err != nil {
		return fmt.Errorf("cannot query photos: %w", err)
	}

	for _, row := range rs {
		p := toPhoto(row)
		if e, ok := entries[row.Zuuid.String]; ok {
			e.Photos = append(e.Photos, *p)
		}
	}

	return nil
}

func toEntry(row database.GetEntriesRow) *Entry {
	e := &Entry{
		UUID:                row.Zuuid.String,
		CreationDate:        parseDate(row.Zcreationdate),
		ModifiedDate:        parseDate(row.Zmodifieddate),
		TimeZone:            extractTimezoneName(row.Ztimezone),
		Duration:            int(row.Zduration.Int64), // seconds?
		Pinned:              row.Zispinned.Int64 != 0,
		Starred:             row.Zstarred.Int64 != 0,
		AllDay:              row.Zisallday.Int64 != 0,
		Text:                row.Zmarkdowntext.String,
		EditingTime:         row.Zeditingtime.Float64,
		CreationDevice:      row.Zcreationdevice.String,
		CreationDeviceType:  row.Zcreationdevicetype.String,
		CreationDeviceModel: row.Zcreationdevicemodel.String,
		CreationOSName:      row.Zcreationosname.String,
		CreationOSVersion:   row.Zcreationosversion.String,
	}
	if row.Tags.Valid {
		e.Tags = filterEmpty(strings.Split(row.Tags.String, ","))
	}
	if row.Zlocation.Valid {
		e.Location = &Location{
			Label:     row.Zuserlabel.String,
			Address:   row.Zplacename.String,
			City:      row.Zlocalityname.String,
			State:     row.Zadministrativearea.String,
			Country:   row.Zcountry.String,
			Latitude:  row.Zlatitude.Float64,
			Longitude: row.Zlongitude.Float64,
			Altitude:  row.Zaltitude.Float64,
		}
	}
	if row.Zweather.Valid {
		e.Weather = &Weather{
			Conditions:         row.Zconditionsdescription.String,
			MoonPhase:          row.Zmoonphase.Float64,
			MoonPhaseCode:      row.Zmoonphasecode.String,
			PressureMB:         row.Zpressuremb.Float64,
			RelativeHumidity:   int(row.Zrelativehumidity.Float64),
			SunriseDate:        parseDate(row.Zsunrisedate),
			SunsetDate:         parseDate(row.Zsunsetdate),
			TemperatureCelsius: row.Ztemperaturecelsius.Float64,
			VisibilityKM:       row.Zvisibilitykm.Float64,
			WeatherCode:        row.Zweathercode.String,
			WeatherServiceName: row.Zweatherservicename.String,
			WindBearing:        int(row.Zwindbearing.Float64),
			WindSpeedKPH:       row.Zwindspeedkph.Float64,
		}
	}
	return e
}

func toPhoto(row database.GetPhotosRow) *Photo {
	return &Photo{
		Type:       row.Ztype.String,
		Identifier: row.Zidentifier.String,
		Filename:   row.Zfilename.String,
		MD5:        row.Zmd5.String,
	}
}

func parseDate(value sql.NullString) time.Time {
	if !value.Valid || value.String == "" {
		return time.Time{}
	}
	// seconds since unix epoch
	// decimal is fractionsal seconds which can be dropped
	cd1, err := strconv.ParseFloat(value.String, 64)
	if err == nil {
		return time.Unix(int64(cd1), 0).Add(macosEpoch).UTC()
	}
	cd2, err := time.Parse(time.RFC3339, value.String)
	if err == nil {
		return cd2.Add(macosEpoch).UTC()
	}
	return time.Time{}
}

func extractTimezoneName(data []byte) string {
	subslices1 := bytes.SplitAfterN(data, []byte{36, 99, 108, 97, 115, 115}, 2) // split after $class
	if len(subslices1) != 2 {
		return ""
	}
	subslices2 := bytes.SplitN(subslices1[1], []byte{84, 90, 105, 102, 50}, 2) // split on TZif2
	if len(subslices2) != 2 {
		return ""
	}
	// remove cruft before first capital ASCII letter
	firstLetter := bytes.IndexAny(subslices2[0], "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if firstLetter == -1 {
		return ""
	}
	subslice3 := subslices2[0][firstLetter:]
	// remove last 4 bytes
	result := subslice3[0 : len(subslice3)-4]
	return string(result)
}

func filterEmpty(values []string) []string {
	var z []string
	for _, x := range values {
		if t := strings.TrimSpace(x); t != "" {
			z = append(z, t)
		}
	}
	return z
}

func toNullString(value string) sql.NullString {
	valid := len(strings.TrimSpace(value)) != 0
	return sql.NullString{String: value, Valid: valid}
}
