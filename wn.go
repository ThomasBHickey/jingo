package jingo

import (
	"fmt"
	//"os"
	"strconv"
	"strings"
)

func init() {
	s := "1"
	i, _ := strconv.Atoi(s)
	fmt.Println("converted string ", s, i)
}

func ToNumber(s string) (n int64, at AType, err error) {
	n, err = strconv.ParseInt(s, 10, 64)
	return n, INT, err
}

func PossibleNumber(s string) bool {
	return AllIn(s, "0123456789.ejrx_")
}

// check to see that all characters in s are in chars
func AllIn(s, chars string) bool {
	if len(chars) > 0 {
		for _, c := range s {
			if !strings.ContainsRune(chars, rune(c)) {
				return false
			}
		}
		return true
	}
	return false // no chars, not a number
}
