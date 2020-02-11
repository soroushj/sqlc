package sqlite

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParse(t *testing.T) {
	r, err := os.Open(filepath.Join("testdata", "hello.sql"))
	if err != nil {
		t.Fatal(err)
	}
	stmts, err := Parse(r)
	spew.Dump(stmts)
	t.Log(err)
}
