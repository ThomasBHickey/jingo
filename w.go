// Copyright Thomas B. Hickey 2015
// See license.txt in github.com/ThomasBHickey/jingo
package jingo
import (
"fmt"
)
func constr(s string) (A, EventType) {
	fmt.Println("In constr with", s)
	return A{}, 0
}
