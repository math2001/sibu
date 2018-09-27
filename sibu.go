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

// Write a clause to the request, with an automatic space if there is some text
// before.
func (s *Sibu) Write(clause string) {
	if s.b.Len() > 0 {
		s.b.WriteRune(' ')
	}
	s.b.WriteString(clause)
}

// BareWrite exactly writes s to the request. No space is added, so be careful
func (s *Sibu) BareWrite(str string) {
	s.b.WriteString(str)
}

// Add adds an argument value to the builder.
func (s *Sibu) Add(clause string, value interface{}) {
	s.Write(clause)
	s.args = append(s.args, value)
}

func (s *Sibu) p() string {
	s.pcount++
	return fmt.Sprintf("$%d", s.pcount)
}

// Query returns the sql query with the need parameter
func (s *Sibu) Query() (string, []interface{}, error) {
	t := template.New("parameterizer")
	t.Funcs(map[string]interface{}{
		"p": s.p,
	})
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
