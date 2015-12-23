// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"errors"
	"fmt"
)

func vdyad(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vdyad", b, e, stack)
	return z, errors.New("vdyad undefined")
}
func vmonad(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vmonad (not implemented)", b, e, stack)
	return
}
func vadv(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vadv (not implemented)", b, e, stack)
	return
}

func vconj(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vconj (not implemented)", b, e, stack)
	return
}
func vfork(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vfork (not implemented", b, e, stack)
	return
}
func vis(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vis (not implemented)", b, e, stack)
	return
}
func vpunc(jt J, b, e int, stack []A) (z A, err error) {
	fmt.Println("In vpunc (not implemented)", b, e, stack)
	return
}
