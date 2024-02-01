// Code generated by ent, DO NOT EDIT.

package ent

import (
	"sastoj/ent/group"
	"sastoj/ent/schema"
	"sastoj/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescGroupName is the schema descriptor for group_name field.
	groupDescGroupName := groupFields[1].Descriptor()
	// group.DefaultGroupName holds the default value on creation for the group_name field.
	group.DefaultGroupName = groupDescGroupName.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// userDescState is the schema descriptor for state field.
	userDescState := userFields[4].Descriptor()
	// user.StateValidator is a validator for the "state" field. It is called by the builders before save.
	user.StateValidator = userDescState.Validators[0].(func(int) error)
	// userDescGroupID is the schema descriptor for group_id field.
	userDescGroupID := userFields[5].Descriptor()
	// user.GroupIDValidator is a validator for the "group_id" field. It is called by the builders before save.
	user.GroupIDValidator = userDescGroupID.Validators[0].(func(int) error)
}
