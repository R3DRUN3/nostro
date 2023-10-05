package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nostro",
	Short: "Nostr OSINT",
	Long:  `NostrO enables you to do Open Source Intelligence on the Nostr protocol`,
	Run: func(cmd *cobra.Command, args []string) {
		// Your CLI's main logic goes here
		fmt.Println("################### Welcome to NostrO 🔎 𓅦 ###################")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
