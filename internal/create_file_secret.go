package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func CreateSecret(fileName string) {
	filePath := "files/" + fileName

	_, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := getFileHash(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Here's your secret file:", hash)
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
