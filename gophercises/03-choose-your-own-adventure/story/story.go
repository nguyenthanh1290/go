package story

import (
	"encoding/json"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

const BeginArc = "intro"

func New(data []byte) (Story, error) {
	var story Story
	err := json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
