// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
	"strconv"
)

func connum(s string) (A, bool) {
	i, err := strconv.Atoi(s)
	fmt.Println("In connum with", s, i, err)
	if err != nil {
		return A{}, false
	}
	return NewSIntArray(i), true
}
