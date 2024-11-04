package oui

import (
	"os"
	"ouidb"
	"ouisql"
	"path/filepath"
	"testing"
)

const OUI_DB = "testdata/oui_db"
const OUI_SQLITE = "testdata/oui_sqlite"

func TestMain(m *testing.M) {
	_, err := os.Stat(OUI_DB)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果 OUIDB 不存在, 就生成它
			os.MkdirAll(filepath.Dir(OUI_DB), 0600)
			f, err := os.OpenFile(OUI_DB, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			err = ouidb.Generate(f)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	_, err = os.Stat(OUI_SQLITE)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果 OUI_SQLITE 不存在, 就生成它
			os.MkdirAll(filepath.Dir(OUI_SQLITE), 0600)
			err = ouisql.Generate(OUI_SQLITE)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	os.Exit(m.Run())
}

func TestOuidbLookup(t *testing.T) {
	f, err := os.Open(OUI_DB)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	db, err := ouidb.NewDB(f)
	if err != nil {
		t.Fatal(err)
	}
	if company, err := db.Lookup("AC319D"); err != nil {
		t.Fatal(err)
	} else if company != "Shenzhen TG-NET Botone Technology Co.,Ltd." {
		t.Fatal(company)
	}

	if company, err := db.Lookup("ABCDEF"); err != nil {
		t.Fatal(err)
	} else if company != "" {
		t.Fatal(company)
	}

	if company, err := db.Lookup("004023"); err != nil {
		t.Fatal(err)
	} else if company != "LOGIC CORPORATION" {
		t.Fatal(company)
	}
}

func TestOuisqlLookup(t *testing.T) {
	db, err := ouisql.NewDB(OUI_SQLITE)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if company, err := db.Lookup("AC319D"); err != nil {
		t.Fatal(err)
	} else if company != "Shenzhen TG-NET Botone Technology Co.,Ltd." {
		t.Fatal(company)
	}

	if company, err := db.Lookup("ABCDEF"); err != nil {
		t.Fatal(err)
	} else if company != "" {
		t.Fatal(company)
	}

	if company, err := db.Lookup("004023"); err != nil {
		t.Fatal(err)
	} else if company != "LOGIC CORPORATION" {
		t.Fatal(company)
	}
}

func BenchmarkOuidbLookup(b *testing.B) {
	f, err := os.Open(OUI_DB)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	db, err := ouidb.NewDB(f)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		db.Lookup("AC319D")
	}
}

func BenchmarkOuisqlLookup(b *testing.B) {
	db, err := ouisql.NewDB(OUI_SQLITE)
	if err != nil {
		b.Fatal(err)
	}

	for range b.N {
		db.Lookup("AC319D")
	}
}
