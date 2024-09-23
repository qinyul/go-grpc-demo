package utils

import (
	"fmt"
	"os"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println("FileExist:: error checking file: ", err)
		return false
	}

	return true
}

func CreateFile(fileName string, format string) error {

	if FileExist(fileName) {
		fmt.Printf("File %s is already exist, skip creating file \n", fileName)
		return nil
	}
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("error creating file: %s error: %s \n", fileName, err.Error())
	} else {
		if format == "json" {
			file.WriteString("[]")
		}
		fmt.Printf("File %s created succesfully \n", fileName)
	}
	defer file.Close()
	return err

}
