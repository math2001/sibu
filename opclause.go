package sibu

// opclause defines Operator Clauses util functions. One could very easily
// implement an other utility, so long as they satisfy the OpClauser interface

import "strings"

type condition struct {
	op   string
	req  string
	args []interface{}
}

// OpClause is used to build composed the WHERE kind of clauses dynamically
type OpClause []condition

// Add adds a condition preceded by the op operator (if there is any condition
// before)
func (w *OpClause) Add(op, req string, args ...interface{}) {
	*w = append(*w, condition{
		op:   op,
		req:  req,
		args: args,
	})
}

// Open opens a bracket
func (w *OpClause) Open(op string) {
	*w = append(*w, condition{
		op:   op,
		req:  "(",
		args: nil,
	})
}

// Close closes a bracket
func (w *OpClause) Close() {
	*w = append(*w, condition{
		op:   "",
		req:  ")",
		args: nil,
	})
}

// GetOpClause returns the formatted clause, ready to be plugged in by Extend
func (w *OpClause) GetOpClause() (string, Params) {
	var b strings.Builder
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
