package utils

import (
	"encoding/json"
	"os"
)

func GetJsonFile(filePath string) (interface{}, error) {
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

func JsonToMap(data string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func JsonToArray(data string) ([]interface{}, error) {
	var result []interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
