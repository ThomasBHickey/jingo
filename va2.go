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

// I atype(I cv){
//  if(!(cv&VBB+VII+VDD+VZZ+VQQ+VXX+VXEQ+VXCF+VXFC))R 0;
//  R cv&VBB?B01:cv&VII?INT:cv&VDD?FL:cv&VZZ?CMPX:cv&VQQ?RAT:XNUM;
// }    /* argument conversion */

func atype(cv CVType) AType {
	if (cv & (VBB + VII + VDD + VZZ + VQQ + VXX + VXEQ + VXCF + VXFC)) == 0 {
		return 0
	}
	if (cv & VBB) != 0 {
		return B01
	}
	if (cv & VII) != 0 {
		return INT
	}
	if (cv & VDD) != 0 {
		return FL
	}
	if (cv & VZZ) != 0 {
		return CMPX
	}
	if (cv & VQQ) != 0 {
		return RAT
	}
	return XNUM
}

// I rtype(I cv){R cv&VB?B01:cv&VI?INT:cv&VD?FL:cv&VZ?CMPX:cv&VQ?RAT:cv&VX?XNUM:SBT;}
//      /* result type */
func rtype(cv CVType) AType {
	if (cv & VB) != 0 {
		return B01
	}
	if (cv & VI) != 0 {
		return INT
	}
	if (cv & VD) != 0 {
		return FL
	}
	if (cv & VZ) != 0 {
		return CMPX
	}
	if (cv & VQ) != 0 {
		return RAT
	}
	if (cv & VX) != 0 {
		return XNUM
	}
	return SBT
}
func varr(jt *J, id IDType, a, w A, at, wt AType, ado *dyadFunct, I *CVType) (b bool) {
	return
}
func va2(jt *J, a, w A, id IDType) (z A, evn Event) {
	jt.Log.Println("In va2")
	//oq := jt.rank
	sp := false
	an, as := a.Length, a.Shape
	var at, wt AType
	if an == 0 {
		at = B01
	} else {
		at = a.Type
	}
	ar := len(as)
	wn, ws := w.Length, w.Shape
	if wn == 0 {
		wt = B01
	} else {
		wt = a.Type
	}
	wr := len(w.Shape)
	jt.Log.Println("an, ar, at, aw", an, ar, at, as)
	jt.Log.Println("wn, wr, ws, wt", wn, wr, wt, ws)
	if id == CEXP {
		return z, EVSYSTEM
	} // needs logic for sqroot
	if (SPARSE & at) != 0 {
		sp = true
		at = DTYPE(at)
	}
	if (SPARSE & wt) != 0 {
		sp = true
		wt = DTYPE(at)
	}
	var ado dyadFunct
	var cv CVType
	b := varr(jt, id, a, w, at, wt, &ado, &cv)
	if !b {
		return
	}
	zt := rtype(cv)
	t := atype(cv)
	if (t != 0) && !sp {
		b = (t & XNUM) != 0 // b = 1&&t&XNUM
		if t != at {
			if b {
				var param XMode
				if (cv & VXEQ) != 0 {
					param = XMEXMT
				} else if (cv & VXFC) != 0 {
					param = XMFLR
				} else if (cv & VXCF)!=0 {
					param = XMCEIL
				} else {
					param = XMEXACT
				}
				a = xcvt(param, a)
			} else {
				a = cvt(t, a)
			}
		}
	}
	return z, EVSYSTEM
}
