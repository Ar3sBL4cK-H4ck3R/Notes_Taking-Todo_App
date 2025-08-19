package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes_app/notes"
	"example.com/notes_app/todo"
)

type storageFiler interface {
	StorageFile() error
}

type outputable interface {
	storageFiler
	Display()
}

func main() {
	inputTitle := noted("Notes Title: ")
	inputContent := noted("Notes Content: ")
	inputTodo := noted("Todo Text: ")

	todo, err := todo.New(&inputTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(todo)
	output, err := notes.New(&inputTitle, &inputContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(output)
}

func noted(userprompt string) string {
	fmt.Print(userprompt)
	value := bufio.NewReader(os.Stdin)
	text, err := value.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data storageFiler) error {
	err := data.StorageFile()
	if err != nil {
		fmt.Println("Saving The Note Failed!")
		return err
	}
	fmt.Println("Saving The Note Succeeded!")
	return nil
}

func outputData(data outputable) {
	data.Display()
	err := saveData(data)
	if err != nil {
		return
	}
}
