package jingo

import (
	"errors"
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

func cid2pdef(idt IDType) A {
	if idt < 128 {
		return id2pdef[idt]
	}
	return NewVerbArray(VAData{})
}

type dyadFunct func(x, y A) (rv A, err error)
type monadFunct func(x A) (rv A, err error)
type VAData struct {
	f1 monadFunct
	f2 dyadFunct
	f, // left conj or adverb argument
	g, // right conj. argument
	h A // auxiliary argument
	flag bool //not sure what gets flagged
	mr,  // monadic rank
	lr, // left dyadic rank
	rr, // right dyadic rank
	funcDepth int
	id IDType
}

//	id2pdef[CASGN] = pdef{atype: ASGN, Dyad: asgn} // =.
func add2(x, w A) (A, error) {
	ra := NewIntArray(x.Shape)
	if x.Type == INT && w.Type == INT && x.Length == 1 && w.Length == 1 {
		ra.Data = make([]int64, x.Length)
		ra.Data.([]int64)[0] = x.Data.([]int64)[0] + w.Data.([]int64)[0]
		return ra, nil
	}
	return ra, errors.New("Unexpected arrays in add2")
}

func asgn(a A, w A) (A, error) {
	fmt.Println("In func asgn!")
	return A{}, nil
}

type value struct{ z int }

func init() {
	fmt.Println("Hi from t.go!")
	id2pdef[CASGN] = NewVerbArray(VAData{f2: asgn, id: CASGN})
	id2pdef[CPLUS] = NewVerbArray(VAData{f2: add2})
}
