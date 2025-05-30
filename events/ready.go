package events

import (
	"fmt"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
)

func Ready(args ...any) {
	Client := args[0].(classes.Client)

	Client.SetPresence(classes.PresenceUpdate{
		Activities: []classes.Activity{
			classes.Activity{
				Name: "godiscord",
				Type: enums.ActivityType.Streaming,
				URL:  "https://twitch.tv/godiscord",
			},
		},
		Status: "idle",
		AFK:    false,
	})

	fmt.Println(Client.User.Username, "is ready!")
}
