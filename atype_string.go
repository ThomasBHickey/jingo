// Code generated by "stringer -type=AType"; DO NOT EDIT

package jingo

import "fmt"

const _AType_name = "B01LITINTFLCMPXBOXXNUMRATBITSB01SLITSINTSFLSCMPXSBOXSBTC2TVERBADVCONJASGNMARKSYMBCONWNAMELPAREDGERPARXDXZAVNCAVNANY"

var _AType_map = map[AType]string{
	1:          _AType_name[0:3],
	2:          _AType_name[3:6],
	4:          _AType_name[6:9],
	8:          _AType_name[9:11],
	16:         _AType_name[11:15],
	32:         _AType_name[15:18],
	64:         _AType_name[18:22],
	128:        _AType_name[22:25],
	256:        _AType_name[25:28],
	512:        _AType_name[28:32],
	1024:       _AType_name[32:36],
	2048:       _AType_name[36:40],
	4096:       _AType_name[40:43],
	8192:       _AType_name[43:48],
	16384:      _AType_name[48:52],
	32768:      _AType_name[52:55],
	65536:      _AType_name[55:58],
	131072:     _AType_name[58:62],
	262144:     _AType_name[62:65],
	524288:     _AType_name[65:69],
	1048576:    _AType_name[69:73],
	2097152:    _AType_name[73:77],
	4194304:    _AType_name[77:81],
	8388608:    _AType_name[81:85],
	16777216:   _AType_name[85:89],
	33554432:   _AType_name[89:93],
	36700160:   _AType_name[93:97],
	67108864:   _AType_name[97:101],
	134217728:  _AType_name[101:103],
	268435456:  _AType_name[103:105],
	403177471:  _AType_name[105:108],
	403701759:  _AType_name[108:112],
	4294967295: _AType_name[112:115],
}

func (i AType) String() string {
	if str, ok := _AType_map[i]; ok {
		return str
	}
	return fmt.Sprintf("AType(%d)", i)
}