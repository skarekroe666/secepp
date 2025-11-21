package internal

import "fmt"

func ListAllFiles() {
	var secret []Secret

	DB.Find(&secret)

	for _, v := range secret {
		fmt.Printf("Id: %d\n", v.ID)
		fmt.Printf("Hash: %s\n", v.Hash[:6])
		fmt.Printf("Filename: %s\n", v.FileName)
		fmt.Printf("CreatedAt: %s\n", v.CreatedAt)
		fmt.Println()
	}
}
