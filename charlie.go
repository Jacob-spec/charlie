package main

import (
	"fmt"
	"io/ioutil"
	"github.org/Jacob-spec/charlie/frontend"
)

func readFile(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
	}

	return string(contents)
}


func main() {
	contents := readFile("/home/jake/Go/markdown/misc/charles.note")
	tokens := frontend.Lex(contents)
	frontend.PrintTokens(tokens)
}
