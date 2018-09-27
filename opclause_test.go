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
		o    OpClause
		b    Sibu
		req  string
		args Params
	)
	o = OpClause{}
	o.Add("AND", "u.userid={{ p }}", 10)
	o.Add("OR", "u.followingid={{ p }}", 20)
	req, args = o.GetOpClause()
	if msg := resultsEquals(req, args, "u.userid={{ p }} OR u.followingid={{ p }}", Params{10, 20}); msg != "" {
		t.Fatalf("Invalid result: %s", msg)
	}
	b = Sibu{}
	b.Add("SELECT * FROM table")
	b.Extend("WHERE", &o)
}

// func TestGroups(t *testing.T) {
// 	var (
// 		w    Where
// 		b    Sibu
// 		req  string
// 		args Params
// 	)
// 	w = Where{}
// 	w.And("a.b={{ p }}", 10)
// 	w.Open()
// 	w.Or("c.d={{ p }}", 15)
// 	w.Or("e.f={{ p }}", 20)
// 	w.Close()
// 	req, args = w.GetClause()
// 	if msg := resultsEquals(req, args, "WHERE a.b=$1 ()"); msg != "" {

// 	}
// }
