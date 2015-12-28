// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	//"errors"
	"fmt"
)

var ctype = [128]CBType{
	00, 00, 00, 00, 00, 00, 00, 00, 00, CS, 00, 00, 00, 00, 00, 00, /* 0                  */
	00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, /* 1                  */
	CS, 00, 00, 00, 00, 00, 00, CQ, 00, 00, 00, 00, 00, 00, CD, 00, /* 2  !"#$%&'()*+,-./ */
	C9, C9, C9, C9, C9, C9, C9, C9, C9, C9, CC, 00, 00, 00, 00, 00, /* 3 0123456789:;<=>? */
	00, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 4 @ABCDEFGHIJKLMNO */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 00, 00, 00, 00, C9, /* 5 PQRSTUVWXYZ[\]^_ */
	00, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 6 `abcdefghijklmno */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 00, 00, 00, 00, 00} /* 7 pqrstuvwxyz{|}~  */
/*   1   2   3   4   5   6   7   8   9   a   b   c   d   e   f   */

var wtype [128]CBType

func init() {
	for i, v := range ctype {
		wtype[i] = v
	}
	wtype['N'] = CN
	wtype['B'] = CB
}

var id2pdef = map[IDType]A{}

func cid2pdef(c rune, idt IDType) (z A, OK bool) {
	if c < 128 {
		z = id2pdef[idt]
		return z, z.Type != NoAType
	}
	return
}

type dyadFunct func(jt *J, x, y A) (rv A, evn Event)
type monadFunct func(jt *J, x A) (rv A, evn Event)
type VAData struct {
	f1 monadFunct
	f2 dyadFunct
	f, // left conj or adverb argument
	g, // right conj. argument
	h A // auxiliary argument
	isAsgn bool
	flag   bool //not sure what gets flagged
	mr,    // monadic rank
	lr, // left dyadic rank
	rr, // right dyadic rank
	funcDepth int
	id IDType
}

//	id2pdef[CASGN] = pdef{atype: ASGN, Dyad: asgn} // =.
func add2(jt *J, x, w A) (A, Event) {
	ra := NewIntArray(x.Shape)
	if x.Type == INT && w.Type == INT && x.Length == 1 && w.Length == 1 {
		ra.Data = make([]int, x.Length)
		ra.Data.([]int)[0] = x.Data.([]int)[0] + w.Data.([]int)[0]
		fmt.Println("add2 returning", ra)
		return ra, 0
	}
	fmt.Println("add2 failed")
	return ra, EVVALUE
}

func asgn(jt *J, a A, w A) (A, Event) {
	fmt.Println("In func asgn!")
	return A{}, 0
}

type value struct{ z int }

func init() {
	fmt.Println("Hi from t.go!")
	id2pdef[CASGN] = NewVerbArray(VAData{f2: asgn, id: CASGN, isAsgn: true})
	id2pdef[CPLUS] = NewVerbArray(VAData{f2: add2})
}
