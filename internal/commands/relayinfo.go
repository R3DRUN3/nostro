// package commands

// import (
// 	"context"
// 	"fmt"

// 	"github.com/nbd-wtf/go-nostr/nip11"
// 	"github.com/spf13/cobra"
// )

// var infoCmd = &cobra.Command{
// 	Use:   "info [RELAY]",
// 	Short: "Retrieve relay information document (nip-11)",
// 	Long:  `The info subcommand retrieves the relay capabilities, administrative contacts, and various server attributes.`,
// 	Args:  cobra.ExactArgs(1), // Expects exactly 1 argument (relay name)
// 	Run: func(cmd *cobra.Command, args []string) {
// 		url := args[0]
// 		data, err := nip11.Fetch(context.Background(), url)

// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(data)

// 	},
// }

// func init() {
// 	// Add the "info" subcommand to the "relay" command
// 	relayCmd.AddCommand(infoCmd)
// }

package commands

import (
	"context"
	"fmt"

	"github.com/nbd-wtf/go-nostr/nip11"
	"github.com/spf13/cobra"
)

var relayInfoCmd = &cobra.Command{
	Use:   "info [RELAY]",
	Short: "Retrieve relay information document (nip-11)",
	Long:  `The info subcommand retrieves the relay capabilities, administrative contacts, and various server attributes.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		relay := args[0]
		data, err := nip11.Fetch(context.Background(), relay)
		if err != nil {
			panic(err)
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

	},
}

func init() {
	RelayCmd.AddCommand(relayInfoCmd)
}
