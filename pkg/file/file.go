package file

import (
	"fmt"
	"os"
)

func GetWorkingDirectory(logRef string) string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s Error getting working directory: %v\n", logRef, err)
	}

	return cwd
}
