package sibu

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/pkg/errors"
)

// Sibu is simplistic sql request buidler
type Sibu struct {
	args   []interface{}
	b      strings.Builder
	pcount int
}

// Unconditionnaly write string to the request.
// It is good practice to make one call per clause
func (s *Sibu) Write(clause string) {
	s.b.WriteString(clause)
}

// Add adds an argument value to the builder. If do is false, does nothing
func (s *Sibu) Add(clause string, value interface{}) {
	s.b.WriteString(clause)
	s.args = append(s.args, value)
}

func (s *Sibu) p() string {
	s.pcount++
	return fmt.Sprintf("$%d", s.pcount)
}

// Parameterize returns the sql query with the need parameter
func (s *Sibu) Parameterize() (string, []interface{}, error) {
	t := template.New("parameterizer")
	t.Funcs(map[string]interface{}{
		"p": s.p,
	})
	req := s.b.String()
	if _, err := t.Parse(req); err != nil {
		return "", nil, errors.Wrapf(err, "failed to parse request %q", req)
	}
	var b bytes.Buffer
	t.Execute(&b, nil)
	return b.String(), s.args, nil
}
