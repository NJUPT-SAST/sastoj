package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("group_name").Default("unknown"),
		field.Int("root_group_id").Default(1),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admins", Contest.Type),
		edge.To("contestants", Contest.Type),
		edge.To("problems", Problem.Type),
		edge.To("users", User.Type),
		edge.To("subgroups", Group.Type).From("root_group").Field("root_group_id").Unique().Required(),
	}
}
