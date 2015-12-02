package jingo

import (
	"fmt"
)

func init() {
	fmt.Println("SPARSE", SPARSE, "LIT", LIT, "SINT", SINT)
}

//go:generate stringer -type=AType
type AType int

const (
	/* Types for AT(x) field of type A                                         */
	/* Note: BOOL name conflict with ???; SCHAR name conflict with sqltypes.h  */

	B01   AType = 1 << iota // (i)1L           /* B  boolean                      */
	LIT                     // (i)2L           /* C  literal (character)          */
	INT                     // (i)4L           /* I  integer                      */
	FL                      // (i)8L           /* D  double (IEEE floating point) */
	CMPX                    // (i)16L          /* Z  complex                      */
	BOX                     // (i)32L          /* A  boxed                        */
	XNUM                    // (i)64L          /* X  extended precision integer   */
	RAT                     // (i)128L         /* Q  rational number              */
	BIT                     // (i)256L         /* BT bit boolean                  */
	SB01                    // (i)1024L        /* P  sparse boolean               */
	SLIT                    // (i)2048L        /* P  sparse literal (character)   */
	SINT                    // (i)4096L        /* P  sparse integer               */
	SFL                     // (i)8192L        /* P  sparse floating point        */
	SCMPX                   // (i)16384L       /* P  sparse complex               */
	SBOX                    // (i)32768L       /* P  sparse boxed                 */
	SBT                     // (i)65536L       /* SB symbol                       */
	C2T                     // (i)131072L      /* C2 unicode (2-byte characters)  */
	VERB                    // (i)262144L      /* V  verb                         */
	ADV                     // (i)524288L      /* V  adverb                       */
	CONJ                    // (i)1048576L     /* V  conjunction                  */
	ASGN                    // (i)2097152L     /* I  assignment                   */
	MARK                    // (i)4194304L     /* I  end-of-stack marker          */
	SYMB                    // (i)8388608L     /* I  locale (symbol table)        */
	CONW                    // (i)16777216L    /* CW control word                 */
	NAME                    // (i)33554432L    /* NM name                         */
	LPAR                    // (i)67108864L    /* I  left  parenthesis            */
	RPAR                    // (i)134217728L   /* I  right parenthesis            */
	XD                      // (i)268435456L   /* DX extended floating point      */
	XZ                      // (i)536870912L   /* ZX extended complex             */

	ANY     AType = 0xffffffff
	SPARSE        = (SB01 | SINT | SFL | SCMPX | SLIT | SBOX)
	DENSE         = (NOUN &^ SPARSE)
	NUMERIC       = (B01 | BIT | INT | FL | CMPX | XNUM | RAT | XD | XZ | SB01 | SINT | SFL | SCMPX)
	DIRECT        = (LIT | C2T | B01 | BIT | INT | FL | CMPX | SBT)
	JCHAR         = (LIT | C2T | SLIT)
	NOUN          = (NUMERIC | JCHAR | BOX | SBOX | SBT)
	FUNC          = (VERB | ADV | CONJ)
	RHS           = (NOUN | FUNC)
	IS1BYTE       = (B01 | LIT)
	LAST0         = (B01 | LIT | C2T | NAME)
)

// func init() {
// 	fmt.Printf("XZ:  %x\n", XZ)
// 	fmt.Printf("ANY: %x\n", ANY)
// 	fmt.Printf("ANY & XZ %x\n", ANY & XZ)
// }
