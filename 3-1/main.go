package main

import (
	"fmt"
	"golang_course_one/3-1/cmd"

	"github.com/spf13/cobra"
)

func main() {

	var input string
	command := &cobra.Command{
		Use:   "main",
		Short: "get string in command line",
		Run: func(command *cobra.Command, args []string) {
			cmd.Run(input)
		},
	}
	fmt.Println("Calling command.Execute()!")

	// Add a flag to the root command
	command.Flags().StringVarP(&input, "input", "i", "", "Input data")

	command.Execute()
}
