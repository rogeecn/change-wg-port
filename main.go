/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"
	"wg/cmd"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "server")
	}
	cmd.Execute()
}
