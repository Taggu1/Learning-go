package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileType = ".json"

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("note missing some details")
	}

	return Todo{
		Text: content,
	}, nil
}

func (n Todo) Display() {
	fmt.Println(n.Text)
}

func (n Todo) Save() error {
	fileName := "todo"

	jsonString, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName+fileType, jsonString, os.ModePerm)

}
