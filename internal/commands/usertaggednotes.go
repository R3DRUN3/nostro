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

var userTaggedNoteCmd = &cobra.Command{
	Use:   "usertagged [NPUB] [RELAY]",
	Short: "Retrieve notes in which a specific user has been tagged",
	Long:  `The usertagged subcommand retrieves from the specified relay the last 30 notes in which the specified user has been tagged.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
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
				Kinds: []int{nostr.KindTextNote},
				Tags:  t,
				Limit: 30,
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

		filename := "user_tagged_notes.json"
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

	},
}

func init() {
	NotesCmd.AddCommand(userTaggedNoteCmd)
}
