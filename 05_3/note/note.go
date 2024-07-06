package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title string, description string) (*Note, error) {
	if title == "" || description == "" {
		return nil, errors.New("title and description are required")
	}

	return &Note{
		title:     title,
		content:   description,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) Save() error {
	formatedTitle := strings.ReplaceAll(n.title, " ", "_")
	formatedTitle = strings.ToLower(formatedTitle)
	fileName := "todo_" + formatedTitle + ".json"

	marshalledJSON, errMarshalJSON := n.marshalJSON()
	if errMarshalJSON != nil {
		return errors.New("marshal json failed")
	}

	errWriteFile := os.WriteFile(fileName, []byte(marshalledJSON), 0644)
	if errWriteFile != nil {
		return errWriteFile
	}
	n.Print()
	return nil
}

func (n *Note) Print() {
	fmt.Println("Note saved!")
	fmt.Println("Title: ", n.title)
	fmt.Println("Content: ", n.content)
	fmt.Println("Created at: ", n.createdAt)
}

func (n *Note) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAt,
	})
}
