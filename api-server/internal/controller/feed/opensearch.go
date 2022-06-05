package feed

import "time"

const index = "feed"

const scrollTTL = 5 * time.Minute

type post struct {
	From      int       `json:"from"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
