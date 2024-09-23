package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSONFile(filePath string, data interface{}) error {
	fmt.Println("WriteJSONFile:: Starting writting json data:", data)

	file, err := os.Create(filePath)
	fmt.Println("WriteJSONFile:: file created")
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
		return err
	}
	defer file.Close()

	fmt.Println("WriteJSONFile:: start marshalling  data")
	jsonData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Printf("WriteJSONFile:: error marshaling data: %v\n", err)
		return err
	}

	fmt.Println("WriteJSONFile:: start writing to file  data\n", string(jsonData))

	_, err = file.WriteString(string(jsonData))

	if err != nil {
		fmt.Println("WriteJSONFile:: error writing file, error: ", err)
		return err
	}

	err = file.Sync()
	if err != nil {
		fmt.Println("Error syncing file: ", err)
	}

	fmt.Println("WriteJSONFile:: finish writting  json to file")

	return nil
}
