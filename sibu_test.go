package sibu

import "testing"

func TestNoArgs(t *testing.T) {
	var b Sibu
	b = Sibu{}
	b.Write("SELECT *")
	b.Write("FROM table")
	if s, a, err := b.Parameterize(); s != "SELECT * FROM table" || len(a) != 0 || err != nil {
		t.Errorf("Invalid return values: %q, %v, %s", s, a, err)
	}
	b = Sibu{}
	b.Write("SELECT *")
	b.Write("FROM table")
	b.Write("JOIN other")
	b.Write("ON other.a = table.b")
	if s, a, err := b.Parameterize(); s != "SELECT * FROM table JOIN other ON other.a = table.b" || len(a) != 0 || err != nil {
		t.Errorf("Invalid return values: %q, %v, %s", s, a, err)
	}
}
