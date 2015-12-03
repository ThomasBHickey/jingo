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
	snpdefs :=jingo.Scan(strings.TrimSpace(text))
	fmt.Println("snpdefs", snpdefs)
	for _, spd := range(snpdefs){
		fmt.Println(spd)
		}
	jingo.Enqueue(snpdefs)
}
