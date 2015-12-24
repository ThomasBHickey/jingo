// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
	//"unicode/utf8"
)

type CSType int // char state type
const (
	SS  CSType = iota // Space
	SX                // Other
	SA                // Alphanumerics
	SN                // N
	SNB               // NB
	SNZ               // NB.
	S9                //Numeric
	SQ                // Quote
	SQQ               // Even quotes
	SZ                // Trailing comment
)

//go:generate stringer -type=CSType

type EffType int // effect
const (
	E0 EffType = iota
	EI         // emit
	EN
)

//go:generate stringer -type=EffType

type sa struct {
	new    CSType
	effect EffType
}

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

func runeToCType(r rune) CBType {
	if r < 128 {
		return ctype[r]
	} else {
		return CA
	}
}
func runeToWType(r rune) CBType {
	if r < 128 {
		return wtype[r]
	} else {
		return CA
	}
}

type wp struct{ Start, End int } // word position

func Scan(jt J, text string) []wp {
	fmt.Println("In Scan", text)
	nv := false    // numeric value being built
	cs := SS       // current state
	wps := []wp{}  // word positions
	t := false     // true if building numeric vector (S9)
	var b int      // beginning index of current word
	var xb, xe int // beginning/end index of current numeric vector
	var e EffType  // effect associated with state

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
					fmt.Println("emit 1:", text[xb:xe])
					wps = append(wps, wp{xb, xe})
				}
				fmt.Println("emit 2:", text[b:bpos])
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
	//snpdefs := []snpdef{}
	if cs == SQ {
		return wps // needs error condition
	}
	t = t && (cs == S9)
	if t {
		if nv {
			wps = append(wps, wp{xb, len(text)})
			fmt.Println("emit 3a:", xb, len(text), text[xb:len(text)])
		} else {
			wps = append(wps, wp{b, len(text)})
			fmt.Println("emit 3b:", b, len(text), text[b:len(text)])
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
	// for _, wp := range wps {
	// 	s := text[wp.Start:wp.End]
	// 	fmt.Println("wp", s)
	// 	if len(s) >= 0 {
	// 		c0 := []rune(s)[0]
	// 		fmt.Println("wtype", c0, runeToWType(c0), "spellin", spellIn[s])
	// 		id := spellIn[s]
	// 		pdef, ok := id2pdef[id]
	// 		snpdefs = append(snpdefs, snpdef{s: s, id: id, pd: pdef})
	// 		if ok {
	// 			dyres, _ := pdef.Dyad(A{}, A{})
	// 			fmt.Println("pdef dyres", dyres)
	// 		}
	// 	}
	// }
	// for _, sp := range snpdefs {
	// 	fmt.Println("s", sp.s, "id", spellOut[sp.id], "pd", sp.pd)
	// }
	return wps
}

func runeIfNotb(p rune, b bool) rune {
	if b {
		return 0
	}
	return p
}

func Enqueue(jt J, wps []wp, text string) ([]A, Event) {
	fmt.Println("In word.Enqueue")
	queue := []A{}
	var y A
	var b bool
	for _, wp := range wps {
		s := text[wp.Start:wp.End] // string in utf-8
		runes := ([]rune)(s)
		wl := len(runes)
		fmt.Println("s:", s, "wlength wl", wl)
		c := runes[0]
		e := IDType(c)
		p := runeToCType(c)
		fmt.Println("p: ctype[firstchar]", p)
		if wl > 1 {
			d := ESCType(runes[len(runes)-1])
			fmt.Println("d last char", d)
			if b = p != C9 && d == CESC1 || d == CESC2; b {
				e = spellIn[s]
				fmt.Println("b is true, e:", e)
			}
		}
		if y = cid2pdef(e); y.Type != NoAType {
			y = id2pdef[e]
			fmt.Println("c<128, y=", y)
		}
		y, ok := id2pdef[e]
		if ok {
			fmt.Println("c<128, e=", e, "ok", ok, "y=id2pdef[e]", y)
			queue = append(queue, y)
		} else if b {
			fmt.Println("UNEXPECTED b?")
		} else {
			switch p {
			case C9:
				x, ok := connum(s)
				if !ok {
					return queue, 0
				}
				queue = append(queue, x)
			case CQ:
				x, err := constr(s)
				if err != 0 {
					return queue, err
				}
				queue = append(queue, x)
			default:
				jsignal2(EVSPELL, wp)
				return queue, EVSPELL
			}
		}
	}
	return queue, 0
}
