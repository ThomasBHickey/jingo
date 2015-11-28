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
	Shape    []int
	Data     Vector
}

func NewIntArray(shape []int) (a Array) {
	a.Type = jingo.INT
	a.RefCount = 1
	length := 1
	for _, sp := range(shape) { length *= sp }
	a.Shape = shape
	a.Data = make([]int, length)
	return
}
func (array Array) ShowArray() {
	switch array.Type {
	case jingo.INT:
		fmt.Println("Found INT array")
		fmt.Println(array)
		fmt.Println("Shape", array.Shape)
		fmt.Println("Array length", len(array.Data.([]int)))
	}
}
