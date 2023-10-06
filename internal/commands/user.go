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

type User struct {
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Banner      string `json:"banner"`
	Website     string `json:"website"`
	About       string `json:"about"`
	Nip05       string `json:"nip05"`
	Lud16       string `json:"lud16"`
	Lud06       string `json:"lud06"`
	CreatedAt   int64  `json:"created_at"`
	Nip05Valid  bool   `json:"nip05valid"`
}

var userData bool

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Operations on users",
	Long:  `Search and retrieve users info`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if userData {
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
					Kinds:   []int{int(nostr.KindSetMetadata)},
					Authors: []string{v.(string)},
					//Tags:  t,
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
			fmt.Println("####################### USER INFO #######################")
			fmt.Println(evs[0].Tags)
			var user User
			if err := json.Unmarshal([]byte(evs[0].Content), &user); err != nil {
				panic("Error decoding user JSON")
			}
			fmt.Println("Name:", user.Name)
			fmt.Println("Picture:", user.Picture)
			fmt.Println("Username:", user.Username)
			fmt.Println("Display Name:", user.DisplayName)
			fmt.Println("Banner:", user.Banner)
			fmt.Println("Website:", user.Website)
			fmt.Println("About:", user.About)
			fmt.Println("Nip05:", user.Nip05)
			fmt.Println("Lud16:", user.Lud16)
			fmt.Println("Lud06:", user.Lud06)
			fmt.Println("Created At:", user.CreatedAt)
			fmt.Println("Nip05 Valid:", user.Nip05Valid)
			fmt.Println("##########################################################")
			// Uncomment to save user info into file
			// filename := "user_data.json"
			// if f, err := os.Create(filename); err == nil {
			// 	fmt.Fprintf(os.Stderr, "returned events saved to %s\n", filename)
			// 	// encode the returned events in a file
			// 	enc := json.NewEncoder(f)
			// 	enc.SetIndent("", " ")
			// 	enc.Encode(evs)
			// 	f.Close()
			// } else {
			// 	panic(err)
			// }
		} else {
			cmd.Help()
		}
		return nil
	},
}

func init() {
	UserCmd.Flags().BoolVarP(&userData, "info", "", false, "Retrieve user info from the specified relay.")
}
