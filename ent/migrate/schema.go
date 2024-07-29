// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContestsColumns holds the columns for the "contests" table.
	ContestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt16},
		{Name: "type", Type: field.TypeInt16},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "language", Type: field.TypeString},
		{Name: "extra_time", Type: field.TypeInt16, Default: 0},
		{Name: "create_time", Type: field.TypeTime},
	}
	// ContestsTable holds the schema information for the "contests" table.
	ContestsTable = &schema.Table{
		Name:       "contests",
		Columns:    ContestsColumns,
		PrimaryKey: []*schema.Column{ContestsColumns[0]},
	}
	// ContestResultsColumns holds the columns for the "contest_results" table.
	ContestResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "score", Type: field.TypeInt32},
		{Name: "rank", Type: field.TypeInt32},
		{Name: "score_time", Type: field.TypeInt32},
		{Name: "penalty", Type: field.TypeInt32},
		{Name: "contest_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// ContestResultsTable holds the schema information for the "contest_results" table.
	ContestResultsTable = &schema.Table{
		Name:       "contest_results",
		Columns:    ContestResultsColumns,
		PrimaryKey: []*schema.Column{ContestResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contest_results_contests_contest_results",
				Columns:    []*schema.Column{ContestResultsColumns[5]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "contest_results_users_contest_results",
				Columns:    []*schema.Column{ContestResultsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "group_name", Type: field.TypeString, Default: "unknown"},
		{Name: "root_group_id", Type: field.TypeInt64, Nullable: true, Default: 1},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_groups_subgroups",
				Columns:    []*schema.Column{GroupsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LoginSessionColumns holds the columns for the "login_session" table.
	LoginSessionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// LoginSessionTable holds the schema information for the "login_session" table.
	LoginSessionTable = &schema.Table{
		Name:       "login_session",
		Columns:    LoginSessionColumns,
		PrimaryKey: []*schema.Column{LoginSessionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "login_session_users_login_sessions",
				Columns:    []*schema.Column{LoginSessionColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProblemsColumns holds the columns for the "problems" table.
	ProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "point", Type: field.TypeInt16},
		{Name: "case_version", Type: field.TypeInt16, Default: 1},
		{Name: "index", Type: field.TypeInt16},
		{Name: "restrict_presentation", Type: field.TypeBool, Default: true},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "config", Type: field.TypeString},
		{Name: "visibility", Type: field.TypeInt8, Default: 0},
		{Name: "contest_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// ProblemsTable holds the schema information for the "problems" table.
	ProblemsTable = &schema.Table{
		Name:       "problems",
		Columns:    ProblemsColumns,
		PrimaryKey: []*schema.Column{ProblemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "problems_contests_problems",
				Columns:    []*schema.Column{ProblemsColumns[10]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "problems_users_owned_problems",
				Columns:    []*schema.Column{ProblemsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProblemCasesColumns holds the columns for the "problem_cases" table.
	ProblemCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "point", Type: field.TypeInt16},
		{Name: "index", Type: field.TypeInt16},
		{Name: "is_auto", Type: field.TypeBool, Default: false},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "problem_id", Type: field.TypeInt64},
	}
	// ProblemCasesTable holds the schema information for the "problem_cases" table.
	ProblemCasesTable = &schema.Table{
		Name:       "problem_cases",
		Columns:    ProblemCasesColumns,
		PrimaryKey: []*schema.Column{ProblemCasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "problem_cases_problems_problem_cases",
				Columns:    []*schema.Column{ProblemCasesColumns[5]},
				RefColumns: []*schema.Column{ProblemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubmissionsColumns holds the columns for the "submissions" table.
	SubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "code", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeInt16},
		{Name: "compile_message", Type: field.TypeString, Nullable: true},
		{Name: "point", Type: field.TypeInt16},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "total_time", Type: field.TypeInt32},
		{Name: "max_memory", Type: field.TypeInt32},
		{Name: "language", Type: field.TypeString},
		{Name: "case_version", Type: field.TypeInt8},
		{Name: "problem_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// SubmissionsTable holds the schema information for the "submissions" table.
	SubmissionsTable = &schema.Table{
		Name:       "submissions",
		Columns:    SubmissionsColumns,
		PrimaryKey: []*schema.Column{SubmissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submissions_problems_submission",
				Columns:    []*schema.Column{SubmissionsColumns[10]},
				RefColumns: []*schema.Column{ProblemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "submissions_users_submission",
				Columns:    []*schema.Column{SubmissionsColumns[11]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SubmissionCasesColumns holds the columns for the "submission_cases" table.
	SubmissionCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "state", Type: field.TypeInt16},
		{Name: "point", Type: field.TypeInt16},
		{Name: "message", Type: field.TypeString, Size: 2147483647},
		{Name: "time", Type: field.TypeInt32},
		{Name: "memory", Type: field.TypeInt32},
		{Name: "problem_case_id", Type: field.TypeInt64},
		{Name: "submission_id", Type: field.TypeInt64},
	}
	// SubmissionCasesTable holds the schema information for the "submission_cases" table.
	SubmissionCasesTable = &schema.Table{
		Name:       "submission_cases",
		Columns:    SubmissionCasesColumns,
		PrimaryKey: []*schema.Column{SubmissionCasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submission_cases_problem_cases_submission_cases",
				Columns:    []*schema.Column{SubmissionCasesColumns[6]},
				RefColumns: []*schema.Column{ProblemCasesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "submission_cases_submissions_submission_cases",
				Columns:    []*schema.Column{SubmissionCasesColumns[7]},
				RefColumns: []*schema.Column{SubmissionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt16},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_groups_users",
				Columns:    []*schema.Column{UsersColumns[5]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ContestContestantsColumns holds the columns for the "contest_contestants" table.
	ContestContestantsColumns = []*schema.Column{
		{Name: "contest_id", Type: field.TypeInt64},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// ContestContestantsTable holds the schema information for the "contest_contestants" table.
	ContestContestantsTable = &schema.Table{
		Name:       "contest_contestants",
		Columns:    ContestContestantsColumns,
		PrimaryKey: []*schema.Column{ContestContestantsColumns[0], ContestContestantsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contest_contestants_contest_id",
				Columns:    []*schema.Column{ContestContestantsColumns[0]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "contest_contestants_group_id",
				Columns:    []*schema.Column{ContestContestantsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ContestManagersColumns holds the columns for the "contest_managers" table.
	ContestManagersColumns = []*schema.Column{
		{Name: "contest_id", Type: field.TypeInt64},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// ContestManagersTable holds the schema information for the "contest_managers" table.
	ContestManagersTable = &schema.Table{
		Name:       "contest_managers",
		Columns:    ContestManagersColumns,
		PrimaryKey: []*schema.Column{ContestManagersColumns[0], ContestManagersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contest_managers_contest_id",
				Columns:    []*schema.Column{ContestManagersColumns[0]},
				RefColumns: []*schema.Column{ContestsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "contest_managers_group_id",
				Columns:    []*schema.Column{ContestManagersColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProblemJudgersColumns holds the columns for the "problem_judgers" table.
	ProblemJudgersColumns = []*schema.Column{
		{Name: "problem_id", Type: field.TypeInt64},
		{Name: "group_id", Type: field.TypeInt64},
	}
	// ProblemJudgersTable holds the schema information for the "problem_judgers" table.
	ProblemJudgersTable = &schema.Table{
		Name:       "problem_judgers",
		Columns:    ProblemJudgersColumns,
		PrimaryKey: []*schema.Column{ProblemJudgersColumns[0], ProblemJudgersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "problem_judgers_problem_id",
				Columns:    []*schema.Column{ProblemJudgersColumns[0]},
				RefColumns: []*schema.Column{ProblemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "problem_judgers_group_id",
				Columns:    []*schema.Column{ProblemJudgersColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubmissionContestResultsColumns holds the columns for the "submission_contest_results" table.
	SubmissionContestResultsColumns = []*schema.Column{
		{Name: "submission_id", Type: field.TypeInt64},
		{Name: "contest_result_id", Type: field.TypeInt},
	}
	// SubmissionContestResultsTable holds the schema information for the "submission_contest_results" table.
	SubmissionContestResultsTable = &schema.Table{
		Name:       "submission_contest_results",
		Columns:    SubmissionContestResultsColumns,
		PrimaryKey: []*schema.Column{SubmissionContestResultsColumns[0], SubmissionContestResultsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "submission_contest_results_submission_id",
				Columns:    []*schema.Column{SubmissionContestResultsColumns[0]},
				RefColumns: []*schema.Column{SubmissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "submission_contest_results_contest_result_id",
				Columns:    []*schema.Column{SubmissionContestResultsColumns[1]},
				RefColumns: []*schema.Column{ContestResultsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContestsTable,
		ContestResultsTable,
		GroupsTable,
		LoginSessionTable,
		ProblemsTable,
		ProblemCasesTable,
		SubmissionsTable,
		SubmissionCasesTable,
		UsersTable,
		ContestContestantsTable,
		ContestManagersTable,
		ProblemJudgersTable,
		SubmissionContestResultsTable,
	}
)

func init() {
	ContestResultsTable.ForeignKeys[0].RefTable = ContestsTable
	ContestResultsTable.ForeignKeys[1].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = GroupsTable
	LoginSessionTable.ForeignKeys[0].RefTable = UsersTable
	LoginSessionTable.Annotation = &entsql.Annotation{
		Table: "login_session",
	}
	ProblemsTable.ForeignKeys[0].RefTable = ContestsTable
	ProblemsTable.ForeignKeys[1].RefTable = UsersTable
	ProblemCasesTable.ForeignKeys[0].RefTable = ProblemsTable
	SubmissionsTable.ForeignKeys[0].RefTable = ProblemsTable
	SubmissionsTable.ForeignKeys[1].RefTable = UsersTable
	SubmissionCasesTable.ForeignKeys[0].RefTable = ProblemCasesTable
	SubmissionCasesTable.ForeignKeys[1].RefTable = SubmissionsTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	ContestContestantsTable.ForeignKeys[0].RefTable = ContestsTable
	ContestContestantsTable.ForeignKeys[1].RefTable = GroupsTable
	ContestManagersTable.ForeignKeys[0].RefTable = ContestsTable
	ContestManagersTable.ForeignKeys[1].RefTable = GroupsTable
	ProblemJudgersTable.ForeignKeys[0].RefTable = ProblemsTable
	ProblemJudgersTable.ForeignKeys[1].RefTable = GroupsTable
	SubmissionContestResultsTable.ForeignKeys[0].RefTable = SubmissionsTable
	SubmissionContestResultsTable.ForeignKeys[1].RefTable = ContestResultsTable
}
