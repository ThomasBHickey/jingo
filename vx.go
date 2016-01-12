// Copyright 2016 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

/* values for jt->xmode */
type XMode int

const (
	XMFLR   XMode = iota //       0                    /* floor,   round down        */
	XMCEIL               //		  1                    /* ceiling, round up          */
	XMEXACT              //       2                    /* exact, error if impossible */
	XMEXMT               //       3                    /* exact, empty if impossible */
	XMRND                //       4                    /* round,   round to nearest  */
)
