// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
	//"unicode/utf8"
)

// From dsusp.c:
/* deba() and debz() must be coded and executed in pairs */
/* in particular, do NOT do error exits between them     */
/* e.g. the following is a NO NO:                        */
/*    d=deba(...);                                       */
/*    ASSERT(blah,EVDOMAIN);                             */
/*    debz()                                             */

type DCType int // from jtype.h
const (
	DCPARSE DCType = 1 + iota
	DCSCRIPT
	DCCALL
	DCJUNK
)

func deba(jt *J, t DCType, x, y, fs A) (d []DST, evn Event) {
	fmt.Println("Call to jtdeba")
	return
}

func debz() {}
