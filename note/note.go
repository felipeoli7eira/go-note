package note

import (
	"errors"
	"time"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

// struct tags

type Note struct {
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("[NOTE DATA]\n\n[TITLE]: %v\n\n[CONTENT]: %v\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ToLower(
		strings.ReplaceAll(n.Title, " ", "_")) + ".json"

	fileContent, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileContent, 0644)
}
