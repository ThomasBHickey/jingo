// Code generated by "stringer -type=AType"; DO NOT EDIT

package jingo

import "fmt"

const _AType_name = "NoATypeB01LITIS1BYTEINTFLCMPXBOXXNUMRATBITSB01SLITSINTSFLSCMPXSBOXSPARSESBTC2TJCHARDIRECTVERBADVCONJFUNCASGNMARKSYMBCONWNAMELAST0LPAREDGERPARXDXZNUMERICDENSENOUNAVNRHSANY"

var _AType_map = map[AType]string{
	0:          _AType_name[0:7],
	2:          _AType_name[7:10],
	4:          _AType_name[10:13],
	6:          _AType_name[13:20],
	8:          _AType_name[20:23],
	16:         _AType_name[23:25],
	32:         _AType_name[25:29],
	64:         _AType_name[29:32],
	128:        _AType_name[32:36],
	256:        _AType_name[36:39],
	512:        _AType_name[39:42],
	1024:       _AType_name[42:46],
	2048:       _AType_name[46:50],
	4096:       _AType_name[50:54],
	8192:       _AType_name[54:57],
	16384:      _AType_name[57:62],
	32768:      _AType_name[62:66],
	64512:      _AType_name[66:72],
	65536:      _AType_name[72:75],
	131072:     _AType_name[75:78],
	133124:     _AType_name[78:83],
	197182:     _AType_name[83:89],
	262144:     _AType_name[89:93],
	524288:     _AType_name[93:96],
	1048576:    _AType_name[96:100],
	1835008:    _AType_name[100:104],
	2097152:    _AType_name[104:108],
	4194304:    _AType_name[108:112],
	8388608:    _AType_name[112:116],
	16777216:   _AType_name[116:120],
	33554432:   _AType_name[120:124],
	33685510:   _AType_name[124:129],
	67108864:   _AType_name[129:133],
	73400320:   _AType_name[133:137],
	134217728:  _AType_name[137:141],
	268435456:  _AType_name[141:143],
	536870912:  _AType_name[143:145],
	805337018:  _AType_name[145:152],
	805503998:  _AType_name[152:157],
	805568510:  _AType_name[157:161],
	806354942:  _AType_name[161:164],
	807403518:  _AType_name[164:167],
	4294967295: _AType_name[167:170],
}

func (i AType) String() string {
	if str, ok := _AType_map[i]; ok {
		return str
	}
	return fmt.Sprintf("AType(%d)", i)
}
