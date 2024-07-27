package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var factPatStr = "\\[(.*?)\\] \\[(.*?)\\] \\[(.*?)\\]"

type Fact struct {
	Subj string
	Pred string
	Obj  string
}

func (f *Fact) String() string {
	return fmt.Sprintf("[[ %s :: %s :: %s ]]", f.Subj, f.Obj, f.Pred)
}

func NewFact(s, p, o string) *Fact {
	return &Fact{s, p, o}
}

type PO struct {
	Pred string
	Obj  string
}

func NewPO(p, o string) *PO {
	return &PO{p, o}
}

func main() {
	factPat := regexp.MustCompile(factPatStr)

	subjects := map[string][]*Fact{}
	//predicates := map[string]Fact{}
	//objects := map[string]Fact{}

	matches, err := filepath.Glob("*.md")
	check(err)
	if len(matches) == 0 {
		fmt.Println("No matches!")
	}
	for _, mdfile := range matches {
		//fmt.Printf("Match: %s\n", mdfile)
		mdtext, err := os.ReadFile(mdfile)
		facts := factPat.FindAllStringSubmatch(string(mdtext), -1)
		for _, f := range facts {
			s := f[1]
			p := f[2]
			o := f[3]
			fact := NewFact(s, p, o)
			subjects[s] = append(subjects[s], fact)
		}
		check(err)
	}

	for s, pos := range subjects {
		fmt.Printf("%s\n", s)
		for _, po := range pos {
			fmt.Printf("  %s :: %s\n", po.Pred, po.Obj)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
