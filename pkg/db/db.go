package db

import (
	"fmt"
	"os"

	"github.com/qinyul/go-grpc-demo/internal/utils"
)

func InitDB() error {
	fmt.Println("Starting init DB")

	fmt.Println("Create DB directory")
	err := os.MkdirAll("data", 0777)
	if err != nil {
		fmt.Println("Failed to create DB Directory err: ", err)
		return err
	}
	fmt.Println("DB Directory Created")

	fmt.Println("Create item.json")
	err = utils.CreateFile("data/item.json", "json")

	fmt.Println("Finish init db")
	return err
}
