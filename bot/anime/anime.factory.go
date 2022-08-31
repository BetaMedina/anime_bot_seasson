package anime

import (
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Create(s *discordgo.Session, m *discordgo.MessageCreate) IController {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	service := GetAnimeRows{client: &client}
	return &Get{service: service, s: s, m: m}
}
