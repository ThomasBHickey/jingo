package main

import (
	"bufio"
	"fmt"
	"github.com/ThomasBHickey/jingo"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	ttext := strings.TrimSpace(text)
	jt := jingo.GetJ()
	wps := jingo.Scan(jt, ttext)
	fmt.Println("wps")
	for _, w := range wps {
		fmt.Println(text[w.Start:w.End])
	}
	q, event := jingo.Enqueue(jt, wps, ttext)
	if event!=0 {
		fmt.Println("enqueue failed", event)
	} else {
		fmt.Println("enqueue", q)
		z, err := jingo.Parse(jt, q)
		fmt.Println("result of Parse", err, z)
	}
}
