package main

import (
	"fmt"
	"github.com/felipeoli7eira/go-note/note"
	"bufio"
	"strings"
	"os"
)

func main() {
	title, content := getNoteData()

	n, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	n.Display()
	saveErr := n.Save()

	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := input("title: ")
	content := input("content: ")

	return title, content
}

func input(label string) string {
	fmt.Print(label)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
