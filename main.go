package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var factPatStr = "\\[(.*)\\] \\[(.*)\\] \\[(.*)\\]"

func main() {
	factPat := regexp.MustCompile(factPatStr)

	matches, err := filepath.Glob("*.md")
	check(err)
	if len(matches) == 0 {
		fmt.Println("No matches!")
	}
	for _, mdfile := range matches {
		//fmt.Printf("Match: %s\n", mdfile)
		mdtext, err := os.ReadFile(mdfile)
		facts := factPat.FindAllString(string(mdtext), -1)
		for _, fact := range facts {
			fmt.Printf("Fact: %s\n", fact)
		}
		check(err)
	}

}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
