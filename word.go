package jingo

import (
	"fmt"
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

/*const (
	CX = iota //0            // other
	CS        //1            // space or tab
	CA        //2            // letter
	CN        //3            // N (capital N)
	CB        //4            // B (capital B)
	C9        //5            // digit or sign (underscore)
	CD        //6            // dot
	CC        //7            // colon
	CQ        //8            // quote
)*/ //moved to t.go

// var ctype = [128]int{  // back in t.go
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, CS, 0, 0, 0, 0, 0, 0, // 0
// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 1
// 	CS, 0, 0, 0, 0, 0, 0, CQ, 0, 0, 0, 0, 0, 0, CD, 0, // 2  !"#$%&'()*+,-./
// 	C9, C9, C9, C9, C9, C9, C9, C9, C9, C9, CC, 0, 0, 0, 0, 0, // 3 0123456789:;<=>?
// 	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, // 4 @ABCDEFGHIJKLMNO
// 	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, C9, // 5 PQRSTUVWXYZ[\]^_
// 	0, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, // 6 `abcdefghijklmno
// 	CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, CA, 0, 0, 0, 0, 0} // 7 pqrstuvwxyz{|}~
const (
	E0 = iota
	EI // emit
	EN
)

type sa struct{ new, effect int }
type wp struct{ Start, End int } // word position
var state = [10][9]sa{
	/*SS */ {sa{SX, EN}, sa{SS, E0}, sa{SA, EN}, sa{SN, EN}, sa{SA, EN}, sa{S9, EN}, sa{SX, EN}, sa{SX, EN}, sa{SQ, EN}},
	/*SX */ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SA */ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SN */ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SNB, E0}, sa{SA, E0}, sa{SX, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SNB*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SA, E0}, sa{SNZ, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SNZ*/ {sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SX, E0}, sa{SX, E0}, sa{SZ, E0}},
	/*S9 */ {sa{SX, EI}, sa{SS, EI}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{S9, E0}, sa{SX, E0}, sa{SQ, EI}},
	/*SQ */ {sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQ, E0}, sa{SQQ, E0}},
	/*SQQ*/ {sa{SX, EI}, sa{SS, EI}, sa{SA, EI}, sa{SN, EI}, sa{SA, EI}, sa{S9, EI}, sa{SX, EI}, sa{SX, EI}, sa{SQ, E0}},
	/*SZ */ {sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}, sa{SZ, E0}}}

func runeToCType(r rune) int {
	if r < 128 {
		return ctype[r]
	} else {
		return CA
	}
}
func runeToWType(r rune) int {
	if r < 128 {
		return wtype[r]
	} else {
		return CA
	}
}
func Scan(text string) []wp {
	fmt.Println("In Scan", text)
	nv := false    // numeric value being built
	cs := SS       // current state
	wps := []wp{}  // word positions
	t := false     // true if building numeric vector (S9)
	var b int      // beginning index of current word
	var xb, xe int // beginning/end index of current numeric vector
	var e int      // effect associated with state
	//var bpos int
	for bpos, rune := range text {
		//fmt.Printf("%#U starts at byte position %d\n", rune, bpos)
		ct := runeToCType(rune)
		// ct := CA // default current char type
		// if rune < 128 {
		// 	ct = ctype[rune]
		// }
		fmt.Println("curState", cs, "ctype", ct, "rune", rune)
		p := state[cs][ct]
		if e = p.effect; e == EI {
			if t := t && (cs == S9); t {
				if !nv {
					nv = true
					xb = b
				}
				xe = bpos
			} else {
				if nv {
					nv = false
					fmt.Println("emit 1", text[xb:xe])
					wps = append(wps, wp{xb, xe})
				}
				fmt.Println("emit 2", text[b:bpos])
				wps = append(wps, wp{b, bpos})
			}
		}
		cs = p.new
		if e != E0 {
			b = bpos
			t = cs == S9
		}
	}
	//bpos = bpos+1
	//fmt.Println("finished loop", "cs", cs, "t", t, "nv", nv, "xb", xb, "xe", xe, "b", b, "bpos", bpos)
	if cs == SQ {
		return []wp{} // needs error condition
	}
	t = t && (cs == S9)
	if t {
		if nv {
			wps = append(wps, wp{xb, len(text)})
			fmt.Println("emit 3a", xb, len(text), text[xb:len(text)])
		} else {
			wps = append(wps, wp{b, len(text)})
			fmt.Println("emit 3b", b, len(text), text[b:len(text)])
		}
	} else {
		if nv {
			wps = append(wps, wp{xb, xe})
			fmt.Println("emit 4:", xb, xe, text[xb:xe])
		}
		if cs != SS {
			wps = append(wps, wp{b, len(text)})
			fmt.Println("emit 5:", b, len(text), text[b:len(text)])
		}
	}
	for _, wp := range wps {
		s := text[wp.Start:wp.End]
		fmt.Println("wp", s)
		if len(s) >= 0 {
			c0 := []rune(s)[0]
			fmt.Println("wtype", c0, runeToWType(c0), "spellin", spellIn[s])
			id := spellIn[s]
			pdef, ok := id2pdef[id]
			if ok{
				dyres, _ := pdef.dyadFunc(2,3)
				fmt.Println("pdef dyres", dyres)
			}
		}
	}
	return wps
}
