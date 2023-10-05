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
	Long:  `NostrO enables you to do Open Source Intelligence on the Nostr protocol.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display a message indicating that a subcommand is required
		fmt.Println("################### Welcome to NostrO 🔎 𓅦 ###################")
		fmt.Println("Please specify a subcommand (e.g., 'relay').")
	},
}

func init() {
	rootCmd.AddCommand(commands.RelayCmd)
	rootCmd.AddCommand(commands.NotesCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
