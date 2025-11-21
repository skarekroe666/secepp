/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/skarekroe666/secepp/cmd"
	"github.com/skarekroe666/secepp/internal"
)

func main() {
	// Initialize DB before executing commands so commands can use DB
	internal.InitDB()
	cmd.Execute()
}
