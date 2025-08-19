package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(todoContent *string) (*Todo, error) {
	if *todoContent == "" {
		return nil, errors.New("empty file error")
	}
	return &Todo{
		Text: *todoContent,
	}, nil
}

func (todo Todo) StorageFile() error {
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile("todo.json", json, 0644)

}

func (todo *Todo) Display() {
	_, err := fmt.Printf("\nYour Todo Text: %v\n", todo.Text)
	if err != nil {
		return
	}
}
