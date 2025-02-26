package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const fileType = ".json"

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
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

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	jsonString, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName+fileType, jsonString, os.ModePerm)

}
