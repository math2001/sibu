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
		req  string
		args Params
	)
	o = OpClause{}
	o.Add("AND", "u.userid={{ p }}", 10)
	o.Add("OR", "u.followingid={{ p }}", 20)
	req, args = o.GetOpClause()
	if msg := resultsEquals(req, args, "u.userid={{ p }} OR u.followingid={{ p }}", Params{10, 20}); msg != "" {
		t.Errorf("Invalid result: %s", msg)
	}
}

func TestGroups(t *testing.T) {
	var (
		o    OpClause
		req  string
		args Params
	)
	o = OpClause{}
	o.Add("AND", "a.b={{ p }}", 10)
	o.Open("AND")
	o.Add("OR", "c.d={{ p }}", 15)
	o.Add("OR", "e.f={{ p }}", 20)
	o.Close()
	req, args = o.GetOpClause()
	if msg := resultsEquals(req, args, "a.b={{ p }} AND ( c.d={{ p }} OR e.f={{ p }} )", Params{10, 15, 20}); msg != "" {
		t.Errorf("Invalid result: %s", msg)
	}
}
