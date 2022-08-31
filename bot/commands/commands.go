package commands

import (
	"anime-seasson-bot/bot/anime"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func defaultMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "I couldn't understand your message,please send in the same model below:\n  >>> ** !news 0 2022 fall {} **\n** !news [seasons] [year] [page] **")
	return
}

func ExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := strings.Split(strings.TrimSpace(m.Content), " ")
	switch message[0] {
	case "!news":
		controller := anime.Create(s, m)
		page := 0
		if message[3] != "" {
			page, _ = strconv.Atoi(string(message[1]))
		}
		controller.Get(message[1], message[2], page)
		break
	default:
		defaultMessage(s, m)
		break
	}
}
