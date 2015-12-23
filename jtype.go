// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file

package jingo

import (
	"fmt"
)

func init() {
	fmt.Println("SPARSE", SPARSE, "LIT", LIT, "SINT", SINT)
}

/* Values for AFLAG(x) field of type A                                     */
type AFLAG int

const (
	AFRO  = 1 << iota //1            /* read only; can't change data    */
	AFNJA             //2            /* non-J alloc; i.e. mem mapped    */
	AFSMM             //4            /* SMM managed                     */
//AFREL         //8            /* uses relative addressing        */
)

//go:generate stringer -type=AType
type AType int

const (
	/* Types for AT(x) field of type A                                         */
	/* Note: BOOL name conflict with ???; SCHAR name conflict with sqltypes.h  */
	NoAType AType = 0
	B01     AType = 1 << iota // (i)1L           /* B  boolean                      */
	LIT                       // (i)2L           /* C  literal (character)          */
	INT                       // (i)4L           /* I  integer                      */
	FL                        // (i)8L           /* D  double (IEEE floating point) */
	CMPX                      // (i)16L          /* Z  complex                      */
	BOX                       // (i)32L          /* A  boxed                        */
	XNUM                      // (i)64L          /* X  extended precision integer   */
	RAT                       // (i)128L         /* Q  rational number              */
	BIT                       // (i)256L         /* BT bit boolean                  */
	SB01                      // (i)1024L        /* P  sparse boolean               */
	SLIT                      // (i)2048L        /* P  sparse literal (character)   */
	SINT                      // (i)4096L        /* P  sparse integer               */
	SFL                       // (i)8192L        /* P  sparse floating point        */
	SCMPX                     // (i)16384L       /* P  sparse complex               */
	SBOX                      // (i)32768L       /* P  sparse boxed                 */
	SBT                       // (i)65536L       /* SB symbol                       */
	C2T                       // (i)131072L      /* C2 unicode (2-byte characters)  */
	VERB                      // (i)262144L      /* V  verb                         */
	ADV                       // (i)524288L      /* V  adverb                       */
	CONJ                      // (i)1048576L     /* V  conjunction                  */
	ASGN                      // (i)2097152L     /* I  assignment                   */
	MARK                      // (i)4194304L     /* I  end-of-stack marker          */
	SYMB                      // (i)8388608L     /* I  locale (symbol table)        */
	CONW                      // (i)16777216L    /* CW control word                 */
	NAME                      // (i)33554432L    /* NM name                         */
	LPAR                      // (i)67108864L    /* I  left  parenthesis            */
	RPAR                      // (i)134217728L   /* I  right parenthesis            */
	XD                        // (i)268435456L   /* DX extended floating point      */
	XZ                        // (i)536870912L   /* ZX extended complex             */

	ANY     AType = 0xffffffff
	SPARSE  AType = (SB01 | SINT | SFL | SCMPX | SLIT | SBOX)
	DENSE   AType = (NOUN &^ SPARSE)
	NUMERIC AType = (B01 | BIT | INT | FL | CMPX | XNUM | RAT | XD | XZ | SB01 | SINT | SFL | SCMPX)
	DIRECT  AType = (LIT | C2T | B01 | BIT | INT | FL | CMPX | SBT)
	JCHAR   AType = (LIT | C2T | SLIT)
	NOUN    AType = (NUMERIC | JCHAR | BOX | SBOX | SBT)
	FUNC    AType = (VERB | ADV | CONJ)
	RHS     AType = (NOUN | FUNC)
	IS1BYTE AType = (B01 | LIT)
	LAST0   AType = (B01 | LIT | C2T | NAME)
)

// func init() {
// 	fmt.Printf("XZ:  %x\n", XZ)
// 	fmt.Printf("ANY: %x\n", ANY)
// 	fmt.Printf("ANY & XZ %x\n", ANY & XZ)
// }

type DST struct { /* 1 2 3                                                        */
	dclnk  *DST /* x x x  link to next stack entry                              */
	dca    A    /*     x  fn/op name                                            */
	dcf    A    /*     x  fn/op                                                 */
	dcx    A    /*     x  left argument                                         */
	dcy    A    /* x x x  tokens text        ; right argument                  */
	dcloc  A    /*     x  local symb table (0 if not explicit)                  */
	dcc    A    /*     x  control matrix   (0 if not explicit)                  */
	dci    int  /* x x x  index  next index  ; ptr to line #                   */
	dcj    int  /*   x x        ; prev index  ; error #                         */
	dcn    int  /*   x x        ; line #      ; ptr to symb entry               */
	dcm    int  /*   x x        ; script index; # of non-locale part of name    */
	dcstop int  /*     x  the last stop in this function                        */
	dctype byte /* x x x  type of entry (see #define DC*)                       */
	dcsusp bool /* x   x  1 iff begins a debug suspension                       */
	dcss   byte /*     x  single step code                                      */
}
