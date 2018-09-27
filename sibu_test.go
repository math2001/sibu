package sibu

import (
	"fmt"
	"reflect"
	"testing"
)

// just a shortcut for params
type p []interface{}

func resultsEqual(s string, a p, err error, ss string, aa p, haserror bool) string {
	// the double letter argument are the expected values
	if s != ss {
		return fmt.Sprintf("sql requests don't match: %q != %q", s, ss)
	} else if !reflect.DeepEqual(a, aa) {
		return fmt.Sprintf("arguments don't match: %#v != %#v", a, aa)
	} else if haserror == true && err == nil {
		return fmt.Sprintf("expected error, got nil")
	} else if haserror == false && err != nil {
		return fmt.Sprintf("got unexpected error: %s", err)
	}
	return ""
}

func TestNoArgs(t *testing.T) {
	var (
		b   Sibu
		s   string
		a   p
		err error
	)
	b = Sibu{}
	b.Write("SELECT *")
	b.Write("FROM table")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table", nil, false); msg != "" {
		t.Errorf("Unexpected return values #basic: %s", msg)
	}
	b = Sibu{}
	b.Write("SELECT *")
	b.Write("FROM table")
	b.Write("JOIN other")
	b.Write("ON other.a = table.b")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table JOIN other ON other.a = table.b", nil, false); msg != "" {
		t.Errorf("Unexpected return values #join: %s", msg)
	}
}

func TestArgs(t *testing.T) {
	var (
		b   Sibu
		s   string
		a   p
		err error
	)
	b = Sibu{}
	b.Write("SELECT * FROM table")
	b.Add("WHERE userid={{ p }}", 10)
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table WHERE userid=$1", p{10}, false); msg != "" {
		t.Errorf("Unexpected return values #single: %s", msg)
	}
	b = Sibu{}
	b.Write("SELECT * FROM table")
	b.Add("WHERE userid={{ p }} AND tags LIKE {{ p }}", 10, "%go%")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table WHERE userid=$1 AND tags LIKE $2", p{10, "%go%"}, false); msg != "" {
		t.Errorf("Unexpected return values #double: %s", msg)
	}
}
