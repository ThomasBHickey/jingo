package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ThomasBHickey/jingo"
)


func main() {
	//fmt.Println("In Main")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	ttext := strings.TrimSpace(text)
	wps :=jingo.Scan(ttext)
	fmt.Println("wps")
	for _, w := range(wps){
		fmt.Println(text[w.Start:w.End])
		}
	jingo.Enqueue(wps, ttext)
}
