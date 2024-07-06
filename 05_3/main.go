package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	userNote, errNewNote := note.New(title, content)

	if errNewNote != nil {
		fmt.Println("ERROR")
		fmt.Println(errNewNote)
	} else {
		errSaveNote := userNote.Save()
		if errSaveNote != nil {
			fmt.Println("ERROR")
			fmt.Println(errSaveNote)
		}
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
