package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoginSession holds the schema definition for the LoginSession entity.
type LoginSession struct {
	ent.Schema
}

// Fields of the LoginSession.
func (LoginSession) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("user_id").Positive(),
		field.Time("create_time"),
	}
}

// Edges of the LoginSession.
func (LoginSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}
