package wiz

import (
	"fmt"
	"github.com/pkg/errors"
)

//Antipanic - usage: defer Antipanic(&e, "label") at the beginning of any function.
// Label is usually the name of the function prefixed by its package name, e.g.
// "package.func". This results in traditionally wrapped errors (mimics errors.Wrap).
// Main purpose: deferring this function transforms panics into errors when they
// occur inside the function it was deferred in, and requires only one line.
// Additionally it allows gung-ho use of panics inside parent function to replace
// error checking, which can shorten code somewhat (as panics force a return).
func Antipanic(e *error, label string) {
	if e == nil {
		irony := errors.New(label + ": Antipanic: do not use nil error receiver.")
		panic(irony)
	}
	//If called using defer, will detect any active panics
	if pan := recover(); pan != nil {
		//A panic was caught above.
		s := fmt.Sprintf("%v", pan) //Get it as string
		*e = errors.New(s)
	}
	*e = errors.Wrap(*e, label)
}

//Panics with a string message
func Panic(s string) {
	err := errors.New(s)
	Check(err)
}

//Panics if an error is non-nil. Use after defer Antipanic to shorten code.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

//Panics if a condition is false. The "label" string is the message of the panic.
func Assert(condition bool, label string) {
	if condition == false {
		Panic("Assertion failed: " + label)
	}
}
