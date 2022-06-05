package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// FriendRequest holds the schema definition for the FriendRequest entity.
type FriendRequest struct {
	ent.Schema
}

// Fields of the FriendRequest.
func (FriendRequest) Fields() []ent.Field {
	return nil
}

// Edges of the FriendRequest.
func (FriendRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("from", User.Type).
			Required().
			Unique(),
		edge.To("to", User.Type).
			Required().
			Unique(),
	}
}
