package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/ThomasBHickey/jingo"
)


func main() {
	fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	wps :=jingo.Scan(text)
	fmt.Println("wps", wps)
	for _, wp := range(wps){
		fmt.Println(wp, text[wp.Start:wp.End])
		}
}
