package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("examples/epic.md")
	if err != nil {
		log.Fatal(err)
	}

	p := NewParser()

	tasks, _ := p.Parse(f)
	for _, t := range tasks {
		fmt.Println(t.ID, t.Title)
		for _, c := range t.Children {
			fmt.Println("--", c.ID, c.Title)
		}
	}
}
