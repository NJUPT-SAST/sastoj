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
		field.String("username").Default("unknown"),
		field.String("password"),
		field.String("salt"),
		field.Int16("status").NonNegative(),
		field.Int64("group_id"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submission", Submission.Type),
		edge.To("login_sessions", LoginSession.Type),
		edge.To("owned_problems", Problem.Type),
		edge.From("groups", Group.Type).Ref("users").Field("group_id").Unique().Required(),
		edge.To("contest_results", ContestResult.Type),
	}
}
