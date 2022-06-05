package feed

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/auth"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/feed"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

type Controller struct {
	client *opensearch.Client
}

func NewController(
	client *opensearch.Client,
) *Controller {
	return &Controller{
		client: client,
	}
}

func (c *Controller) Feed(
	ctx context.Context,
	first int,
	after *string,
	search *string,
) (*feed.FeedPayload, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}

	searchRequest, err := buildFeedSearchRequest(ctx, t, first, after, search)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	res, err := searchRequest.Do(ctx, c.client)
	if err != nil {
		return nil, fmt.Errorf("search request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("search request failed")
	}

	type resBody struct {
		ScrollID string `json:"_scroll_id"`
		Hits     struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				ID     string `json:"_id"`
				Source post   `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	var b resBody
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	payload := &feed.FeedPayload{
		TotalCount: b.Hits.Total.Value,
	}
	for _, h := range b.Hits.Hits {
		payload.Posts = append(payload.Posts, &feed.Post{
			ID:        h.ID,
			From:      h.Source.From,
			Text:      h.Source.Text,
			CreatedAt: h.Source.CreatedAt,
		})
	}
	if len(b.Hits.Hits) > 0 {
		payload.HasNextPage = true
		payload.Scroll = &b.ScrollID
	}

	return payload, nil
}

func (c *Controller) PublishPost(
	ctx context.Context,
	input feed.PostInput,
) (*feed.Post, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return nil, auth.ErrNotAuthenticated
	}

	p := post{
		From:      t.UserID,
		Text:      input.Text,
		CreatedAt: time.Now(),
	}

	var reqBody bytes.Buffer
	if err := json.NewEncoder(&reqBody).Encode(p); err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}
	indexRequest := opensearchapi.IndexRequest{
		Index: index,
		Body:  &reqBody,
	}

	res, err := indexRequest.Do(ctx, c.client)
	if err != nil {
		return nil, fmt.Errorf("index request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, errors.New("post is not created")
	}

	type resBody struct {
		ID string `json:"_id"`
	}
	var b resBody
	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &feed.Post{
		ID:        b.ID,
		From:      p.From,
		Text:      p.Text,
		CreatedAt: p.CreatedAt,
	}, nil
}

func buildFeedSearchRequest(
	ctx context.Context,
	t token.Token,
	first int,
	after *string,
	search *string,
) (opensearchapi.Request, error) {
	if after != nil {
		return opensearchapi.ScrollRequest{
			ScrollID: *after,
			Scroll:   scrollTTL,
		}, nil
	}

	client := ent.FromContext(ctx)
	u, err := client.User.Get(ctx, t.UserID)
	if err != nil {
		return nil, fmt.Errorf("load user: %w", err)
	}
	friends, err := u.Friends(ctx)
	if err != nil {
		return nil, fmt.Errorf("load friends: %w", err)
	}

	from := []int{u.ID}
	for _, f := range friends {
		from = append(from, f.ID)
	}
	query := buildSearchFeedQuery(from, search)

	var reqBody bytes.Buffer
	if err := json.NewEncoder(&reqBody).Encode(query); err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	return opensearchapi.SearchRequest{
		Index:  []string{index},
		Scroll: scrollTTL,
		Size:   &first,
		Sort:   []string{"created_at:desc"},
		Body:   &reqBody,
	}, nil
}
