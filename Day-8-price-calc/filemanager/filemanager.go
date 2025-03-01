package filemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (manager FileManager) ReadLines() ([]string, error) {
	data, err := os.ReadFile(manager.InputFilePath)

	if err != nil {
		fmt.Println(err)
		return []string{}, errors.New("failed to read file")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	return lines, nil

}

func (manager FileManager) WriteResault(data interface{}) error {
	file, err := os.Create(manager.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to encode data")
	}

	file.Close()
	return nil

}

func New(inputFilePath, outPutFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outPutFilePath,
	}
}
