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

var id2pdef = map[IDType]pdef{}

type dyadFunct func(x, y A) (rv A, err error)
type monadFunct func(x A) (rv A, err error)
type pdef struct {
	atype AType
	Monad monadFunct
	Dyad  dyadFunct
	monadicRank,
	leftRank,
	rightRank,
	funcDepth int
	spelling IDType
}

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
	// var f2 dyad
	// f2 = add2
	// res, _ := f2(1, 2)
	// fmt.Println("f2: ", res)
	id2pdef[CASGN] = pdef{atype: ASGN, Dyad: asgn} // =.
	id2pdef[CPLUS] = pdef{atype: VERB, Dyad: add2} // +
}
