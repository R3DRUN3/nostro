package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/spf13/cobra"
)

var userReceivedDm bool

var DirectMessagesCmd = &cobra.Command{
	Use:   "dm",
	Short: "Operations on direct messages",
	Long:  `Search and retrieve direct messages on nostr`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userReceivedDm {
			if len(args) != 2 {
				return fmt.Errorf("user npbu key and relay name are required")
			}
			npub := args[0]
			url := args[1]
			// connect to relay
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			relay, err := nostr.RelayConnect(ctx, url)
			if err != nil {
				panic(err)
			}
			// create filters
			var filters nostr.Filters
			if _, v, err := nip19.Decode(npub); err == nil {
				t := make(map[string][]string)
				// making a "p" tag for the above public key.
				// this filters for messages tagged with the user, mainly replies.
				t["p"] = []string{v.(string)}
				filters = []nostr.Filter{{
					Kinds: []int{int(nostr.KindEncryptedDirectMessage)},
					Tags:  t,
					Limit: 300,
				}}
			} else {
				panic("not a valid npub!")
			}
			// create a subscription and submit to relay
			// results will be returned on the sub.Events channel
			sub, _ := relay.Subscribe(ctx, filters)

			// we will append the returned events to this slice
			evs := make([]nostr.Event, 0)

			go func() {
				<-sub.EndOfStoredEvents
				cancel()
			}()
			for ev := range sub.Events {
				evs = append(evs, *ev)
			}

			filename := "user_received_direct_messages.json"
			if f, err := os.Create(filename); err == nil {
				fmt.Fprintf(os.Stderr, "returned events saved to %s\n", filename)
				// encode the returned events in a file
				enc := json.NewEncoder(f)
				enc.SetIndent("", " ")
				enc.Encode(evs)
				f.Close()
			} else {
				panic(err)
			}
		} else {
			cmd.Help()
		}
		return nil
	},
}

func init() {
	DirectMessagesCmd.Flags().BoolVarP(&userReceivedDm, "userreceived", "", false, "Retrieve from the specified relay the last 300 direct messages that the specified user received.")
}
