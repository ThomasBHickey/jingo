// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package array

import (
	"fmt"
	"github.com/ThomasBHickey/jingo"
)

type Vector interface{}
type Array struct {
	Type     byte
	RefCount int
	length   int
	Shape    []int
	Data     Vector
}

func  newArray(shape []int) (a Array){
	a.RefCount = 1
	a.length = 1
	for _, sp := range(shape) { a.length *= sp }
	a.Shape = shape
	return
}

func NewIntArray(shape []int) (a Array) {
	a = newArray(shape)
	a.Type = jingo.INT
	a.Data = make([]int64, a.length)
	return
}

func NewByteArray(shape []int)(a Array){
	a = newArray(shape)
	a.Type = jingo.LIT
	a.Data = make([]byte, a.length)
	return
}
func (array Array) ShowArray() {
	switch array.Type {
	case jingo.INT:
		fmt.Println("Found INT array")
		fmt.Println(array)
		fmt.Println("Shape", array.Shape)
		fmt.Println("Array length", len(array.Data.([]int64)))
	case jingo.LIT:
		fmt.Println("Found LIT array")
		fmt.Println("Array length", len(array.Data.([]byte)))
	}
}
