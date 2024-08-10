// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/loginsession"
	"sastoj/ent/problem"
	"sastoj/ent/problemcase"
	"sastoj/ent/schema"
	"sastoj/ent/submission"
	"sastoj/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contestFields := schema.Contest{}.Fields()
	_ = contestFields
	// contestDescStatus is the schema descriptor for status field.
	contestDescStatus := contestFields[3].Descriptor()
	// contest.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	contest.StatusValidator = contestDescStatus.Validators[0].(func(int16) error)
	// contestDescType is the schema descriptor for type field.
	contestDescType := contestFields[4].Descriptor()
	// contest.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	contest.TypeValidator = contestDescType.Validators[0].(func(int16) error)
	// contestDescExtraTime is the schema descriptor for extra_time field.
	contestDescExtraTime := contestFields[8].Descriptor()
	// contest.DefaultExtraTime holds the default value on creation for the extra_time field.
	contest.DefaultExtraTime = contestDescExtraTime.Default.(int16)
	// contest.ExtraTimeValidator is a validator for the "extra_time" field. It is called by the builders before save.
	contest.ExtraTimeValidator = contestDescExtraTime.Validators[0].(func(int16) error)
	// contestDescCreateTime is the schema descriptor for create_time field.
	contestDescCreateTime := contestFields[9].Descriptor()
	// contest.DefaultCreateTime holds the default value on creation for the create_time field.
	contest.DefaultCreateTime = contestDescCreateTime.Default.(time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescGroupName is the schema descriptor for group_name field.
	groupDescGroupName := groupFields[1].Descriptor()
	// group.DefaultGroupName holds the default value on creation for the group_name field.
	group.DefaultGroupName = groupDescGroupName.Default.(string)
	// groupDescRootGroupID is the schema descriptor for root_group_id field.
	groupDescRootGroupID := groupFields[2].Descriptor()
	// group.DefaultRootGroupID holds the default value on creation for the root_group_id field.
	group.DefaultRootGroupID = groupDescRootGroupID.Default.(int64)
	loginsessionFields := schema.LoginSession{}.Fields()
	_ = loginsessionFields
	// loginsessionDescCreateTime is the schema descriptor for create_time field.
	loginsessionDescCreateTime := loginsessionFields[1].Descriptor()
	// loginsession.DefaultCreateTime holds the default value on creation for the create_time field.
	loginsession.DefaultCreateTime = loginsessionDescCreateTime.Default.(time.Time)
	problemFields := schema.Problem{}.Fields()
	_ = problemFields
	// problemDescPoint is the schema descriptor for point field.
	problemDescPoint := problemFields[3].Descriptor()
	// problem.PointValidator is a validator for the "point" field. It is called by the builders before save.
	problem.PointValidator = problemDescPoint.Validators[0].(func(int16) error)
	// problemDescCaseVersion is the schema descriptor for case_version field.
	problemDescCaseVersion := problemFields[4].Descriptor()
	// problem.DefaultCaseVersion holds the default value on creation for the case_version field.
	problem.DefaultCaseVersion = problemDescCaseVersion.Default.(int16)
	// problemDescIndex is the schema descriptor for index field.
	problemDescIndex := problemFields[5].Descriptor()
	// problem.IndexValidator is a validator for the "index" field. It is called by the builders before save.
	problem.IndexValidator = problemDescIndex.Validators[0].(func(int16) error)
	// problemDescRestrictPresentation is the schema descriptor for restrict_presentation field.
	problemDescRestrictPresentation := problemFields[6].Descriptor()
	// problem.DefaultRestrictPresentation holds the default value on creation for the restrict_presentation field.
	problem.DefaultRestrictPresentation = problemDescRestrictPresentation.Default.(bool)
	// problemDescIsDeleted is the schema descriptor for is_deleted field.
	problemDescIsDeleted := problemFields[7].Descriptor()
	// problem.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	problem.DefaultIsDeleted = problemDescIsDeleted.Default.(bool)
	// problemDescVisibility is the schema descriptor for visibility field.
	problemDescVisibility := problemFields[11].Descriptor()
	// problem.DefaultVisibility holds the default value on creation for the visibility field.
	problem.DefaultVisibility = problemDescVisibility.Default.(int8)
	problemcaseHooks := schema.ProblemCase{}.Hooks()
	problemcase.Hooks[0] = problemcaseHooks[0]
	problemcaseFields := schema.ProblemCase{}.Fields()
	_ = problemcaseFields
	// problemcaseDescPoint is the schema descriptor for point field.
	problemcaseDescPoint := problemcaseFields[1].Descriptor()
	// problemcase.PointValidator is a validator for the "point" field. It is called by the builders before save.
	problemcase.PointValidator = problemcaseDescPoint.Validators[0].(func(int16) error)
	// problemcaseDescIndex is the schema descriptor for index field.
	problemcaseDescIndex := problemcaseFields[2].Descriptor()
	// problemcase.IndexValidator is a validator for the "index" field. It is called by the builders before save.
	problemcase.IndexValidator = problemcaseDescIndex.Validators[0].(func(int16) error)
	// problemcaseDescIsAuto is the schema descriptor for is_auto field.
	problemcaseDescIsAuto := problemcaseFields[3].Descriptor()
	// problemcase.DefaultIsAuto holds the default value on creation for the is_auto field.
	problemcase.DefaultIsAuto = problemcaseDescIsAuto.Default.(bool)
	// problemcaseDescIsDeleted is the schema descriptor for is_deleted field.
	problemcaseDescIsDeleted := problemcaseFields[4].Descriptor()
	// problemcase.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	problemcase.DefaultIsDeleted = problemcaseDescIsDeleted.Default.(bool)
	submissionFields := schema.Submission{}.Fields()
	_ = submissionFields
	// submissionDescCreateTime is the schema descriptor for create_time field.
	submissionDescCreateTime := submissionFields[5].Descriptor()
	// submission.DefaultCreateTime holds the default value on creation for the create_time field.
	submission.DefaultCreateTime = submissionDescCreateTime.Default.(time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[4].Descriptor()
	// user.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	user.StatusValidator = userDescStatus.Validators[0].(func(int16) error)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)
