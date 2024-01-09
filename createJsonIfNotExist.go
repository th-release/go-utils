package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateJsonIfNotExist(filePath string, initialContent interface{}) error {
	_, err := os.Open(filePath)
	if err != nil {
		err := os.MkdirAll(GetFilePath(""), os.ModePerm)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("setup error: %v", err)
		}
		defer file.Close()

		if initialContent != nil {
			data, err := json.Marshal(initialContent)
			if err != nil {
				return fmt.Errorf("error marshaling initial content: %v", err)
			}

			_, err = file.Write(data)
			if err != nil {
				return fmt.Errorf("error writing initial content to file: %v", err)
			}
		}
	}

	return nil
}
