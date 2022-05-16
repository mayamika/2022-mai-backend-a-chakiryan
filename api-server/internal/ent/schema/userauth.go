package schema

import "entgo.io/ent"

// UserAuth holds the schema definition for the UserAuth entity.
type UserAuth struct {
	ent.Schema
}

// Fields of the UserAuth.
func (UserAuth) Fields() []ent.Field {
	return nil
}

// Edges of the UserAuth.
func (UserAuth) Edges() []ent.Edge {
	return nil
}
