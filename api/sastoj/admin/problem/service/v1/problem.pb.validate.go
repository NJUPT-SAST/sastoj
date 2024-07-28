// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sastoj/admin/problem/service/v1/problem.proto

package problem

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

// Validate checks the field values on CreateProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProblemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProblemRequestMultiError, or nil if none found.
func (m *CreateProblemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProblemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Point

	// no validation rules for ContestId

	// no validation rules for CaseVersion

	// no validation rules for Index

	// no validation rules for Config

	// no validation rules for OwnerId

	// no validation rules for Visibility

	if len(errors) > 0 {
		return CreateProblemRequestMultiError(errors)
	}

	return nil
}

// CreateProblemRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProblemRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProblemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProblemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProblemRequestMultiError) AllErrors() []error { return m }

// CreateProblemRequestValidationError is the validation error returned by
// CreateProblemRequest.Validate if the designated constraints aren't met.
type CreateProblemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProblemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProblemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProblemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProblemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProblemRequestValidationError) ErrorName() string {
	return "CreateProblemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProblemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProblemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProblemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProblemRequestValidationError{}

// Validate checks the field values on CreateProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProblemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProblemReplyMultiError, or nil if none found.
func (m *CreateProblemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProblemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateProblemReplyMultiError(errors)
	}

	return nil
}

// CreateProblemReplyMultiError is an error wrapping multiple validation errors
// returned by CreateProblemReply.ValidateAll() if the designated constraints
// aren't met.
type CreateProblemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProblemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProblemReplyMultiError) AllErrors() []error { return m }

// CreateProblemReplyValidationError is the validation error returned by
// CreateProblemReply.Validate if the designated constraints aren't met.
type CreateProblemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProblemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProblemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProblemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProblemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProblemReplyValidationError) ErrorName() string {
	return "CreateProblemReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProblemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProblemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProblemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProblemReplyValidationError{}

// Validate checks the field values on UpdateProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProblemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProblemRequestMultiError, or nil if none found.
func (m *UpdateProblemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProblemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Point

	// no validation rules for ContestId

	// no validation rules for CaseVersion

	// no validation rules for Index

	// no validation rules for Config

	// no validation rules for OwnerId

	// no validation rules for Visibility

	if len(errors) > 0 {
		return UpdateProblemRequestMultiError(errors)
	}

	return nil
}

// UpdateProblemRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateProblemRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateProblemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProblemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProblemRequestMultiError) AllErrors() []error { return m }

// UpdateProblemRequestValidationError is the validation error returned by
// UpdateProblemRequest.Validate if the designated constraints aren't met.
type UpdateProblemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProblemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProblemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProblemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProblemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProblemRequestValidationError) ErrorName() string {
	return "UpdateProblemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProblemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProblemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProblemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProblemRequestValidationError{}

// Validate checks the field values on UpdateProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProblemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProblemReplyMultiError, or nil if none found.
func (m *UpdateProblemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProblemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return UpdateProblemReplyMultiError(errors)
	}

	return nil
}

// UpdateProblemReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateProblemReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateProblemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProblemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProblemReplyMultiError) AllErrors() []error { return m }

// UpdateProblemReplyValidationError is the validation error returned by
// UpdateProblemReply.Validate if the designated constraints aren't met.
type UpdateProblemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProblemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProblemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProblemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProblemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProblemReplyValidationError) ErrorName() string {
	return "UpdateProblemReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProblemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProblemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProblemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProblemReplyValidationError{}

// Validate checks the field values on DeleteProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProblemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProblemRequestMultiError, or nil if none found.
func (m *DeleteProblemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProblemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProblemRequestMultiError(errors)
	}

	return nil
}

// DeleteProblemRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteProblemRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteProblemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProblemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProblemRequestMultiError) AllErrors() []error { return m }

// DeleteProblemRequestValidationError is the validation error returned by
// DeleteProblemRequest.Validate if the designated constraints aren't met.
type DeleteProblemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProblemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProblemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProblemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProblemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProblemRequestValidationError) ErrorName() string {
	return "DeleteProblemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProblemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProblemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProblemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProblemRequestValidationError{}

// Validate checks the field values on DeleteProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProblemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProblemReplyMultiError, or nil if none found.
func (m *DeleteProblemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProblemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return DeleteProblemReplyMultiError(errors)
	}

	return nil
}

// DeleteProblemReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteProblemReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteProblemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProblemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProblemReplyMultiError) AllErrors() []error { return m }

// DeleteProblemReplyValidationError is the validation error returned by
// DeleteProblemReply.Validate if the designated constraints aren't met.
type DeleteProblemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProblemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProblemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProblemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProblemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProblemReplyValidationError) ErrorName() string {
	return "DeleteProblemReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProblemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProblemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProblemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProblemReplyValidationError{}

// Validate checks the field values on GetProblemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProblemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProblemRequestMultiError, or nil if none found.
func (m *GetProblemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProblemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProblemRequestMultiError(errors)
	}

	return nil
}

// GetProblemRequestMultiError is an error wrapping multiple validation errors
// returned by GetProblemRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProblemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProblemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProblemRequestMultiError) AllErrors() []error { return m }

// GetProblemRequestValidationError is the validation error returned by
// GetProblemRequest.Validate if the designated constraints aren't met.
type GetProblemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProblemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProblemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProblemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProblemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProblemRequestValidationError) ErrorName() string {
	return "GetProblemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProblemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProblemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProblemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProblemRequestValidationError{}

// Validate checks the field values on GetProblemReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProblemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProblemReplyMultiError, or nil if none found.
func (m *GetProblemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProblemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Point

	// no validation rules for ContestId

	// no validation rules for CaseVersion

	// no validation rules for Index

	// no validation rules for Config

	// no validation rules for OwnerId

	// no validation rules for Visibility

	if len(errors) > 0 {
		return GetProblemReplyMultiError(errors)
	}

	return nil
}

// GetProblemReplyMultiError is an error wrapping multiple validation errors
// returned by GetProblemReply.ValidateAll() if the designated constraints
// aren't met.
type GetProblemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProblemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProblemReplyMultiError) AllErrors() []error { return m }

// GetProblemReplyValidationError is the validation error returned by
// GetProblemReply.Validate if the designated constraints aren't met.
type GetProblemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProblemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProblemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProblemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProblemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProblemReplyValidationError) ErrorName() string { return "GetProblemReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetProblemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProblemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProblemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProblemReplyValidationError{}

// Validate checks the field values on ListProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProblemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProblemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProblemRequestMultiError, or nil if none found.
func (m *ListProblemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProblemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Size

	// no validation rules for Currency

	if len(errors) > 0 {
		return ListProblemRequestMultiError(errors)
	}

	return nil
}

// ListProblemRequestMultiError is an error wrapping multiple validation errors
// returned by ListProblemRequest.ValidateAll() if the designated constraints
// aren't met.
type ListProblemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProblemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProblemRequestMultiError) AllErrors() []error { return m }

// ListProblemRequestValidationError is the validation error returned by
// ListProblemRequest.Validate if the designated constraints aren't met.
type ListProblemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProblemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProblemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProblemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProblemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProblemRequestValidationError) ErrorName() string {
	return "ListProblemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProblemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProblemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProblemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProblemRequestValidationError{}

// Validate checks the field values on ListProblemReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListProblemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProblemReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProblemReplyMultiError, or nil if none found.
func (m *ListProblemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProblemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Currency

	// no validation rules for Total

	for idx, item := range m.GetProblems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProblemReplyValidationError{
						field:  fmt.Sprintf("Problems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProblemReplyValidationError{
						field:  fmt.Sprintf("Problems[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProblemReplyValidationError{
					field:  fmt.Sprintf("Problems[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListProblemReplyMultiError(errors)
	}

	return nil
}

// ListProblemReplyMultiError is an error wrapping multiple validation errors
// returned by ListProblemReply.ValidateAll() if the designated constraints
// aren't met.
type ListProblemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProblemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProblemReplyMultiError) AllErrors() []error { return m }

// ListProblemReplyValidationError is the validation error returned by
// ListProblemReply.Validate if the designated constraints aren't met.
type ListProblemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProblemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProblemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProblemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProblemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProblemReplyValidationError) ErrorName() string { return "ListProblemReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListProblemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProblemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProblemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProblemReplyValidationError{}

// Validate checks the field values on ListProblemReply_Problem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProblemReply_Problem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProblemReply_Problem with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProblemReply_ProblemMultiError, or nil if none found.
func (m *ListProblemReply_Problem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProblemReply_Problem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Content

	// no validation rules for Point

	// no validation rules for ContestId

	// no validation rules for CaseVersion

	// no validation rules for Index

	// no validation rules for Config

	// no validation rules for OwnerId

	// no validation rules for Visibility

	if len(errors) > 0 {
		return ListProblemReply_ProblemMultiError(errors)
	}

	return nil
}

// ListProblemReply_ProblemMultiError is an error wrapping multiple validation
// errors returned by ListProblemReply_Problem.ValidateAll() if the designated
// constraints aren't met.
type ListProblemReply_ProblemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProblemReply_ProblemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProblemReply_ProblemMultiError) AllErrors() []error { return m }

// ListProblemReply_ProblemValidationError is the validation error returned by
// ListProblemReply_Problem.Validate if the designated constraints aren't met.
type ListProblemReply_ProblemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProblemReply_ProblemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProblemReply_ProblemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProblemReply_ProblemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProblemReply_ProblemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProblemReply_ProblemValidationError) ErrorName() string {
	return "ListProblemReply_ProblemValidationError"
}

// Error satisfies the builtin error interface
func (e ListProblemReply_ProblemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProblemReply_Problem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProblemReply_ProblemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProblemReply_ProblemValidationError{}
