package utils

import (
	"encoding/json"
	"os"
)

func GetJson(filePath string) (interface{}, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// JSON 파싱
	var result interface{}
	err = json.Unmarshal(fileContent, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
