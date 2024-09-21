// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/case/service/v1/case.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Case with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Case) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Case with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CaseMultiError, or nil if none found.
func (m *Case) ValidateAll() error {
	return m.validate(true)
}

func (m *Case) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Point

	// no validation rules for Index

	// no validation rules for IsAuto

	if len(errors) > 0 {
		return CaseMultiError(errors)
	}

	return nil
}

// CaseMultiError is an error wrapping multiple validation errors returned by
// Case.ValidateAll() if the designated constraints aren't met.
type CaseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CaseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CaseMultiError) AllErrors() []error { return m }

// CaseValidationError is the validation error returned by Case.Validate if the
// designated constraints aren't met.
type CaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CaseValidationError) ErrorName() string { return "CaseValidationError" }

// Error satisfies the builtin error interface
func (e CaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CaseValidationError{}

// Validate checks the field values on DeleteCasesByProblemIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCasesByProblemIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCasesByProblemIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeleteCasesByProblemIdRequestMultiError, or nil if none found.
func (m *DeleteCasesByProblemIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCasesByProblemIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProblemId

	if len(errors) > 0 {
		return DeleteCasesByProblemIdRequestMultiError(errors)
	}

	return nil
}

// DeleteCasesByProblemIdRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteCasesByProblemIdRequest.ValidateAll()
// if the designated constraints aren't met.
type DeleteCasesByProblemIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCasesByProblemIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCasesByProblemIdRequestMultiError) AllErrors() []error { return m }

// DeleteCasesByProblemIdRequestValidationError is the validation error
// returned by DeleteCasesByProblemIdRequest.Validate if the designated
// constraints aren't met.
type DeleteCasesByProblemIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCasesByProblemIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCasesByProblemIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCasesByProblemIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCasesByProblemIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCasesByProblemIdRequestValidationError) ErrorName() string {
	return "DeleteCasesByProblemIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCasesByProblemIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCasesByProblemIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCasesByProblemIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCasesByProblemIdRequestValidationError{}

// Validate checks the field values on DeleteCasesByProblemIdReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCasesByProblemIdReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCasesByProblemIdReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCasesByProblemIdReplyMultiError, or nil if none found.
func (m *DeleteCasesByProblemIdReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCasesByProblemIdReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCasesByProblemIdReplyMultiError(errors)
	}

	return nil
}

// DeleteCasesByProblemIdReplyMultiError is an error wrapping multiple
// validation errors returned by DeleteCasesByProblemIdReply.ValidateAll() if
// the designated constraints aren't met.
type DeleteCasesByProblemIdReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCasesByProblemIdReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCasesByProblemIdReplyMultiError) AllErrors() []error { return m }

// DeleteCasesByProblemIdReplyValidationError is the validation error returned
// by DeleteCasesByProblemIdReply.Validate if the designated constraints
// aren't met.
type DeleteCasesByProblemIdReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCasesByProblemIdReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCasesByProblemIdReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCasesByProblemIdReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCasesByProblemIdReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCasesByProblemIdReplyValidationError) ErrorName() string {
	return "DeleteCasesByProblemIdReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCasesByProblemIdReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCasesByProblemIdReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCasesByProblemIdReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCasesByProblemIdReplyValidationError{}
