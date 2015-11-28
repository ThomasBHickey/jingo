package array

import (
	"fmt"
	//"github.com/ThomasBHickey/jingo"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	fmt.Println("Testinit")
	var shape []int
	shape = append(shape, 1)
	a := NewIntArray(shape)
	fmt.Println("a Array", reflect.TypeOf(a.Data))
	fmt.Println("a as []int", a.Data.([]int))
	a.Data = append(a.Data.([]int), 3)
	a.Data = append(a.Data.([]int), 4)
	a.ShowArray()
}
