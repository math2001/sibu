package sibu

import "strings"

type condition struct {
	op   string
	req  string
	args []interface{}
}

// Where is used to build composed WHERE clauses dynamically
type Where []condition

// And adds a condition preceded by the AND operator (if there is any condition
// before)
func (w *Where) And(req string, args ...interface{}) {
	*w = append(*w, condition{
		op:   "AND",
		req:  req,
		args: args,
	})
}

// Or adds a condition preceded by the OR operator (if there is any condition
// before)
func (w *Where) Or(req string, args ...interface{}) {
	*w = append(*w, condition{
		op:   "OR",
		req:  req,
		args: args,
	})
}

// GetClause returns
func (w *Where) GetClause() (string, Params) {
	var b strings.Builder
	b.WriteString("WHERE ")
	// the capacity value is a guess. Most of the time, each condition takes 1
	// parameter.
	var args = make(Params, 0, len(*w))
	for i, cond := range *w {
		if i > 0 && i < len(*w) {
			b.WriteString(" ")
			b.WriteString(cond.op)
			b.WriteString(" ")
		}
		b.WriteString(cond.req)
		args = append(args, cond.args...)
	}
	return b.String(), args
}
