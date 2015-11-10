package main

import (
	"bufio"
	"fmt"
	"os"
	//"github.com/ThomasBHickey/fileserver"
)

const (
	SS  = iota // Space
	SX         // Other
	SA         // Alphanumerics
	SN         // N
	SNB        // NB
	SNZ        // NB.
	S9         //Numeric
	SQ         // Quote
	SQQ        // Even quotes
	SZ         // Trailing comment
)
const (
	CX = iota //0 /* other                                   */
	CS        //1            /* space or tab                            */
	CA        //2            /* letter                                  */
	CN        //3            /* N (capital N)                           */
	CB        //4            /* B (capital B)                           */
	C9        //5            /* digit or sign (underscore)              */
	CD        //6            /* dot                                     */
	CC        //7            /* colon                                   */
	CQ        //8            /* quote    */
)

func ft() {}

var test2 = 3

type Action func(int, int) int

//const A1 Action =3
var A1 = func(j, pos int) int { fmt.Printf("in A1\n"); return j }
var A0 = func(j, pos int) int { fmt.Printf("in A0\n"); return j }
var EN = func(j, pos int) int { fmt.Printf("in EN %d %d\n", j, pos); return pos }
var EI = func(j, pos int) int { fmt.Printf("in EI %d %d\n", j, pos); return pos }

type sa struct { // state pair
	state  int
	action Action
}

var ctype = [128]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, CS, 0, 0, 0, 0, 0, 0, /* 0                  */
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, /* 1                  */
	CS, 0, 0, 0, 0, 0, 0, CQ, 0, 0, 0, 0, 0, 0, CD, 0, /* 2  !"#$%&'()*+,-./ */
	C9, C9, C9, C9, C9, C9, C9, C9, C9, C9, CC, 0, 0, 0, 0, 0, /* 3 0123456789:;<=>? */
	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 4 @ABCDEFGHIJKLMNO */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, C9, /* 5 PQRSTUVWXYZ[\]^_ */
	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, /* 6 `abcdefghijklmno */
	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, 0} /* 7 pqrstuvwxyz{|}~  */

var state = [][]sa{
	/*SS */ {sa{SX, EN}, sa{SS, A0}, sa{SA, EN}, sa{SN, EN}, sa{SA, EN}, sa{S9, EN}, sa{SX, EN}, sa{SX, EN}, sa{SQ, EN}},
	/*SX */ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, A0}, sa{SX, A0}, sa{SQ, EI}},
	/*SA */ {sa{SX, EI}, sa{SS, EI}, sa{SA, A0}, sa{SA, A0}, sa{SA, A0}, sa{SA, A0}, sa{SX, A0}, sa{SX, A0}, sa{SQ, EI}},
	/*SN */ {sa{SX, EI}, sa{SS, EI}, sa{SA, A0}, sa{SA, A0}, sa{SNB, A0}, sa{SA, A0}, sa{SX, A0}, sa{SX, A0}, sa{SQ, EI}},
	/*SNB*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, A0}, sa{SA, A0}, sa{SA, A0}, sa{SA, A0}, sa{SNZ, A0}, sa{SX, A0}, sa{SQ, EI}},
	/*SNZ*/ {sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SX, A0}, sa{SX, A0}, sa{SZ, A0}},
	/*S9 */ {sa{SX, EI}, sa{SS, EI}, sa{S9, A0}, sa{S9, A0}, sa{S9, A0}, sa{S9, A0}, sa{S9, A0}, sa{SX, A0}, sa{SQ, EI}},
	/*SQ */ {sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQ, A0}, sa{SQQ, A0}},
	/*SQQ*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, EI}, sa{SX, EI}, sa{SQ, A0}},
	/*SZ */ {sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}, sa{SZ, A0}}}

func main() {
	fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	//fmt.Println("test:%d", state[0][0].action(3, 2))
	curState := SS
	j := 0
	var thisSA sa
	var thisCtype int
	for pos, rune := range text {
		fmt.Printf("%#U starts at byte position %d\n", rune, pos)
		if rune > 127 {
			fmt.Printf("%d\n", rune)
			thisCtype = CA  // call it a letter
		} else {
			if ctype[rune] > 0 {
				fmt.Printf("ctype %d\n", ctype[rune])
				thisCtype = ctype[rune]
			}
		}
		fmt.Printf("curState %d, ctype %d\n", curState, ctype[rune])
		thisSA = state[curState][thisCtype]
		j = thisSA.action(j,pos)
		fmt.Printf("j=%d\n",j)
		curState = thisSA.state
	}

	//fileserver.Server()
}
