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

type Action func(int, int, []wp) (int, []wp)

var noOp = func(j, pos int, wps []wp)(int, []wp) { return j, wps }
var A1 = noOp
var A0 = noOp
var EN = func(j, pos int, wps []wp)(int, []wp) { return pos, wps}
var emit = func(j, pos int,  wps []wp) (int, []wp) {
	fmt.Printf("in EI %d %d\n", j, pos-1);
	wps = append(wps, wp{j, pos-1})
	fmt.Printf("len of wps in EI %d\n", len(wps))
	return pos, wps }
var EI = emit
type sa struct { // state pair
	state  int
	action Action
}
type wp struct { // word position
	start, end int
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
	var rpos int
	var runes []rune
	var wps []wp
	for bpos, rune := range text {
		fmt.Printf("%#U starts at byte position %d\n", rune, bpos)
		if rune > 127 {
			thisCtype = CA // call it a letter
		} else {
			thisCtype = ctype[rune]
		}
		fmt.Printf("curState %d, ctype %d\n", curState, thisCtype)
		thisSA = state[curState][thisCtype]
		j, wps = thisSA.action(j, bpos, wps)
		fmt.Printf("j=%d\n", j)
		curState = thisSA.state
		runes = append(runes, rune)
		rpos = rpos + 1
	}
	fmt.Printf("length of wps %d\n", len(wps))
	for _, wp := range wps {
		fmt.Print(wp)
		fmt.Printf("%s\n", text[wp.start:wp.end+1])
		//fmt.Printf("%d %d\n", wp.start wp.end)
	}
}
