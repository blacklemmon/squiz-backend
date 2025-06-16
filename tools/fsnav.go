package tools

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkingDir() string {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		fmt.Println("Couldn't get the current working directory")
		return ""
	} else {
		return wd
	}
}

func GetFileName(fullFilepath string) string {
	lastSlash := strings.LastIndex(fullFilepath, string(filepath.Separator))
	return fullFilepath[lastSlash+1:]
}

// If the file does not exist, it will return an empty string
func GetLocalFilePath(filename string, workingDir string, controlExists bool) string {
	var localFilePath string

	if !controlExists {
		localFilePath = filepath.Join(workingDir, filename)
	} else {
		localFilePath = filename
		var fileExists bool = FileExists(localFilePath)

		if !fileExists {
			//There is an error, the file doesn't exist

			localFilePath = filepath.Join(workingDir, filename) //We check in the current directory
			fileExists = FileExists(localFilePath)

			if !fileExists {
				//We tried to look for the file in the current working directory
				//and we still didn't find it, so we give up
				localFilePath = ""
			}
		}
	}

	return localFilePath
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !errors.Is(err, os.ErrNotExist)
}