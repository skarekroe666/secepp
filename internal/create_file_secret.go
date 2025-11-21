package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

/* UNCOMMENT THIS FUNCTION TO GET A SIMPLE CLI INPUT AND RESPONSE
NOTE: FOR THIS FUNCTION USE THE COMMAND AND ENTER FILE NAME AS THE SAME ARGS
AND UNCOMMENT THE FILENAME CHECK IN cmd/crete.go */

// func CreateSecret(fileName string) {
// 	filePath := "files/" + fileName
//
// 	_, err := os.Stat(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	hash, err := getFileHash(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Here's your secret file:", hash)
// }

// THIS FUNCTION USES A UI LIBRARY TO FOR AN INTERACTIVE EXPERIENCE
func CreateSecret() {
	userInput := func(input string) error {
		if input == "" {
			return errors.New("please mention file name")
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
	}

	filePath := "files/" + result

	hash, err := getFileHash(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Here's your secret:", hash)

	output := Secret{
		Hash:      hash,
		FileName:  result,
		CreatedAt: time.Now(),
	}

	fmt.Println()
	fmt.Printf("Here's your hash: %s\n", output.Hash)
	fmt.Printf("File name: %s\n", output.FileName)
	fmt.Printf("Created at: %s\n", output.CreatedAt.Format(time.UnixDate))
	DB.Create(&output)
}

func getFileHash(filePath string) (string, error) {
	//Use Open for reading large files
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("error hashing file: %v", err)
	}

	hexHash := hex.EncodeToString(hasher.Sum(nil))

	return hexHash, nil
}
