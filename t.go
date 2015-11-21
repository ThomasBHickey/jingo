package jingo

import (
	"fmt"
)

var ctype = [128]int{
	00, 0, 0, 0, 0, 0, 0, 0, 0, CS, 0, 0, 0, 0, 0, 0, /* 0                  */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 1                  */
	CS, 0, 0, 0, 0, 0, 0, CQ, 0, 0, 0, 0, 0, 0, CD, 0, /* 2  !"#$%&'()*+,-./ */
	C9, C9, C9, C9, C9, C9, C9, C9, C9, C9, CC, 0, 0, 0, 0, 0, /* 3 0123456789:;<=>? */
	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 4 @ABCDEFGHIJKLMNO */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, C9, /* 5 PQRSTUVWXYZ[\]^_ */
	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 6 `abcdefghijklmno */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, 0} /* 7 pqrstuvwxyz{|}~  */
/*   1   2   3   4   5   6   7   8   9   a   b   c   d   e   f   */

var wtype [128]int

func init() {
	for i, v := range ctype {
		wtype[i] = v
	}
	wtype['N'] = CN
	wtype['B'] = CB
}

type dyad func(x, y int) (rv int, err error)
type monand func(x int) (rv int, err error)
type pdef struct {
	monadFunc monand
	dyadFunc  dyad
	monadicRank,
	leftRank,
	rightRank,
	funcDepth,
	spelling int
	pst [256]int
}

func add2(x, y int) (int, error) {
	return x + y, nil
}

type value struct{ z int }

func init() {
	fmt.Println("Hi from t.go!")
	var f2 dyad
	f2 = add2
	res, _ := f2(1, 2)
	fmt.Println("f2: ", res)
}
