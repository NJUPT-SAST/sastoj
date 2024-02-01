// Code generated by ent, DO NOT EDIT.

package group

import (
	"sastoj/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldID, id))
}

// GroupName applies equality check predicate on the "group_name" field. It's identical to GroupNameEQ.
func GroupName(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldGroupName, v))
}

// GroupNameEQ applies the EQ predicate on the "group_name" field.
func GroupNameEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldEQ(FieldGroupName, v))
}

// GroupNameNEQ applies the NEQ predicate on the "group_name" field.
func GroupNameNEQ(v string) predicate.Group {
	return predicate.Group(sql.FieldNEQ(FieldGroupName, v))
}

// GroupNameIn applies the In predicate on the "group_name" field.
func GroupNameIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldIn(FieldGroupName, vs...))
}

// GroupNameNotIn applies the NotIn predicate on the "group_name" field.
func GroupNameNotIn(vs ...string) predicate.Group {
	return predicate.Group(sql.FieldNotIn(FieldGroupName, vs...))
}

// GroupNameGT applies the GT predicate on the "group_name" field.
func GroupNameGT(v string) predicate.Group {
	return predicate.Group(sql.FieldGT(FieldGroupName, v))
}

// GroupNameGTE applies the GTE predicate on the "group_name" field.
func GroupNameGTE(v string) predicate.Group {
	return predicate.Group(sql.FieldGTE(FieldGroupName, v))
}

// GroupNameLT applies the LT predicate on the "group_name" field.
func GroupNameLT(v string) predicate.Group {
	return predicate.Group(sql.FieldLT(FieldGroupName, v))
}

// GroupNameLTE applies the LTE predicate on the "group_name" field.
func GroupNameLTE(v string) predicate.Group {
	return predicate.Group(sql.FieldLTE(FieldGroupName, v))
}

// GroupNameContains applies the Contains predicate on the "group_name" field.
func GroupNameContains(v string) predicate.Group {
	return predicate.Group(sql.FieldContains(FieldGroupName, v))
}

// GroupNameHasPrefix applies the HasPrefix predicate on the "group_name" field.
func GroupNameHasPrefix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasPrefix(FieldGroupName, v))
}

// GroupNameHasSuffix applies the HasSuffix predicate on the "group_name" field.
func GroupNameHasSuffix(v string) predicate.Group {
	return predicate.Group(sql.FieldHasSuffix(FieldGroupName, v))
}

// GroupNameEqualFold applies the EqualFold predicate on the "group_name" field.
func GroupNameEqualFold(v string) predicate.Group {
	return predicate.Group(sql.FieldEqualFold(FieldGroupName, v))
}

// GroupNameContainsFold applies the ContainsFold predicate on the "group_name" field.
func GroupNameContainsFold(v string) predicate.Group {
	return predicate.Group(sql.FieldContainsFold(FieldGroupName, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.Group(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.Group(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Group) predicate.Group {
	return predicate.Group(sql.NotPredicates(p))
}
