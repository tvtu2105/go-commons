package files

import (
	"github.com/tvtu2105/go-commons/utils"
	"os"
)

func CheckFolderExist(path string) bool {
	return checkExist(path)
}

func CheckFileExist(path string) bool {
	return checkExist(path)
}

func checkExist(path string) bool {
	if utils.IsEmpty(path) {
		return false
	}
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
