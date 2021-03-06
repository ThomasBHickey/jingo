// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"fmt"
)

type AData interface{}

// J source has additional fields:
// k: offset of ravel
// RefCount : reference count
// m: # bytes in ravel

type A struct {
	Type   AType // defined in jtype
	Flag   AFLAG // defined in jtype
	Length int
	Shape  []int
	Data   AData
}

func shape2length(shape []int) (length int) {
	length = 1
	for _, sp := range shape {
		length *= sp
	}
	return
}
func NewArray(typ AType, shape []int, adata AData) (a A) {
	a.Type = typ
	a.Length = shape2length(shape)
	a.Shape = shape
	a.Data = adata
	return
}

func NewVerbArray(vd VAData) A {
	a := NewArray(VERB, []int{}, vd)
	return a
}
func NewASGNArray(vd VAData) A {
	a := NewArray(ASGN, []int{}, vd)
	return a
}

func NewIntArray(shape []int) A {
	a := NewArray(INT, shape, make([]int, shape2length(shape)))
	return a
}
func NewSIntArray(i int) A {
	a := NewIntArray([]int{})
	a.Data = []int{i}
	return a
}

type NameData struct {
	name  string
	value A
}

func NewNameArray(nm string) A {
	return NewArray(NAME, []int{1}, NameData{name: nm})
}
func NewByteArray(shape []int) (a A) {
	a = NewArray(LIT, shape, make([]byte, shape2length(shape)))
	return a
}
func (array A) ShowArray() {
	switch array.Type {
	case INT:
		fmt.Println("Found INT array")
		fmt.Println(array)
		fmt.Println("Shape", array.Shape)
		fmt.Println("Array length", len(array.Data.([]int)))
	case LIT:
		fmt.Println("Found LIT array")
		fmt.Println("Array length", len(array.Data.([]byte)))
	}
}
