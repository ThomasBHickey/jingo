// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
	//"unicode/utf8"
)

func JDo(jt *J, text string) {
	q, event := tokens(jt, text)
	if event != 0 {
		fmt.Println("enqueue failed", event)
	} else {
		fmt.Println("enqueue", q)
		z, err := parse(jt, q)
		if jt.Asgn {
			//fmt.Println("Assignment")
		} else {
			fmt.Println("result of Parse", err, z)
		}
	}

}
