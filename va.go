// Copyright 2016 Thomas B. Hickey
// Use of this code is goverened by
// license that can be found in the LICENSE file
package jingo

type CVType int

const (
	NoCVType CVType = 0
	VBB      CVType = 1 << iota //(I)1        /* convert arguments to B              */
	VII                         //(I)2        /* convert arguments to I              */
	VDD                         //(I)4        /* convert arguments to D              */
	VZZ                         //(I)8        /* convert arguments to Z              */
	VXX                         //(I)16       /* convert arguments to XNUM           */
	VQQ                         //(I)32       /* convert arguments to RAT            */
	VB                          //(I)256      /* result type B                       */
	VI                          //(I)512      /* result type I                       */
	VD                          //(I)1024     /* result type D                       */
	VZ                          //(I)2048     /* result type Z                       */
	VX                          //(I)4096     /* result type XNUM                    */
	VQ                          //(I)8192     /* result type RAT                     */
	VSB                         //(I)16384    /* result type SBT                     */
	VRD                         //(I)65536    /* convert result to D if possible     */
	VRI                         //(I)131072   /* convert result to I if possible     */
	VXEQ                        //(I)262144   /* convert to XNUM for = ~:            */
	VXCF                        //(I)524288   /* convert to XNUM ceiling/floor       */
	VXFC                        //(I)1048576  /* convert to XNUM floor/ceiling       */
)
