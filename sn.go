// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
	"fmt"
	"strings"
)

func valnm(jt *J, s string) bool {
	if len(s) == 0 {
		return false
	}
	fb := s[0] // first byte
	if jt.Dotnames && len(s) == 2 && s[1] == '.' && (fb == 'm' || fb == 'n' || fb == 'u' || fb == 'v' || fb == 'x' || fb == 'y') {
		fmt.Println("dotnames not supported")
		return false
	}
	//d := 'a'
	//b := false
	// j := 0
	for bp, c := range s {
		t := runeToCType(c)
		if bp == 0 && t != CA {
			return false // first position must be alpha
		}
		if t != CA && t != C9 {
			return false
		}
		// 	if c=='_'&&d=='_' && !b && bp!=len(s)-1{j=bp+1;b=true}
		// 	d:= c
	}
	if strings.ContainsRune(s, '_') {
		fmt.Println("underscore found in valnm")
		return false
	}
	return true
}

// Name from string
func nfs(jt *J, s string) (z A, ok bool) {
	ts := strings.Trim(s, " ")
	if len(ts) == 0 {
		return
	}
	z = NewNameArray(ts)
	//jt.Symb[s] = z
	//fmt.Println("in nfs (not implemented), passed", s)
	return z, true
}
