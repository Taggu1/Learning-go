package main

import (
	"fmt"

	"example.org/noto/Day-4-note/note"
	"example.org/noto/Day-4-note/utility"
	"example.org/noto/Day-5/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	note, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := getTodoData()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)

	printSomething(1)
	printSomething("Nice")
	printSomething(note)

}

func printSomething(value interface{}) {
	typedValue, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", typedValue)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the Todo failed")
		return err
	}

	fmt.Println("Note saved successfully")
	return nil
}

func getTodoData() (todo.Todo, error) {
	todoContent := utility.Input("Todo: ")

	return todo.New(todoContent)
}

func getNoteData() (note.Note, error) {
	noteTitle := utility.Input("Note title: ")
	noteContent := utility.Input("Note content: ")

	return note.New(noteTitle, noteContent)
}
