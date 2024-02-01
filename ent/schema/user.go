package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("username").Default("unknown"),
		field.String("password"),
		field.String("salt"),
		field.Int("state").Positive(),
		field.Int("group_id").Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
