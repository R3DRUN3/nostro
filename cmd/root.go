package cmd

import (
	"fmt"
	"os"

	"github.com/r3drun3/nostro/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nostro",
	Short: "Nostr OSINT.",
	Long:  `Welcome to NostrO 🔎 𓅦`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display a message indicating that a subcommand is required
		fmt.Println("################### Welcome to NostrO 🔎 𓅦 ###################")
		fmt.Println("Please specify a subcommand (e.g., 'relay').")
	},
}

func init() {
	rootCmd.AddCommand(commands.RelayCmd)
	rootCmd.AddCommand(commands.NotesCmd)
	rootCmd.AddCommand(commands.DirectMessagesCmd)
	rootCmd.AddCommand(commands.UserCmd)
	rootCmd.AddCommand(commands.EventsCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
