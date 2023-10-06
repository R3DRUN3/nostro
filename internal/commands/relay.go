package commands

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr/nip11"
	"github.com/spf13/cobra"
)

var relayInfo bool

var RelayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Operations on relays",
	Long:  `Retrieve data on nostr relays`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if relayInfo {
			if len(args) != 1 {
				return fmt.Errorf("relay name is required")
			}
			relay := args[0]
			data, err := nip11.Fetch(context.Background(), relay)
			if err != nil {
				return err
			}
			fmt.Println("####################### RELAY INFO #######################")
			fmt.Println("NAME: ", data.Name)
			fmt.Println("DESCRIPTION: ", data.Description)
			fmt.Println("PUB KEY: ", data.PubKey)
			fmt.Println("CONTACT: ", data.Contact)
			fmt.Println("SUPPORTED NIPS: ", data.SupportedNIPs)
			fmt.Println("SOFTWARE: ", data.Software)
			fmt.Println("VERSION: ", data.Version)
			fmt.Println("LIMITATION: ", data.Limitation)
			fmt.Println("PAYMENTSURL: ", data.PaymentsURL)
			fmt.Println("##########################################################")
		} else {
			cmd.Help()
		}
		return nil
	},
}

func init() {
	RelayCmd.Flags().BoolVarP(&relayInfo, "info", "", false, "Retrieve relay information document (nip-11)")
}
