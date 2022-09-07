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
	s.ChannelMessageSend(m.ChannelID, "I couldn't understand your message,please send in the same model below:\n  >>> ** !news fall 2022 0  {} **\n** !news [seasons] [year] [page] **")
	return
}

func ExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	message := strings.Split(strings.TrimSpace(m.Content), " ")
	if message[0] != "!news" {
		return
	}
	if len(message) < 4 {
		defaultMessage(s, m)
		return
	}
	controller := anime.Create(s, m)
	page := 0
	if message[3] != "" {
		page, _ = strconv.Atoi(string(message[1]))
	}
	controller.Get(message[1], message[2], page)
}
