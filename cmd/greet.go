package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Prints a greeting",
	Long:  `This command prints a greeting to the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
