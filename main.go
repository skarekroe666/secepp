/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/skarekroe666/secepp/cmd"
	"github.com/skarekroe666/secepp/internal"
)

func main() {
	cmd.Execute()
	internal.InitDB()
}
