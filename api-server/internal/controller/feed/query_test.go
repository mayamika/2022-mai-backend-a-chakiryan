package feed

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestShit(t *testing.T) {
	s := "awd"
	q := buildSearchFeedQuery(nil, &s)
	b, err := json.MarshalIndent(q, "", "\t")
	fmt.Printf("err: %v\nwow: %s\n", err, string(b))
}
