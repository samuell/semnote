package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, err := filepath.Glob("*.md")
	check(err)
	if len(matches) == 0 {
		fmt.Println("No matches!")
	}
	for _, m := range matches {
		fmt.Printf("Match: %s\n", m)
	}
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
