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
	shape = append(append(shape, 3), 4)
	a := NewIntArray(shape)
	fmt.Println("a Array", reflect.TypeOf(a.Data))
	fmt.Println("a as []int64", a.Data.([]int64))
	a.Data.([]int64)[3] = 4
	a.ShowArray()
	b := NewByteArray(shape)
	b.Data.([]byte)[0] = 'b'
	b.ShowArray()
}
