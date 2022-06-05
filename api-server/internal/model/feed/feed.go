package feed

type Post struct {
	ID   string
	From int
	Text string
}

type PostEdge struct {
	Node   *Post
	Cursor string
}

type PageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     *string
	EndCursor       *string
}

type PostConnection struct {
	TotalCount int
	PageInfo   PageInfo
	Edges      []*PostEdge
}
