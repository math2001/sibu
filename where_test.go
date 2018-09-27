package sibu

import (
	"fmt"
	"reflect"
	"testing"
)

func resultsEquals(req string, args Params, rreq string, aargs Params) string {
	if req != rreq {
		return fmt.Sprintf("req didn't match: %q != %q", req, rreq)
	} else if !reflect.DeepEqual(args, aargs) {
		return fmt.Sprintf("args didn't match: %#v != %#v", args, aargs)
	}
	return ""
}

func TestLinear(t *testing.T) {
	var (
		w    Where
		b    Sibu
		req  string
		args Params
	)
	w = Where{}
	w.And("u.userid={{ p }}", 10)
	w.And("u.followingid={{ p }}", 20)
	req, args = w.GetClause()
	if msg := resultsEquals(req, args, "WHERE u.userid={{ p }} AND u.followingid={{ p }}", Params{10, 20}); msg != "" {
		t.Fatalf("Invalid result: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Extend(&w)

	w = Where{}
	w.Or("u.userid={{ p }}", 10)
	w.Or("u.followingid={{ p }}", 20)
	req, args = w.GetClause()
	if msg := resultsEquals(req, args, "WHERE u.userid={{ p }} OR u.followingid={{ p }}", Params{10, 20}); msg != "" {
		t.Fatalf("Invalid result: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Extend(&w)
}
