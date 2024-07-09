package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	contentTodo := getUserInput("Todo content: ")
	titleNote := getUserInput("Note title: ")
	contentNote := getUserInput("Note content: ")

	userTodo, errNewTodo := todo.New(contentTodo)
	userNote, errNewNote := note.New(titleNote, contentNote)

	if errNewTodo != nil || errNewNote != nil {
		fmt.Println("ERROR")
		fmt.Println(errNewTodo)
		fmt.Println(errNewNote)
	} else {
		err := saveData(userTodo)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
		err = saveData(userNote)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
	}
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

func getUserInput(message string) (text string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return
}
