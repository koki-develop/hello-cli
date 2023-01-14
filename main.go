package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "hello-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, World")
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
