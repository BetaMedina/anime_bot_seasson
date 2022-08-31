package anime

type AnimeAttributesTitleEntity struct {
	En   string `json:"en"`
	EnJp string `json:"en_jp"`
	JaJp string `json:"ja_jp"`
}

type AnimeAttributesEntity struct {
	Slug           string                     `json:"slug"`
	Type           string                     `json:"type"`
	Titles         AnimeAttributesTitleEntity `json:"titles"`
	CanonicalTitle string                     `json:"canonicalTitle"`
	AverageRating  string                     `json:"averageRating"`
	PopularityRank int32                      `json:"popularityRank"`
	Status         string                     `json:"status"`
	PosterImage    AnimePosterImageEntity     `json:"posterImage"`
	EpisodeCount   int32                      `json:"episodeCount"`
	EpisodeLength  int32                      `json:"episodeLength"`
	Nsfw           bool                       `json:"nsfw"`
	AgeRatingGuide string                     `json:"ageRatingGuide"`
	Synopsis       string                     `json:"synopsis"`
}

type AnimePosterImageEntity struct {
	Original string `json:"original"`
	Large    string `json:"large"`
	Medium   string `json:"medium"`
}

type AnimeLinksEntity struct {
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type AnimeDataEntity struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes AnimeAttributesEntity `json:"attributes"`
}

type AnimeEntity struct {
	Data  []AnimeDataEntity `json:"data"`
	Links AnimeLinksEntity  `json:"links"`
}
