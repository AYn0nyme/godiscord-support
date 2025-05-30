package main

import (
	"os"

	"github.com/AYn0nyme/godiscord-support/events"
	_ "github.com/joho/godotenv/autoload"
	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"
	"godiscord.foo.ng/lib/pkg/slash"
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

	Client.On("READY", events.Ready)

	Client.On("INTERACTION_CREATE", events.InteractionCreate)

	err := Client.Connect()
	if err != nil {
		return
	}
}
