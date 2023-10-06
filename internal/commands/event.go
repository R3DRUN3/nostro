package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/spf13/cobra"
)

type Event struct {
	ID        string     `json:"id"`
	PubKey    string     `json:"pubkey"`
	Content   string     `json:"content"`
	Kind      int        `json:"kind"`
	Tags      [][]string `json:"tags"`
	Sig       string     `json:"sig"`
	CreatedAt int64      `json:"created_at"`
}

var eventData bool

var EventsCmd = &cobra.Command{
	Use:   "event",
	Short: "Operations on events",
	Long:  `Search and retrieve events on nostr`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if eventData {
			if len(args) != 2 {
				return fmt.Errorf("user npbu key and relay name are required")
			}
			id := args[0]
			url := args[1]
			// connect to relay
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			relay, err := nostr.RelayConnect(ctx, url)
			if err != nil {
				panic(err)
			}
			// create filters
			var filters nostr.Filters
			if _, v, err := nip19.Decode(id); err == nil {
				filters = []nostr.Filter{{
					IDs: []string{v.(string)},
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

			fmt.Println("####################### EVENT INFO #######################")
			//fmt.Println(evs[0].Tags)
			var event Event
			eventBytes, err := json.Marshal(evs[0])
			if err != nil {
				// Handle the error
				panic("Error marshaling event")
			}
			if err := json.Unmarshal(eventBytes, &event); err != nil {
				panic("Error decoding user JSON")
			}
			fmt.Println("ID:", event.ID)
			fmt.Println("PubKey:", event.PubKey)
			fmt.Println("Kind:", event.Kind)
			fmt.Println("Created At:", event.CreatedAt)
			fmt.Println("Tags:", event.Tags)
			fmt.Println("Content:", event.Content)
			fmt.Println("Signature:", event.Sig)
			fmt.Println("##########################################################")
		} else {
			cmd.Help()
		}
		return nil
	},
}

func init() {
	EventsCmd.Flags().BoolVarP(&eventData, "info", "", false, "Retrieve the specified event body from the specified relay.")
}
