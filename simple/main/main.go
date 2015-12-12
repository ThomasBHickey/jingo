// My attempt at a simple interpreter
// Copyright Thomas B. Hickey 2015
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type V interface{}
type AType int

const (
	Loc AType = iota
	Op
	Value
	Box
)

type A struct {
	Type   AType
	Length int
	Shape  []int
	Data   V
}
type vMonad func(x A) A
type vDyad func(x, y A) A

func ASize(shape []int) int {
	sz := 1
	for _, sp := range shape {
		sz *= sp
	}
	return sz
}
func mkA(atype AType, shape []int) (na A) {
	na.Type = atype
	na.Length = ASize(shape)
	na.Shape = shape
	if na.Length > 0 {
		na.Data = make([]int, na.Length)
	} else {
		na.Data = 0
	}
	return
}
func mkIntA(typ AType, i int) (na A) {
	na = mkA(typ, []int{})
	na.Data = i
	return
}

func miot(w A) (z A) {
	//fmt.Println("iot")
	//pr(w)
	if w.Type == Loc {
		fmt.Println("found location")
		w = st[w.Data.(int)]
		//fmt.Println("pulled from location", w.Data)
		//pr(w)
	}
	if len(w.Shape) == 0 {
		n := w.Data.(int)
		d := make([]int, n)
		for i := 0; i < n; i++ {
			d[i] = i
		}
		z = mkA(Value, []int{n})
		z.Data = d
		//fmt.Println("miot returning", z)
	} else {
		fmt.Println("miot expects single values")
	}
	return z
}
func diot(a, w A) (na A) {
	fmt.Println("diot not implemented")
	return
}
func dasgn(a, w A) (na A) {
	if a.Type != Loc {
		fmt.Println("asgn expected location in a")
		return
	}
	loc := a.Data.(int)
	if loc < 0 || loc >= 26 {
		fmt.Println("Location out of range")
	} else {
		st[loc] = w
		return w
	}
	return
}
func dplus(a, w A) (na A) {
	fmt.Println("dplus a", a, "dplus w", w)
	if a.Type == Loc {
		a = st[a.Data.(int)]
	}
	if len(a.Shape) == 0 && len(w.Shape) == 0 {
		return mkIntA(Value, a.Data.(int)+w.Data.(int))
	}
	if len(a.Shape) == 0 {
		av := a.Data.(int)
		nd := make([]int, w.Length)
		od := w.Data.([]int)
		for i := 0; i < w.Length; i++ {
			nd[i] = av + od[i]
		}
		na.Type = Value
		na.Shape = w.Shape
		na.Length = w.Length
		na.Data = nd
		return
	}
	fmt.Println("dplus not complete")
	return
}
func mplus(a A) (na A) {
	fmt.Println("mplus not yet implemented")
	return
}
func newLine() {
	fmt.Println()
}
func prInt(i int) {
	fmt.Print(i, " ")
}
func pr(w A) {
	//fmt.Println("Just called 'pr' on", w)
	//fmt.Println("shape", w.Shape)
	for _, d := range w.Shape {
		prInt(d)
	}
	if len(w.Shape) > 0 {
		newLine()
	}
	switch w.Type {
	case Loc:
		fmt.Print(st[w.Data.(int)])
	case Op:
		fmt.Print(string(vt[w.Data.(int)]))
	case Value:
		//fmt.Println("printing", w)
		if len(w.Shape) == 0 {
			prInt(w.Data.(int))
		} else {
			for i := 0; i < w.Length; i++ {
				prInt(w.Data.([]int)[i])
			}
		}
	case Box:
		fmt.Print("< ")
		for i := 0; i < w.Length; i++ {
			pr(w.Data.([]A)[i])
		}
	}
}

var vt = "=+{~<#,"
var vDyads = []vDyad{dasgn, dplus, nil, diot, nil, nil}
var vMonads = []vMonad{nil, mplus, nil, miot}
var st [26]A

func ex(e A) (z A) {
	//fmt.Println("ex", e)
	switch e.Type {
	case Loc:
		return st[e.Data.(int)]
	case Op:
		fmt.Println("Found Op unexpectedly")
		return
	case Value:
		return e
	case Box:
		if e.Length == 0 {
			fmt.Println("ex:Empty box")
			return
		} else if e.Length == 1 {
			return ex(e.Data.([]A)[0])
		} else {
			a := e.Data.([]A)[0]
			b := e.Data.([]A)[1]
			if a.Type == Op { // monad
				rest := mkA(Box, []int{e.Length - 1})
				rest.Data = e.Data.([]A)[1:]
				return vMonads[a.Data.(int)](ex(rest))
			}
			if b.Type == Op {
				rest := mkA(Box, []int{e.Length - 2})
				rest.Data = e.Data.([]A)[2:]
				return vDyads[b.Data.(int)](a, ex(rest))
			}
		}
		fmt.Println("Don't know what to execute")
		return
	default:
		fmt.Println("Unexpected Type", e.Type)
	}
	return
}
func verbPos(ct byte) (pos int, ok bool) {
	pos = strings.IndexByte(vt, ct)
	if pos < 0 {
		return 0, false
	}
	return pos, true
}
func mkNoun(c byte) (z A, ok bool) {
	if c < '0' || c > '9' {
		return z, false
	}
	return mkIntA(Value, int(c-'0')), true
}
func words(s string) (z A) {
	//fmt.Println("just called words")
	n := len(s)
	e := make([]A, n)
	for i := 0; i < n; i++ {
		c := s[i]
		//fmt.Println("looking at", c, string(c))
		if a, ok := mkNoun(c); ok {
			e[i] = a
			//fmt.Println("wordsA", e[i])
		} else if a, ok := verbPos(c); ok {
			e[i] = mkIntA(Op, a)
			//fmt.Println("wordsB", e[i])
		} else if c >= 'a' && c <= 'z' {
			e[i] = mkIntA(Loc, int(c-'a'))
		} else {
			e[i] = mkIntA(Value, int(c))
			//fmt.Println("wordsC", e[i])
		}
	}
	z.Type = Box
	z.Shape = []int{n}
	z.Length = n
	z.Data = e
	//fmt.Println("wordsDone", z)
	return
}
func getString(reader *bufio.Reader) string {
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
func main() {
	//fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	for {
		//w := words(getString(reader))
		// w := words("~4")
		// fmt.Println("words:", w)
		// pr(w)
		// res := ex(w)
		// //fmt.Println("Result:", res)
		// pr(res)
		// newLine()
		// w = words("d=8")
		// fmt.Println("words:", w)
		// res = ex(w)
		// fmt.Println("Result:", res, st)
		// pr(res)
		// w = words("d")
		// res = ex(w)
		// fmt.Println("Result of d:", res)
		// pr(res)
		// break
		s := getString(reader)
		if s == "quit" || s == "exit" {
			break
		}
		//if w=="quit" || w=="exit"{break}
		pr(ex(words(s)))
		newLine()
	}
}
