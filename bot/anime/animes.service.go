package anime

import (
	"anime-seasson-bot/internal/envs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetAnimeRows struct {
	client *http.Client
}

func (c *GetAnimeRows) GetAnimeRows(page int, season string, year string) (AnimeEntity, error) {
	var animeRow AnimeEntity
	url := envs.LoadApiEnvs()
	resp, err := c.client.Get(fmt.Sprintf("%s?filter[seasonYear]=%s&filter[season]=%s&page[limit]=10&page[offset]=%v", url.API_URL, year, season, page*10))
	if err != nil {
		return animeRow, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &animeRow)
	return animeRow, nil
}
