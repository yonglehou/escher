// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	// "fmt"

	"github.com/gocircuit/escher/kit/plumb"
	. "github.com/gocircuit/escher/circuit"
	. "github.com/gocircuit/escher/be"
)

func MaterializeUnion(name string, field ...string) Reflex {
	reflex, eye := plumb.NewEye(append(field, name)...)
	go func() {
		conj := New()
		for {
			dvalve, dvalue := eye.See()
			if dvalve == name { // conjunction updated
				y := make(chan struct{}) // block and
				for _, f_ := range field { // send updated conjunction to all field valves
					f := f_
					go func() {
						eye.Show(f, dvalue.(Circuit).AtNil(f))
						y <- struct{}{}
					}()
				}
				for _ = range field {
					<-y
				}
			} else { // field updated
				conj.Abandon(dvalve).Grow(dvalve, dvalue)
				if conj.Len() == len(field) {
					eye.Show(name, conj)
				}
			}
		}
	}()
	return reflex
}
