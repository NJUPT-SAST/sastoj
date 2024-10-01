package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.String("salt"),
		field.Enum("state").Values("NORMAL", "BANNED", "INACTIVE").Default("NORMAL"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission", Submission.Type),
		edge.To("login_sessions", LoginSession.Type),
		edge.To("owned_problems", Problem.Type),
		edge.From("groups", Group.Type).Ref("users"),
		edge.To("contest_results", ContestResult.Type),
	}
}
