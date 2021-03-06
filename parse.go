// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"fmt"
)

//type Action int

type Action func(jt *J, b, e int, stack []A) (rv A, evn Event)

// const (
// 	dyad Action = dyadFunct
// )

// const (
// 	adv Action = iota
// 	bident
// 	cmonad
// 	conj
// 	dyad
// 	fork
// 	hook
// 	is
// 	monad
// 	punc
// 	trident
// 	vadv
// 	vconj
// 	vdyad
// 	vfork
// 	vhook
// 	vis
// 	vmonad
// 	vpunc
// )

func dyad(jt *J, b, e int, stack []A) (z A, evn Event) {
	showArraySliceR(jt, "In dyad", stack)
	if (b - e) != 2 {
		return z, EVSYNTAX
	}
	jt.Log.Println("dyad 1st param", stack[e+2])
	jt.Log.Println("dyad 2nd param", stack[e])
	//return stack[e+1].Data.(VAData).f2(jt, stack[e+2], stack[e])
	sfold := jt.sf
	verbA := stack[e+1]
	jt.sf = verbA
	z, evn = verbA.Data.(VAData).f2(jt, stack[e+2], stack[e])
	if evn != 0 {
		jt.Log.Println("dyad call", Event2String[evn])
	}
	jt.sf = sfold
	return
}

func monad(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In monad (not implemented)", b, e, stack)
	return
}

func adv(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In adv (not implemented)", b, e, stack)
	return
}

func conj(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In conj (not implemented)", b, e, stack)
	return
}

func trident(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In trident (not implemented)", b, e, stack)
	return
}

func bident(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In bident (not implemented", b, e, stack)
	return
}
func vhook(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In vhook (not implemented)", b, e, stack)
	return
}
func is(jt *J, b, e int, stack []A) (z A, evn Event) {
	showArraySliceR(jt, "In 'is'", stack)
	if (b - e) != 2 {
		return z, EVSYNTAX
	}
	jt.Log.Println("is 1st param", stack[e+2])
	jt.Log.Println("is 2nd param", stack[e])
	return stack[e+1].Data.(VAData).f2(jt, stack[e+2], stack[e])
}

func punc(jt *J, b, e int, stack []A) (z A, evn Event) {
	jt.Log.Println("In punc (not implemented)", b, e, stack)
	return
}

const (
	AVN      AType = ADV | VERB | NOUN
	CAVN     AType = CONJ | ADV | VERB | NOUN
	EDGE     AType = MARK | ASGN | LPAR
	NAMENOUN AType = NAME | NOUN
)

type ptCase struct {
	pattern       [4]AType
	actions       [2]Action
	begin, end, k int // don't know what k is yet
}

var ptCases [9]ptCase

const NCASES = len(ptCases)

func init() {
	ptCases[0] = ptCase{[4]AType{EDGE, VERB, NOUN, ANY}, [2]Action{monad, vmonad}, 1, 2, 1}
	ptCases[1] = ptCase{[4]AType{EDGE + AVN, VERB, VERB, NOUN}, [2]Action{monad, vmonad}, 2, 3, 2}
	ptCases[2] = ptCase{[4]AType{EDGE + AVN, NOUN, VERB, NOUN}, [2]Action{dyad, vdyad}, 1, 3, 2}
	ptCases[3] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, ADV, ANY}, [2]Action{adv, vadv}, 1, 2, 1}
	ptCases[4] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, CONJ, VERB + NOUN}, [2]Action{conj, vconj}, 1, 3, 1}
	ptCases[5] = ptCase{[4]AType{EDGE + AVN, VERB + NOUN, VERB, VERB}, [2]Action{trident, vfork}, 1, 3, 1}
	ptCases[6] = ptCase{[4]AType{EDGE, CAVN, CAVN, ANY}, [2]Action{bident, vhook}, 1, 2, 1}
	//ptCases[7] = ptCase{[4]AType{NAME + NOUN, ASGN, CAVN, ANY}, [2]Action{is, vis}, 0, 2, 1}
	ptCases[7] = ptCase{[4]AType{NAMENOUN, ASGN, CAVN, ANY}, [2]Action{is, vis}, 0, 2, 1}
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

func parse(jt *J, q []A) (z A, evn Event) {
	jt.Log.Println("in Parse")
	// problem:  deba expects an array, but we've got a slice of A's
	if _, evn = deba(jt, DCPARSE, A{}, A{}, A{}); evn != 0 {
		return
	}
	q = append([]A{mark}, append(q, []A{mark, mark, mark, mark}...)...)
	z, evn = parsea(jt, q)
	debz()
	if evn != 0 {
		jt.Log.Println("Error on parsea", evn, Event2String[evn])
		return
	}
	return
}
func showArraySlice(jt *J, msg string, aslice []A) {
	jt.Log.Println(msg)
	s := ""
	for i := 0; i < len(aslice); i++ {
		a := aslice[i]
		s = s + fmt.Sprintf(" (%d:", i)
		switch a.Type {
		case INT:
			s = s + fmt.Sprintf(" INT=%d)", a.Data.([]int)[0])
		case VERB:
			s = s + fmt.Sprintf(" VERB=%s)", a.Data.(VAData).id)
		default:
			s = s + fmt.Sprintf(" %s)", a.Type)
		}
	}
	jt.Log.Println(s)
}

func showArraySliceR(jt *J, msg string, aslice []A) {
	jt.Log.Println(msg)
	s := ""
	for i := len(aslice) - 1; i >= 0; i-- {
		a := aslice[i]
		s = s + fmt.Sprintf(" (%d:", i)
		switch a.Type {
		case INT:
			s = s + fmt.Sprintf(" INT=%d)", a.Data.([]int)[0])
		case VERB:
			s = s + fmt.Sprintf(" VERB=%s)", a.Data.(VAData).id)
		default:
			s = s + fmt.Sprintf(" %s)", a.Type)
		}
	}
	jt.Log.Println(s)
}

func parsea(jt *J, q []A) (z A, evn Event) {
	showArraySlice(jt, "In Parsea", q)
	// Return if empty
	if len(q) == 0 {
		return z, EVVALUE
	}
	// other setup
	jt.Asgn = false
	jt.Parsercalls++
	stack := []A{}
	var i int
	// C code manages stack & queue as one list
	for i = 0; i < 4; i++ { // probably could be done more efficiently
		stack = append(stack, q[len(q)-1])
		q = q[0 : len(q)-1]
	}
	for len(stack) > 1 {
		showArraySliceR(jt, "Main execution loop", stack)
		// cycle through cases
		stp := len(stack) - 1 // stack top pos
		for i = 0; i < len(ptCases); i++ {
			pat := ptCases[i].pattern
			if ((pat[0] & stack[stp-0].Type) != 0) &&
				((pat[1] & stack[stp-1].Type) != 0) &&
				((pat[2] & stack[stp-2].Type) != 0) &&
				((pat[3] & stack[stp-3].Type) != 0) {
				jt.Log.Println("found match", i, "length of q", len(q))
				break
			}
		}
		if i < NCASES {
			// execute the case
			jt.Log.Println("Executing pattern", i)
			ptCase := ptCases[i]
			b, e := ptCase.begin, ptCase.end
			j, k := stp-b, stp-e
			f := ptCase.actions[0]
			jt.Asgn = stack[k+1].Type == ASGN
			if z, evn = f(jt, j, k, stack); evn != 0 {
				jt.Log.Println("ERROR event from f call:", Event2String[evn])
				return
			}
			// finish execution
			jt.Log.Println("updating stack at", k, "using", z)
			stack[k] = z
			showArraySliceR(jt, "", stack)
			// jt.Log.Println("changing stack at :k+1 and j:", k+1, j)
			stack = append(stack[:k+1], stack[j+1:]...)
			showArraySliceR(jt, "stack after trunc", stack)
		} else {
			// move from queue to stack
			if len(q) == 0 {
				break
			}
			jt.Log.Println("moving from q to stack: qend=", q[len(q)-1])
			stack = append(stack, q[len(q)-1])
			q = q[0 : len(q)-1]
			stop := len(stack) - 1
			jt.Log.Println("stack[stop, stop-1]", stack[stop].Type, stack[stop-1].Type)
			if (stack[stop].Type&NAME != 0) && ((stack[stop-1].Type & ASGN) == 0) {
				jt.Log.Println("Need to replace name with value", stack[stop])
				jt.Log.Println("value of name", jt.Symb[stack[stop].Data.(NameData).name])
				stack[stop] = jt.Symb[stack[stop].Data.(NameData).name]
			}
		}
	}
	// cleanup
	stack = stack[4:] // drop those 4 MARK arrays
	showArraySliceR(jt, "stack at end of parsea", stack)
	if ((stack[0].Type & CAVN) == 0) || (stack[1].Type != MARK) {
		return z, EVSYNTAX
	}
	// return value
	return stack[0], 0
}
