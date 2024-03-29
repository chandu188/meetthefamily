package main

import (
	"log"
	"os"
	"strings"

	"github.com/chandu188/familytree"

	"github.com/chandu188/platform"
)

func main() {

	if len(os.Args) != 2 {
		log.Println("missing file path information")
		os.Exit(1)
	}
	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	p := platform.NewPlatform(file, os.Stdout)
	commands := make([]string, 0)
	for {
		in, err := p.RetrieveInput()
		if err != nil {
			break
		}
		commands = append(commands, in)
	}

	ft := familytree.NewFamilyTree()

	for _, v := range commands {
		tokens := strings.Split(v, " ")
		var res string
		if len(tokens) == 4 && tokens[0] == "ADD_CHILD" {
			res = ft.AddChild(tokens[1], tokens[2], tokens[3])
		} else if len(tokens) == 3 && tokens[0] == "GET_RELATIONSHIP" {
			res = ft.GetRelationShip(tokens[1], tokens[2])
		}
		p.WriteOutput(res)
	}

}
