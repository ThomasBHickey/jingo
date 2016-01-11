// Copyright 2015 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

import (
//"fmt"
)

func plus(jt *J, a, w A) (z A, evn Event) {
	jt.Log.Println("In plus")
	return va2(jt, a, w, CPLUS)
}

func va2(jt *J, a, w A, id IDType) (z A, evn Event) {
	jt.Log.Println("In va2")
	//oq := jt.rank
	an, as :=a.Length, a.Shape
	var at, wt AType
	if an==0 { at = B01}else{at = a.Type}
	ar:= len(as)
	wn, ws := w.Length, w.Shape
	if wn==0 {wt = B01}else{wt = a.Type}
	wr := len(w.Shape)
	jt.Log.Println("an, ar, at, aw", an, ar, at, as)
	jt.Log.Println("wn, wr, ws, wt", wn, wr, wt, ws)
	return z, EVSYSTEM
}
