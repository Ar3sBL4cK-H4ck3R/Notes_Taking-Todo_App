package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notes struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"creat_at"`
}

func New(noteTitle, noteContent *string) (*Notes, error) {
	if *noteTitle == "" || *noteContent == "" {
		return nil, errors.New("empty file error")
	}
	return &Notes{
		Title:     *noteTitle,
		Content:   *noteContent,
		CreatedAt: time.Now(),
	}, nil
}

func (note Notes) StorageFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func (note *Notes) Display() {
	_, err := fmt.Printf("\nYour Notes of Title \"%s\" have the Content Given Below:-\n\n%s\n\n", note.Title, note.Content)
	if err != nil {
		return
	}
}
