// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"log"

	. "github.com/gocircuit/escher/a"
	. "github.com/gocircuit/escher/circuit"
)

func ParseVerb(src string) (verb Verb) {
	defer func() {
		if r := recover(); r != nil {
			verb = Verb{}
		}
	}()
	t := NewSrcString(src)
	verb = Verb(SeeVerb(t).(Circuit))
	if t.Len() != 0 {
		log.Printf("Non-address characters at end of %q", src)
		panic(1)
	}
	return verb
}

func Parse(src string) (Name, Value) {
	return SeePeer(NewSrcString(src))
}

func ParseCircuit(src string) Circuit {
	n, v := Parse(src)
	if _, ok := n.(Nameless); !ok {
		panic("not a circuit")
	}
	return v.(Circuit)
}
