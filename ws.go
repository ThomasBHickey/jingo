package jingo

import (
	"fmt"
)
var spellt = [3][3]byte{
	{'a', 'a', 'a'},
	{'a', 'a', 'a'},
	{'a', 'a', 'a'}}

var spell = [3][70]byte{
	{'=', '<', '>', '_', '+', '*', '-', '%',
		'^', '$', '~', '|', '.', ':', ',', ';',
		'#', '@', '/', CBSLASH, '[', ']', '{', '}',
		'`', CQQ, '&', '!', '?', 'a', 'A', 'b',
		'c', 'C', 'd', 'D', 'e', 'E', 'f', 'H',
		'i', 'I', 'j', 'L', 'm', 'M', 'n', 'o',
		'p', 'q', 'r', 's', 'S', 't', 'T', 'u',
		'v', 'x', 'y', '0', '1', '2', '3', '4',
		'5', '6', '7', '8', '9', 0},
	{CASGN, CFLOOR, CCEIL, 1, CPLUSDOT, CSTARDOT, CNOT, CDOMINO,
		CLOG, CSPARSE, CNUB, CREV, CEVEN, COBVERSE, CCOMDOT, CCUT,
		CBASE, CATDOT, CSLDOT, CBSDOT, CLEV, CDEX, CTAKE, CDROP,
		CGRDOT, CEXEC, CUNDER, CFIT, CQRYDOT, CALP, CATOMIC, CBDOT,
		CCDOT, CCYCLE, CDDOT, CDCAP, CEPS, CEBAR, CFIX, CHGEOM,
		CIOTA, CICAP, CJDOT, CLDOT, CMDOT, CMCAP, CNDOT, CCIRCLE,
		CPOLY, 1, CRDOT, 1, 1, CTDOT, CTCAP, CUDOT,
		CVDOT, CXDOT, CYDOT, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 0},
	{CGASGN, CLE, CGE, CFCONS, CPLUSCO, CSTARCO, CMATCH, CROOT,
		CPOWOP, CSELF, CNE, CCANT, CODD, CADVERSE, CLAMIN, CWORDS,
		CABASE, CATCO, CGRADE, CDGRADE, CCAP, CIDA, CTAIL, CCTAIL,
		CGRCO, CTHORN, CAMPCO, CIBEAM, CQRYCO, CACE, 1, 1,
		1, 1, 1, CDCAPCO, 1, 1, 1, 1,
		CICO, 1, 1, CLCAPCO, 1, 1, 1, 1,
		CPCO, CQCO, 1, CSCO, CSCAPCO, CTCO, 1, CUCO,
		1, CXCO, 1, CFCONS, CFCONS, CFCONS, CFCONS, CFCONS,
		CFCONS, CFCONS, CFCONS, CFCONS, CFCONS, 0}}

var mapS2I = map[string] int{}

// func init() {
// 	for _, row := range(spell){
// 		fmt.Println("row ", row)
// 	}
// }
const (
	a = 1<<iota
	b
	c)
func init() {
	fmt.Println("a,b,c", a, b, c)
}
func init() {
	var r rune
	runes := make([]rune, 2)
	for i:=0; i<70; i++ {
		r = rune(spell[0][i])
		mapS2I[string(r)]=int(r)
		runes[0] = r
		runes[1] = '.'
		mapS2I[string(runes)] = int(spell[1][i])
		runes[1] = ':'
		mapS2I[string(runes)] = int(spell[2][i])
	}
	fmt.Println(len(mapS2I))
	fmt.Println("= :", mapS2I["="], mapS2I["=."], mapS2I["=:"])
	//fmt.Println(mapS2I)
}
