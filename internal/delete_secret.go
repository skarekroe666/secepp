package internal

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func DeleteSecret() {
	userInput := func(input string) error {
		if input == "" {
			return errors.New("please mention a file name")
		} else {
			return nil
		}
	}

	prompt := promptui.Prompt{
		Label:    "Enter Secret",
		Validate: userInput,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	var input Secret
	if err := DB.Where("hash = ?", result).First(&input).Error; err != nil {
		log.Fatal(err)
	}
	DB.Delete(&input)

	fmt.Println("Secret deleted successfully")
}
