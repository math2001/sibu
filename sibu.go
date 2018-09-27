package sibu

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/pkg/errors"
)

// ErrDismatchingParam is returned when the number of parsed argument differs
// from the number of given arguments
type ErrDismatchingParam struct {
	Given, Parsed int
}

func (e ErrDismatchingParam) Error() string {
	return fmt.Sprintf("Dismatching parameter count. Given %d, found %d", e.Given, e.Parsed)
}

// Sibu is simplistic sql request buidler
type Sibu struct {
	args   []interface{}
	b      strings.Builder
	pcount int
}

// BareWrite exactly writes s to the request. No space is added, so be careful
func (s *Sibu) BareWrite(str string) {
	s.b.WriteString(str)
}

// Add writes the clause to the request and takes the parameters used, in the
// right order
func (s *Sibu) Add(clause string, value ...interface{}) {
	if s.b.Len() > 0 {
		s.b.WriteRune(' ')
	}
	s.b.WriteString(clause)
	s.args = append(s.args, value...)
}

func (s *Sibu) p() string {
	s.pcount++
	return fmt.Sprintf("$%d", s.pcount)
}

// Query returns the sql query with the need parameter
func (s *Sibu) Query() (string, []interface{}, error) {
	t := template.New("parameterizer").
		Funcs(map[string]interface{}{
			"p": s.p,
		}).
		Option("missingkey=error")
	req := s.b.String()
	if _, err := t.Parse(req); err != nil {
		return "", nil, errors.Wrapf(err, "failed to parse request %q", req)
	}
	var b bytes.Buffer
	if err := t.Execute(&b, nil); err != nil {
		return "", nil, errors.Wrapf(err, "failed to execute template")
	}
	if len(s.args) != s.pcount {
		return "", nil, ErrDismatchingParam{
			Given:  len(s.args),
			Parsed: s.pcount,
		}
	}
	return b.String(), s.args, nil
}
