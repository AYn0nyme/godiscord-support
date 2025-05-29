package main

import (
	"fmt"
	"os"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/slash"
	"godiscord.foo.ng/lib/pkg/types"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	Client := classes.NewClient(
		os.Getenv("DISCORD_TOKEN"),
		enums.GatewayIntent.Guilds,
		enums.GatewayIntent.GuildMembers,
		enums.GatewayIntent.GuildMessages,
		enums.GatewayIntent.GuildModeration,
		enums.GatewayIntent.MessageContent,
	)

	slash.RegisterGuildCommands("1375914465064915144", []classes.SlashCommandData{
		{
			Name:        "ping",
			Description: "Pong! Get the ping of the bot",
			Type:        enums.InteractionType.ChatInput,
		},
		{
			Name:        "links",
			Description: "Get the links of the project",
			Type:        enums.InteractionType.ChatInput,
		},
	}, "1373794354677813290")

	Client.On("READY", func(args ...any) {
		fmt.Println("READY:", args[0].(classes.Client).User.Username)
		Client.SetPresence(classes.PresenceUpdate{
			Activities: []classes.Activity{
				classes.Activity{
					Name: "godiscord",
					Type: enums.ActivityType.Streaming,
					URL:  "https://twitch.tv/godiscord",
				},
			},
			Status: "dnd",
			AFK:    false,
		})
	})

	Client.On("INTERACTION_CREATE", func(args ...any) {
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
	})

	err := Client.Connect()
	if err != nil {
		return
	}
}
