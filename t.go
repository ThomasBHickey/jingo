package jingo

import (
	"fmt"
)

var ctype = [128]int{
	00, 00, 00, 00, 00, 00, 00, 00, 00, CS, 00, 00, 00, 00, 00, 00, /* 0                  */
	00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, 00, /* 1                  */
	CS, 00, 00, 00, 00, 00, 00, CQ, 00, 00, 00, 00, 00, 00, CD, 00, /* 2  !"#$%&'()*+,-./ */
	C9, C9, C9, C9, C9, C9, C9, C9, C9, C9, CC, 00, 00, 00, 00, 00, /* 3 0123456789:;<=>? */
	00, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 4 @ABCDEFGHIJKLMNO */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 00, 00, 00, 00, C9, /* 5 PQRSTUVWXYZ[\]^_ */
	00, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 6 `abcdefghijklmno */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 00, 00, 00, 00, 00} /* 7 pqrstuvwxyz{|}~  */
/*   1   2   3   4   5   6   7   8   9   a   b   c   d   e   f   */

var wtype [128]int

func init() {
	for i, v := range ctype {
		wtype[i] = v
	}
	wtype['N'] = CN
	wtype['B'] = CB
}

// pst [256]int  // not clear what pst is used for
var id2pdef = map[int]pdef{}

type dyad func(x, y int) (rv int, err error)
type monand func(x int) (rv int, err error)
type pdef struct {
	ptype     byte
	monadFunc monand
	dyadFunc  dyad
	monadicRank,
	leftRank,
	rightRank,
	funcDepth,
	spelling int
}

func add2(x, y int) (int, error) {
	return x + y, nil
}

type Val []int
type Array struct {
	atype              byte
	refCount, numAtoms int
	shape              []int
	value              Val
}

func asgn(a Array, w int) (Array, error) {
	fmt.Println("In func asgn!")
	return Array{}, nil
}

type value struct{ z int }

func init() {
	fmt.Println("Hi from t.go!")
	var f2 dyad
	f2 = add2
	res, _ := f2(1, 2)
	fmt.Println("f2: ", res)
	id2pdef[CGASGN] = pdef{dyadFunc: add2}
}
