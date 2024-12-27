package main

import (
	"fmt"
	"github.com/felipeoli7eira/go-note/note"

)

func main() {
	title, content := getNoteData()

	n, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}

func getNoteData() (string, string) {
	title := input("title: ")
	content := input("content: ")

	return title, content
}

func input(label string) string {
	fmt.Println(label)
	var input string
	fmt.Scanln(&input)

	return input
}
