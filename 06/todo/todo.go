package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	content string
}

func New(content string) (*Todo, error) {
	if content == "" {
		return nil, errors.New("text is required")
	}

	return &Todo{
		content: content,
	}, nil
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	marshalledJSON, errMarshalJSON := t.marshalJSON()
	if errMarshalJSON != nil {
		return errors.New("marshal json failed")
	}

	errWriteFile := os.WriteFile(fileName, []byte(marshalledJSON), 0644)
	if errWriteFile != nil {
		return errWriteFile
	}
	return nil
}

func (t *Todo) Print() {
	fmt.Println("Todo saved!")
	fmt.Println("Content: ", t.content)
}

func (t *Todo) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Text string `json:"text"`
	}{
		Text: t.content,
	})
}
