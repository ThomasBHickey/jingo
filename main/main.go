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
	jt := jingo.GetJ()
	//jt.Log = log.New(os.Stderr, "", 0)
	jt.Log.Println("Starting in main")
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		ttext := strings.TrimSpace(text)
		if ttext == "quit" {
			break
		}
		wps := jingo.Scan(jt, ttext)
		fmt.Println("wps")
		for _, w := range wps {
			fmt.Println(text[w.Start:w.End])
		}
		q, event := jingo.Enqueue(jt, wps, ttext)
		fmt.Println("after Enqueue: jt.Symb[\"a\"]", jt.Symb["a"])
		if event != 0 {
			fmt.Println("enqueue failed", event)
		} else {
			fmt.Println("enqueue", q)
			z, err := jingo.Parse(jt, q)
			fmt.Println("after Parse: jt.Symb[\"a\"]", jt.Symb["a"])
			if jt.Asgn {
				fmt.Println("Assignment")
			} else {
				fmt.Println("result of Parse", err, z)
			}
		}
	}
}
