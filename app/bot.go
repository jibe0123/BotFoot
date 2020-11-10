package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jibe0123/discordBot/commands/foot"
	"github.com/jibe0123/discordBot/commands/ping"
	"log"
	"strings"
)

var BotID string

func Start(token string) *discordgo.Session {
	goBot, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println("Bot is running!")
	return goBot
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		pingCommand := ping.MakePingCommand(s, m)
		pingCommand.Execute()
	}

	if strings.HasPrefix(m.Content, "!foot") {
		log.Print(strings.Fields(m.Content))
		footCommand := foot.GetLastCompetitionsCommand(s, m)
		footCommand.Execute()

	}

}
