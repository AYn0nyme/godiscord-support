package main

import (
	"fmt"
	"os"
	"strings"

	"godiscord.foo.ng/lib/pkg/classes"
	"godiscord.foo.ng/lib/pkg/enums"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	Client := classes.NewClient(
		enums.GatewayIntent.Guilds,
		enums.GatewayIntent.GuildMembers,
		enums.GatewayIntent.GuildMessages,
		enums.GatewayIntent.GuildModeration,
		enums.GatewayIntent.MessageContent,
	)

	Client.On("READY", func(args ...any) {
		fmt.Println("READY:", args[0].(classes.Client).User.Username)
	})

	Client.On("MESSAGE_CREATE", func(args ...any) {
		message := args[0].(classes.Message)
		//var message_args []string
		commandName := strings.Fields(message.Content)[0]
		if len(strings.Fields(message.Content)) > 1 {
			//message_args = strings.Fields(message.Content)[1:]
		}
		if commandName == "!ping" {
			message.Reply("Pong!")
		}
		if commandName == "!gh" || commandName == "!github" {
			message.Reply("https://github.com/AYn0nyme/godiscord")
		}
	})

	err := Client.Connect(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return
	}
}
