// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
)

func constr(s string) (A, Event) {
	fmt.Println("In constr with", s)
	return A{}, 0
}
