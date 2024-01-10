package utils

// go get golang.org/x/crypto/sha3

import (
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func HashSha3_256(input string) (string, error) {
	hasher := sha3.New256()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	// 16진수 문자열로 변환
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

func HashSha3_512(input string) (string, error) {
	hasher := sha3.New512()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	// 16진수 문자열로 변환
	hashString := hex.EncodeToString(hashBytes)

	return hashString, nil
}

func Base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}

func Base64Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
