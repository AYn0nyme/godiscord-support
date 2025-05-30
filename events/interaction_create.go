package events

import (
	"fmt"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/types"
)

func InteractionCreate(args ...any) {
	interaction := args[0].(classes.BaseInteraction)
	if interaction.Data.Name == "ping" {
		interaction.Reply(classes.MessageData{
			Embeds: []classes.Embed{
				classes.NewEmbed().SetDescription(fmt.Sprintf("🏓 **%d**ms", Client.GetWSPing())).SetColor("00ADD8"),
			},
		})
	} else if interaction.Data.Name == "links" {
		interaction.Reply(classes.MessageData{
			Embeds: []classes.Embed{
				classes.NewEmbed().SetDescription("[Repo of godiscord](https://github.com/AYn0nyme/godiscord)\n[Repo of this bot](https://github.com/AYn0nyme/godiscord-support)").SetColor("00ADD8"),
			},
			Flags: []types.MessageFlag{
				enums.MessageFlags.Ephemeral,
			},
		})
	}
}
