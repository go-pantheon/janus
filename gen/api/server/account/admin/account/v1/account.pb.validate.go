// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/admin/account/v1/account.proto

package adminv1

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

// Validate checks the field values on GetByIdRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetByIdRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetByIdRequestMultiError,
// or nil if none found.
func (m *GetByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetByIdRequestMultiError(errors)
	}

	return nil
}

// GetByIdRequestMultiError is an error wrapping multiple validation errors
// returned by GetByIdRequest.ValidateAll() if the designated constraints
// aren't met.
type GetByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetByIdRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetByIdRequestMultiError) AllErrors() []error { return m }

// GetByIdRequestValidationError is the validation error returned by
// GetByIdRequest.Validate if the designated constraints aren't met.
type GetByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIdRequestValidationError) ErrorName() string { return "GetByIdRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIdRequestValidationError{}

// Validate checks the field values on GetByIdResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetByIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetByIdResponseMultiError, or nil if none found.
func (m *GetByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetAccount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetByIdResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetByIdResponseValidationError{
					field:  "Account",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetByIdResponseValidationError{
				field:  "Account",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetByIdResponseMultiError(errors)
	}

	return nil
}

// GetByIdResponseMultiError is an error wrapping multiple validation errors
// returned by GetByIdResponse.ValidateAll() if the designated constraints
// aren't met.
type GetByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetByIdResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetByIdResponseMultiError) AllErrors() []error { return m }

// GetByIdResponseValidationError is the validation error returned by
// GetByIdResponse.Validate if the designated constraints aren't met.
type GetByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByIdResponseValidationError) ErrorName() string { return "GetByIdResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByIdResponseValidationError{}

// Validate checks the field values on AccountListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AccountListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListRequestMultiError, or nil if none found.
func (m *AccountListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if all {
		switch v := interface{}(m.GetCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccountListRequestValidationError{
					field:  "Condition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccountListRequestValidationError{
					field:  "Condition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountListRequestValidationError{
				field:  "Condition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccountListRequestMultiError(errors)
	}

	return nil
}

// AccountListRequestMultiError is an error wrapping multiple validation errors
// returned by AccountListRequest.ValidateAll() if the designated constraints
// aren't met.
type AccountListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListRequestMultiError) AllErrors() []error { return m }

// AccountListRequestValidationError is the validation error returned by
// AccountListRequest.Validate if the designated constraints aren't met.
type AccountListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListRequestValidationError) ErrorName() string {
	return "AccountListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AccountListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListRequestValidationError{}

// Validate checks the field values on AccountListCondition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AccountListCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListCondition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListConditionMultiError, or nil if none found.
func (m *AccountListCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for AppleId

	// no validation rules for GoogleId

	if len(errors) > 0 {
		return AccountListConditionMultiError(errors)
	}

	return nil
}

// AccountListConditionMultiError is an error wrapping multiple validation
// errors returned by AccountListCondition.ValidateAll() if the designated
// constraints aren't met.
type AccountListConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListConditionMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListConditionMultiError) AllErrors() []error { return m }

// AccountListConditionValidationError is the validation error returned by
// AccountListCondition.Validate if the designated constraints aren't met.
type AccountListConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListConditionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListConditionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListConditionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListConditionValidationError) ErrorName() string {
	return "AccountListConditionValidationError"
}

// Error satisfies the builtin error interface
func (e AccountListConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListConditionValidationError{}

// Validate checks the field values on AccountListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AccountListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AccountListResponseMultiError, or nil if none found.
func (m *AccountListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	for idx, item := range m.GetAccounts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AccountListResponseValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AccountListResponseValidationError{
						field:  fmt.Sprintf("Accounts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccountListResponseValidationError{
					field:  fmt.Sprintf("Accounts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return AccountListResponseMultiError(errors)
	}

	return nil
}

// AccountListResponseMultiError is an error wrapping multiple validation
// errors returned by AccountListResponse.ValidateAll() if the designated
// constraints aren't met.
type AccountListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountListResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountListResponseMultiError) AllErrors() []error { return m }

// AccountListResponseValidationError is the validation error returned by
// AccountListResponse.Validate if the designated constraints aren't met.
type AccountListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountListResponseValidationError) ErrorName() string {
	return "AccountListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AccountListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountListResponseValidationError{}

// Validate checks the field values on AccountProto with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccountProto) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccountProto with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccountProtoMultiError, or
// nil if none found.
func (m *AccountProto) ValidateAll() error {
	return m.validate(true)
}

func (m *AccountProto) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IdStr

	// no validation rules for AppleId

	// no validation rules for GoogleId

	// no validation rules for Name

	// no validation rules for RegisterIp

	// no validation rules for LastLoginIp

	// no validation rules for Channel

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return AccountProtoMultiError(errors)
	}

	return nil
}

// AccountProtoMultiError is an error wrapping multiple validation errors
// returned by AccountProto.ValidateAll() if the designated constraints aren't met.
type AccountProtoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccountProtoMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccountProtoMultiError) AllErrors() []error { return m }

// AccountProtoValidationError is the validation error returned by
// AccountProto.Validate if the designated constraints aren't met.
type AccountProtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountProtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountProtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountProtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountProtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountProtoValidationError) ErrorName() string { return "AccountProtoValidationError" }

// Error satisfies the builtin error interface
func (e AccountProtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountProto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountProtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountProtoValidationError{}
