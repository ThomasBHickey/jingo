// Code generated by "stringer -type=CBType"; DO NOT EDIT

package jingo

import "fmt"

const _CBType_name = "CXCSCACNCBC9CDCCCQ"

var _CBType_index = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}

func (i CBType) String() string {
	if i < 0 || i >= CBType(len(_CBType_index)-1) {
		return fmt.Sprintf("CBType(%d)", i)
	}
	return _CBType_name[_CBType_index[i]:_CBType_index[i+1]]
}
