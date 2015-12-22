// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"fmt"
)

type Action int

const (
	adv Action = iota
	bident
	cmonad
	conj
	dyad
	fork
	hook
	is
	monad
	punc
	trident
	vadv
	vconj
	vdyad
	vfork
	vhook
	vis
	vmonad
	vpunc
)

const (
	AVN  AType = ADV | VERB | NOUN
	CAVN AType = CONJ | ADV | VERB | NOUN
	EDGE AType = MARK | ASGN | LPAR
)

type ptCase struct {
	pattern       [4]AType
	funcType      [2]Action
	begin, end, k int // don't know what k is yet
}

var ptCases [9]ptCase

func init() {
	ptCases[0] = ptCase{[4]AType{EDGE, VERB, NOUN, ANY}, [2]Action{monad, vmonad}, 1, 2, 1}
	ptCases[1] = ptCase{[4]AType{EDGE + AVN, VERB, VERB, NOUN}, [2]Action{monad, vmonad}, 2, 3, 2}
	ptCases[2] = ptCase{[4]AType{EDGE + AVN, NOUN, VERB, NOUN}, [2]Action{dyad, vdyad}, 1, 3, 2}
	ptCases[3] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, ADV, ANY}, [2]Action{adv, vadv}, 1, 2, 1}
	ptCases[4] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, CONJ, VERB + NOUN}, [2]Action{conj, vconj}, 1, 3, 1}
	ptCases[5] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, VERB, VERB}, [2]Action{trident, vfork}, 1, 3, 1}
	ptCases[6] = ptCase{[4]AType{EDGE, CAVN, CAVN, ANY}, [2]Action{bident, vhook}, 1, 2, 1}
	ptCases[7] = ptCase{[4]AType{NAME + NOUN, ASGN, CAVN, ANY}, [2]Action{is, vis}, 0, 2, 1}
	ptCases[8] = ptCase{[4]AType{LPAR, CAVN, RPAR, ANY}, [2]Action{punc, vpunc}, 0, 2, 0}
}

/*  The original from jsoftware.com
PT cases[] = {
 EDGE,      VERB,      NOUN, ANY,       jtmonad,   jtvmonad, 1,2,1,
 EDGE+AVN,  VERB,      VERB, NOUN,      jtmonad,   jtvmonad, 2,3,2,
 EDGE+AVN,  NOUN,      VERB, NOUN,      jtdyad,    jtvdyad,  1,3,2,
 EDGE+AVN,  VERB+NOUN, ADV,  ANY,       jtadv,     jtvadv,   1,2,1,
 EDGE+AVN,  VERB+NOUN, CONJ, VERB+NOUN, jtconj,    jtvconj,  1,3,1,
 EDGE+AVN,  VERB+NOUN, VERB, VERB,      jttrident, jtvfolk,  1,3,1,
 EDGE,      CAVN,      CAVN, ANY,       jtbident,  jtvhook,  1,2,1,
 NAME+NOUN, ASGN,      CAVN, ANY,       jtis,      jtvis,    0,2,1,
 LPAR,      CAVN,      RPAR, ANY,       jtpunc,    jtvpunc,  0,2,0,
};*/

func Parse(jt J, q []A) (z A, err error) {
	fmt.Println("in Parse")
	// problem:  deba expects an array, but we've got a slice of A's
	if _, err = deba(jt, DCPARSE, A{}, A{}, A{}); err != nil {
		return
	}
	q = append([]A{mark}, append(q, []A{mark, mark, mark, mark}...)...)
	z, err = Parsea(jt, q)
	debz()
	if err != nil {
		fmt.Println("Error on Parsea", err)
		return
	}
	return
}

func Parsea(jt J, q []A) (z A, err error) {
	fmt.Println("in Parsea", q)
	//(NUMERIC | JCHAR | BOX | SBOX | SBT)
	//fmt.Println("NUMERIC, JCHAR, BOX, SBOX, SBT", NUMERIC, JCHAR, BOX, SBOX, SBT)

	var i int
	var ptc ptCase
	stack := []A{}
	for i = 0; i < 4; i++ {
		stack = append(stack, q[len(q)-1])
		q = q[0 : len(q)-1]
	}
	for {
		stack = append(stack, q[len(q)-1])
		q = q[0 : len(q)-1]
		stp := len(stack) - 1 // stack top pos
		fmt.Println("top 4 stack", stack[stp-0].Type, stack[stp-1].Type, stack[stp-2].Type, stack[stp-3].Type)
		for i, ptc = range ptCases {
			fmt.Println("Checking pattern", i)
			pat := ptc.pattern
			if ((pat[0] & stack[stp-0].Type) != 0) &&
				((pat[1] & stack[stp-1].Type) != 0) &&
				((pat[2] & stack[stp-2].Type) != 0) &&
				((pat[3] & stack[stp-3].Type) != 0) {
				fmt.Println("found match", i)
				break
			}
		}
		if len(q) < 1 {
			break
		}
	}
	fmt.Println("pat pos", i)
	stp := len(stack) - 1
	fmt.Println("stack 4 at end", stack[stp-0].Type, stack[stp-1].Type, stack[stp-2].Type, stack[stp-3].Type)
	fmt.Println("q", q)
	if i < len(ptCases) {
		fmt.Println("Executing pattern", i)
	}
	//for i, ptc = range(ptCases){
	// 	fmt.Println("i, pattern", i, ptc.pattern)
	//}
	return
}
