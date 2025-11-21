package internal

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func GetFile() {
	userInput := func(input string) error {
		if input == "" {
			return errors.New("please mention a file name")
		} else {
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    "Enter file name",
		Validate: userInput,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	// input := strings.Split(result, ".")
	// fmt.Println(input)

	var inputSecret Secret
	if err := DB.Where("file_name = ?", result).First(&inputSecret).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Id: %d\n", inputSecret.ID)
	fmt.Printf("Hash: %s\n", inputSecret.Hash)
	fmt.Printf("Filename: %s\n", inputSecret.FileName)
	fmt.Println()
}
