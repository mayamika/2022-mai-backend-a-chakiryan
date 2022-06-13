package feed

import "time"

type PostInput struct {
	Text   string
	Images []string
}

type Post struct {
	ID        string
	From      int
	Text      string
	CreatedAt time.Time
	Images    []string
}

type FeedPayload struct {
	TotalCount  int
	HasNextPage bool
	Scroll      *string
	Posts       []*Post
}
