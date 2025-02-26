package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const jsonFileName = "database.json"

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time `json:"CreatedAt"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("note missing some details")
	}

	return Note{
		Title:   title,
		Content: content,
	}, nil
}

func (n Note) Display() {
	fmt.Printf("Your note titled %s has the following content: \n\n%s\n\n", n.Title, n.Content)
}

func (n Note) Save() {
	jsonString, err := json.Marshal(n)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(jsonFileName, jsonString, os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note saved succesfully")
}
