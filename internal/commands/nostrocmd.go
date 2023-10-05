package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nostro = &cobra.Command{
	Use:   "nostro",
	Short: "A custom command for NostrO",
	Long:  `This is a custom command for NostrO.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Your custom command logic goes here
		fmt.Println("Executing nostro...")
	},
}

func init() {
	// You can add flags and options specific to this command here
	// Example:
	// nostro.Flags().StringP("flagname", "f", "default", "help message")
	// Cobra also supports persistent flags and local flags
	// See https://github.com/spf13/cobra for more options
}

func Getnostro() *cobra.Command {
	return nostro
}
