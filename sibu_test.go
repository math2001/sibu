package sibu

import (
	"fmt"
	"reflect"
	"testing"
)

// just a shortcut for params
func resultsEqual(s string, a Params, err error, ss string, aa Params,
	haserror bool) string {
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
		a   Params
		err error
	)
	b = Sibu{}
	b.Add("SELECT *")
	b.Add("FROM table")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table", nil,
		false); msg != "" {
		t.Errorf("Unexpected return values [basic]: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT *")
	b.Add("FROM table")
	b.Add("JOIN other")
	b.Add("ON other.a = table.b")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err,
		"SELECT * FROM table JOIN other ON other.a = table.b", nil,
		false); msg != "" {
		t.Errorf("Unexpected return values [join]: %s", msg)
	}
}

func fromUserInput(t *testing.T, from string, contains string) *Sibu {
	b := &Sibu{}
	b.Add("SELECT * FROM a JOIN b b ON a.a=b.b")

	where := OpClause{}
	if from != "" {
		where.Add("AND", "a.d={{ p }}", from)
	}
	if contains != "" {
		where.Add("AND", "b.c LIKE {{ p }}", contains)
	}
	b.AddClause("WHERE", where)
	return b
}

func TestArgs(t *testing.T) {
	var (
		b   Sibu
		s   string
		a   Params
		err error
	)
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Add("WHERE userid={{ p }}", 10)
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table WHERE userid=$1",
		Params{10}, false); msg != "" {
		t.Errorf("Unexpected return values [single]: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Add("WHERE userid={{ p }} AND tags LIKE {{ p }}", 10, "%go%")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err,
		"SELECT * FROM table WHERE userid=$1 AND tags LIKE $2",
		Params{10, "%go%"}, false); msg != "" {
		t.Errorf("Unexpected return values [double]: %s", msg)
	}
	s, a, err = fromUserInput(t, "math2001", "go").Query()
	if msg := resultsEqual(s, a, err,
		"SELECT * FROM a JOIN b b ON a.a=b.b WHERE a.d=$1 AND b.c LIKE $2",
		Params{"math2001", "go"},
		false); msg != "" {
		t.Errorf("Unexpected return values [OpClause]: %s", msg)
	}
}

func TestBareAdd(t *testing.T) {
	var (
		b   Sibu
		s   string
		a   Params
		err error
	)
	b = Sibu{}
	b.BareWrite("SELECT * FR")
	b.BareWrite("OM table")
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "SELECT * FROM table", nil, false); msg != "" {
		t.Errorf("Unexpected return values: %s", msg)
	}
}

func TestErrors(t *testing.T) {
	var (
		b   Sibu
		s   string
		a   Params
		err error
	)
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Add("WHERE userid={{ p }} AND extra={{ p }}", 10)
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "", nil, true); msg != "" {
		t.Errorf("Unexpected return values [extra parameter]: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Add("WHERE userid={{ p }", 10)
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "", nil, true); msg != "" {
		t.Errorf("Unexpected return values [invalid tmpl syntax]: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Add("WHERE userid={{ p }} AND error={{ .Error }}", 10)
	s, a, err = b.Query()
	if msg := resultsEqual(s, a, err, "", nil, true); msg != "" {
		t.Errorf("Unexpected return values [invalid data use in tmpl]: %s", msg)
	}
}
