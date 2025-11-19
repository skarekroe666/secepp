package cmd

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "sha for content",
	Run: func(cmd *cobra.Command, args []string) {
		Example()
	},
}

func Example() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Print("here is the hash: ")
	fmt.Printf("%x\n", h.Sum(nil))
}

func init() {
	rootCmd.AddCommand(hashCmd)
}
