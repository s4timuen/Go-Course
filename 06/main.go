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
		err := outputData(userTodo)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
		err = outputData(userNote)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
		}
		printSomething(1)
		printSomething(1.5)
		printSomething("test")
	}
}

func outputData(data Outputtable) error {
	data.Print()
	return data.Save()
}

func printSomething(value any) { // same as interface{}

	typedValue, isInt := value.(int)
	fmt.Println(typedValue, "is an integer: ", isInt)

	switch value.(type) { // works only in switch
	case int:
		fmt.Println("Int: ", value)
	case string:
		fmt.Println("String: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	default:
		fmt.Println("Other type: ", value)
	}
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
