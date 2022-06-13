package feed

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

const index = "feed"

const scrollTTL = 5 * time.Minute

type post struct {
	From      int       `json:"from"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Images    []string  `json:"images"`
}

func (c *Controller) createIndex(ctx context.Context) error {
	d, _ := json.Marshal(&post{
		From:      0,
		Text:      "Text",
		CreatedAt: time.Now(),
		Images:    []string{uuid.NewString(), uuid.NewString()},
	})

	createRequest := opensearchapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(d),
		Refresh: "wait_for",
	}
	_, err := createRequest.Do(ctx, c.client)
	return err
}
