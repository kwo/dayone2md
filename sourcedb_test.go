package dayone2md

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGetJournal(t *testing.T) {
	t.Skip()
	src := newDBSource("/Users/karl/Library/Group Containers/5U8NS4GX82.dayoneapp2/Data/Documents/DayOne.sqlite")

	j, err := src.GetJournal(context.Background(), "Journal")
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range j.Entries {
		t.Logf("entry: %f, %s, %s, %v", e.EditingTime, e.UUID, e.TimeZone, e.Tags)
	}
}

func TestTimeZone(tt *testing.T) {
	testCases := []struct {
		tz      string
		datfile string
	}{
		{tz: "Europe/Amsterdam", datfile: "tz-ams.bin"},
		{tz: "Europe/Berlin", datfile: "tz-ber.bin"},
	}

	for _, tCase := range testCases {
		tt.Run(tCase.tz, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", tCase.datfile))
			if err != nil {
				t.Fatal(err)
			}
			if got, want := extractTimezoneName(data), tCase.tz; got != want {
				t.Errorf("bad timezone: %s, expected: %s", got, want)
			}
		})
	}
}

func TestDbReadonly(t *testing.T) {
	dbFile, err := os.CreateTemp("", "dayone2md-*.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(dbFile.Name())

	dsn := fmt.Sprintf("file:%s", dbFile.Name())
	db, err := sql.Open(dbFlavorSqlite, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE ztest (id number, desc text)"); err != nil {
		t.Fatal(err)
	}
	if rs, err := db.Exec("INSERT into ztest (id, desc) VALUES (1, 'apple')"); err != nil {
		t.Fatal(err)
	} else {
		rows, err := rs.RowsAffected()
		if err != nil {
			t.Fatal(err)
		}
		if rows != 1 {
			t.Fatal("cannot insert into ztest table")
		}
	}

	dsnro := fmt.Sprintf("file:%s?mode=ro", dbFile.Name())
	dbro, err := sql.Open(dbFlavorSqlite, dsnro)
	if err != nil {
		t.Fatal(err)
	}
	defer dbro.Close()

	if _, err := dbro.Exec("UPDATE ztest SET desc = 'orange' WHERE id=1"); err == nil {
		t.Fatal("no error when updating read-only database")
	}
}
