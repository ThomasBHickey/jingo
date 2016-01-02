package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ThomasBHickey/jingo"
	"os"
	"strings"
)

func getString(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	ttext := strings.TrimSpace(text)
	if ttext == "quit" {
		return "", errors.New("quit")
	}
	return ttext, nil
}

func prompt(prompt string) {
	fmt.Print(prompt)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	jt := jingo.GetJ()
	jt.Log.Println("Starting in jconsole")
	for {
		prompt(">")
		text, err := getString(reader)
		if err != nil {
			break
		}
		jingo.JDo(jt, text)
	}
	jt.Log.Println("Finishing jconsole")
}
