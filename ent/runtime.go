// Code generated by ent, DO NOT EDIT.

package ent

import (
	"sastoj/ent/contest"
	"sastoj/ent/group"
	"sastoj/ent/loginsession"
	"sastoj/ent/problem"
	"sastoj/ent/problemtype"
	"sastoj/ent/schema"
	"sastoj/ent/submission"
	"sastoj/ent/submissioncase"
	"sastoj/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contestFields := schema.Contest{}.Fields()
	_ = contestFields
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
	// groupDescIsRoot is the schema descriptor for is_root field.
	groupDescIsRoot := groupFields[2].Descriptor()
	// group.DefaultIsRoot holds the default value on creation for the is_root field.
	group.DefaultIsRoot = groupDescIsRoot.Default.(bool)
	loginsessionFields := schema.LoginSession{}.Fields()
	_ = loginsessionFields
	// loginsessionDescCreateTime is the schema descriptor for create_time field.
	loginsessionDescCreateTime := loginsessionFields[1].Descriptor()
	// loginsession.DefaultCreateTime holds the default value on creation for the create_time field.
	loginsession.DefaultCreateTime = loginsessionDescCreateTime.Default.(time.Time)
	problemFields := schema.Problem{}.Fields()
	_ = problemFields
	// problemDescScore is the schema descriptor for score field.
	problemDescScore := problemFields[4].Descriptor()
	// problem.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	problem.ScoreValidator = problemDescScore.Validators[0].(func(int16) error)
	// problemDescCaseVersion is the schema descriptor for case_version field.
	problemDescCaseVersion := problemFields[5].Descriptor()
	// problem.DefaultCaseVersion holds the default value on creation for the case_version field.
	problem.DefaultCaseVersion = problemDescCaseVersion.Default.(int16)
	// problemDescIndex is the schema descriptor for index field.
	problemDescIndex := problemFields[6].Descriptor()
	// problem.IndexValidator is a validator for the "index" field. It is called by the builders before save.
	problem.IndexValidator = problemDescIndex.Validators[0].(func(int16) error)
	// problemDescIsDeleted is the schema descriptor for is_deleted field.
	problemDescIsDeleted := problemFields[8].Descriptor()
	// problem.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	problem.DefaultIsDeleted = problemDescIsDeleted.Default.(bool)
	// problemDescMetadata is the schema descriptor for metadata field.
	problemDescMetadata := problemFields[12].Descriptor()
	// problem.DefaultMetadata holds the default value on creation for the metadata field.
	problem.DefaultMetadata = problemDescMetadata.Default.(map[string]string)
	problemtypeFields := schema.ProblemType{}.Fields()
	_ = problemtypeFields
	// problemtypeDescJudge is the schema descriptor for judge field.
	problemtypeDescJudge := problemtypeFields[5].Descriptor()
	// problemtype.DefaultJudge holds the default value on creation for the judge field.
	problemtype.DefaultJudge = problemtypeDescJudge.Default.(string)
	submissionFields := schema.Submission{}.Fields()
	_ = submissionFields
	// submissionDescCompileStdout is the schema descriptor for compile_stdout field.
	submissionDescCompileStdout := submissionFields[3].Descriptor()
	// submission.DefaultCompileStdout holds the default value on creation for the compile_stdout field.
	submission.DefaultCompileStdout = submissionDescCompileStdout.Default.(string)
	// submissionDescCompileStderr is the schema descriptor for compile_stderr field.
	submissionDescCompileStderr := submissionFields[4].Descriptor()
	// submission.DefaultCompileStderr holds the default value on creation for the compile_stderr field.
	submission.DefaultCompileStderr = submissionDescCompileStderr.Default.(string)
	// submissionDescCreateTime is the schema descriptor for create_time field.
	submissionDescCreateTime := submissionFields[6].Descriptor()
	// submission.DefaultCreateTime holds the default value on creation for the create_time field.
	submission.DefaultCreateTime = submissionDescCreateTime.Default.(time.Time)
	submissioncaseFields := schema.SubmissionCase{}.Fields()
	_ = submissioncaseFields
	// submissioncaseDescStdout is the schema descriptor for stdout field.
	submissioncaseDescStdout := submissioncaseFields[5].Descriptor()
	// submissioncase.DefaultStdout holds the default value on creation for the stdout field.
	submissioncase.DefaultStdout = submissioncaseDescStdout.Default.(string)
	// submissioncaseDescStderr is the schema descriptor for stderr field.
	submissioncaseDescStderr := submissioncaseFields[6].Descriptor()
	// submissioncase.DefaultStderr holds the default value on creation for the stderr field.
	submissioncase.DefaultStderr = submissioncaseDescStderr.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
}
