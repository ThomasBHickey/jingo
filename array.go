// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"fmt"
)

type Vector interface{}
type A struct {
	Type     AType
	RefCount int
	Length   int
	Shape    []int
	Data     Vector
}

func newArray(shape []int) (a A) {
	a.RefCount = 1
	a.Length = 1
	for _, sp := range shape {
		a.Length *= sp
	}
	a.Shape = shape
	return
}

func NewIntArray(shape []int) (a A) {
	a = newArray(shape)
	a.Type = INT
	a.Data = make([]int64, a.Length)
	return
}

func NewByteArray(shape []int) (a A) {
	a = newArray(shape)
	a.Type = LIT
	a.Data = make([]byte, a.Length)
	return
}
func (array A) ShowArray() {
	switch array.Type {
	case INT:
		fmt.Println("Found INT array")
		fmt.Println(array)
		fmt.Println("Shape", array.Shape)
		fmt.Println("Array length", len(array.Data.([]int64)))
	case LIT:
		fmt.Println("Found LIT array")
		fmt.Println("Array length", len(array.Data.([]byte)))
	}
}
