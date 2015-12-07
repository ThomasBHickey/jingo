package jingo

import ()

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
	pattern        [4]AType
	funcType       [2]Action
	start, stop, q int // don't know what q is yet
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
