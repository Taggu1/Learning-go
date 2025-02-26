package main

import (
	"fmt"

	"example.org/noto/Day-4-note/note"
	"example.org/noto/Day-4-note/utility"
)

func main() {
	noteTitle := utility.Input("Note title: ")
	noteContent := utility.Input("Note content: ")

	note, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	err = note.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Note saved successfully")
}
