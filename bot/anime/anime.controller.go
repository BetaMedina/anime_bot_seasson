package anime

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Get struct {
	service GetAnimeRows
	s       *discordgo.Session
	m       *discordgo.MessageCreate
}

type IController interface {
	Get(season, year string, page int) (AnimeEntity, error)
}

func (s Get) Get(season, year string, page int) (AnimeEntity, error) {
	animeRow, err := s.service.GetAnimeRows(page, season, year)
	if err != nil {
		return AnimeEntity{}, err
	}
	if len(animeRow.Data) <= 0 {
		message := fmt.Sprintf(">>> We didn't find any more anime for the reported conditions ")
		s.s.ChannelMessageSend(s.m.ChannelID, message)

		return AnimeEntity{}, nil
	}
	for _, r := range animeRow.Data {
		message := fmt.Sprintf(">>> \nAnime: ** %s**\nType: ** %s **\nSynopese: **%s**\n\n\n", r.Attributes.CanonicalTitle, r.Type, r.Attributes.Synopsis)
		s.s.ChannelMessageSend(s.m.ChannelID, message)
	}
	return animeRow, nil
}
