// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package text

import (
	"bytes"
	"text/template"

	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

// Template …
type Template struct{}

func (Template) Spark(*be.Eye, *be.Matter, ...interface{}) Value {
	return nil
}

func (Template) CognizeIn(eye *be.Eye, v interface{}) {
	td := v.(Circuit)
	t, err := template.New("").Parse(td.StringAt("Template"))
	if err != nil {
		panic(err)
	}
	var w bytes.Buffer
	if err = t.Execute(&w, td.CircuitAt("Data")); err != nil {
		panic(err)
	}
	eye.Show("Out", w.String())
}

func (Template) CognizeOut(eye *be.Eye, v interface{}) {}

func gateHierarchy(u Circuit) map[string]interface{} { // not used
	r := make(map[string]interface{})
	for g, v := range u.Gates() {
		switch t := v.(type) {
		case Circuit:
			r[g.(string)] = gateHierarchy(t)
		default:
			r[g.(string)] = v
		}
	}
	return r
}