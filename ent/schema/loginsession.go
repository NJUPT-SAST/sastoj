package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LoginSession holds the schema definition for the LoginSession entity.
type LoginSession struct {
	ent.Schema
}

func (LoginSession) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "login_session"},
	}
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
		edge.From("users", User.Type).Ref("login_session"),
	}
}
