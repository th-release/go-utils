package utils

import (
	"os"
	"path/filepath"
)

func CreateDirectoryIfNotExist(filePath string) error {
	dir := filepath.Dir(GetFilePath("") + filePath) // 파일 경로로부터 디렉토리 경로 추출

	// 디렉토리가 이미 존재하는지 확인
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 디렉토리가 존재하지 않으면 생성
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			// 디렉토리 생성 중 오류 발생
			return err
		}
	}

	return nil
}
