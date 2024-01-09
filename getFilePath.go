package utils

import (
	"os"
	"path/filepath"
)

func GetFilePath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(dir, "public", file)
}

/*
example:
func main() {
	fmt.Print(getFilePath("hi.json"))
	// /Users/cth/cth.release/personal/utils/public/hi.json
	// (pwd + public + {name}.json)
}
*/
