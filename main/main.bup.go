package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/ThomasBHickey/jingo"
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

var noOp = func(j, pos int, wps []wp) (int, []wp) { return j, wps }
var A1 = func(j, pos int, wps []wp) (int, []wp) { fmt.Print("A1");return j, wps }
var A0 = func(j, pos int, wps []wp) (int, []wp) { fmt.Print("A0");return j, wps }
//var EN = func(j, pos int, wps []wp) (int, []wp) { fmt.Print("EN");return pos, wps }
var emit = func(j, pos int, wps []wp) (int, []wp) {
	fmt.Printf("in EI %d %d\n", j, pos-1)
	wps = append(wps, wp{j, pos - 1})
	fmt.Printf("len of wps in EI %d\n", len(wps))
	return pos, wps
}
const (
	E0 = iota
	EI // emit
	EN
	)

type sa struct { // state pair
	state, action  int
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
	/*SX */ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SA */ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SN */ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SNB, E0}, sa{SA, E0}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SNB*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SNZ, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SNZ*/ {sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SX, E0}, sa{SX, E0}, sa{SZ, E0}},
	/*S9 */ {sa{SX, EI}, sa{SS, EI}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SQ */ {sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQQ, E0}},
	/*SQQ*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, EI}, sa{SX, EI}, sa{SQ, E0}},
	/*SZ */ {sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}}}

func main() {
	fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	jingo.Scan(text)
	return
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
		//fmt.Printf("%#U starts at byte position %d\n", rune, bpos)
		thisCtype = CA // default
		if rune < 128 {
			thisCtype = ctype[rune]
		}
		//fmt.Printf("curState %d, ctype %d\n", curState, thisCtype)
		thisSA = state[curState][thisCtype]
		j, wps = thisSA.action(j, bpos, wps)
		//fmt.Printf("j=%d\n", j)
		curState = thisSA.state
		runes = append(runes, rune)
		rpos = rpos + 1
	}
	//fmt.Printf("length of wps %d\n", len(wps))
	for _, wp := range wps {
		fmt.Print(wp)
		fmt.Printf("%s\n", text[wp.start:wp.end+1])
		//fmt.Printf("%d %d\n", wp.start wp.end)
	}
}
