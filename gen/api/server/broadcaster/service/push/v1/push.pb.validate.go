// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: broadcaster/service/push/v1/push.proto

package servicev1

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

// Validate checks the field values on PushRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushRequestMultiError, or
// nil if none found.
func (m *PushRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PushRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBodies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PushRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PushRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PushRequestValidationError{
					field:  fmt.Sprintf("Bodies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Color

	// no validation rules for Sid

	// no validation rules for Uid

	if len(errors) > 0 {
		return PushRequestMultiError(errors)
	}

	return nil
}

// PushRequestMultiError is an error wrapping multiple validation errors
// returned by PushRequest.ValidateAll() if the designated constraints aren't met.
type PushRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushRequestMultiError) AllErrors() []error { return m }

// PushRequestValidationError is the validation error returned by
// PushRequest.Validate if the designated constraints aren't met.
type PushRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushRequestValidationError) ErrorName() string { return "PushRequestValidationError" }

// Error satisfies the builtin error interface
func (e PushRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushRequestValidationError{}

// Validate checks the field values on PushResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushResponseMultiError, or
// nil if none found.
func (m *PushResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PushResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pushed

	if len(errors) > 0 {
		return PushResponseMultiError(errors)
	}

	return nil
}

// PushResponseMultiError is an error wrapping multiple validation errors
// returned by PushResponse.ValidateAll() if the designated constraints aren't met.
type PushResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushResponseMultiError) AllErrors() []error { return m }

// PushResponseValidationError is the validation error returned by
// PushResponse.Validate if the designated constraints aren't met.
type PushResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushResponseValidationError) ErrorName() string { return "PushResponseValidationError" }

// Error satisfies the builtin error interface
func (e PushResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushResponseValidationError{}

// Validate checks the field values on MulticastRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MulticastRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MulticastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MulticastRequestMultiError, or nil if none found.
func (m *MulticastRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MulticastRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBodies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MulticastRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MulticastRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MulticastRequestValidationError{
					field:  fmt.Sprintf("Bodies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Color

	if len(errors) > 0 {
		return MulticastRequestMultiError(errors)
	}

	return nil
}

// MulticastRequestMultiError is an error wrapping multiple validation errors
// returned by MulticastRequest.ValidateAll() if the designated constraints
// aren't met.
type MulticastRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MulticastRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MulticastRequestMultiError) AllErrors() []error { return m }

// MulticastRequestValidationError is the validation error returned by
// MulticastRequest.Validate if the designated constraints aren't met.
type MulticastRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MulticastRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MulticastRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MulticastRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MulticastRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MulticastRequestValidationError) ErrorName() string { return "MulticastRequestValidationError" }

// Error satisfies the builtin error interface
func (e MulticastRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMulticastRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MulticastRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MulticastRequestValidationError{}

// Validate checks the field values on MulticastResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MulticastResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MulticastResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MulticastResponseMultiError, or nil if none found.
func (m *MulticastResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MulticastResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MulticastResponseMultiError(errors)
	}

	return nil
}

// MulticastResponseMultiError is an error wrapping multiple validation errors
// returned by MulticastResponse.ValidateAll() if the designated constraints
// aren't met.
type MulticastResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MulticastResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MulticastResponseMultiError) AllErrors() []error { return m }

// MulticastResponseValidationError is the validation error returned by
// MulticastResponse.Validate if the designated constraints aren't met.
type MulticastResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MulticastResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MulticastResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MulticastResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MulticastResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MulticastResponseValidationError) ErrorName() string {
	return "MulticastResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MulticastResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMulticastResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MulticastResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MulticastResponseValidationError{}

// Validate checks the field values on BroadcastRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BroadcastRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BroadcastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BroadcastRequestMultiError, or nil if none found.
func (m *BroadcastRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BroadcastRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBodies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BroadcastRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BroadcastRequestValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BroadcastRequestValidationError{
					field:  fmt.Sprintf("Bodies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Color

	// no validation rules for Sid

	if len(errors) > 0 {
		return BroadcastRequestMultiError(errors)
	}

	return nil
}

// BroadcastRequestMultiError is an error wrapping multiple validation errors
// returned by BroadcastRequest.ValidateAll() if the designated constraints
// aren't met.
type BroadcastRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BroadcastRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BroadcastRequestMultiError) AllErrors() []error { return m }

// BroadcastRequestValidationError is the validation error returned by
// BroadcastRequest.Validate if the designated constraints aren't met.
type BroadcastRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BroadcastRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BroadcastRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BroadcastRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BroadcastRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BroadcastRequestValidationError) ErrorName() string { return "BroadcastRequestValidationError" }

// Error satisfies the builtin error interface
func (e BroadcastRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBroadcastRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BroadcastRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BroadcastRequestValidationError{}

// Validate checks the field values on BroadcastResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BroadcastResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BroadcastResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BroadcastResponseMultiError, or nil if none found.
func (m *BroadcastResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BroadcastResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BroadcastResponseMultiError(errors)
	}

	return nil
}

// BroadcastResponseMultiError is an error wrapping multiple validation errors
// returned by BroadcastResponse.ValidateAll() if the designated constraints
// aren't met.
type BroadcastResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BroadcastResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BroadcastResponseMultiError) AllErrors() []error { return m }

// BroadcastResponseValidationError is the validation error returned by
// BroadcastResponse.Validate if the designated constraints aren't met.
type BroadcastResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BroadcastResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BroadcastResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BroadcastResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BroadcastResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BroadcastResponseValidationError) ErrorName() string {
	return "BroadcastResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BroadcastResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBroadcastResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BroadcastResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BroadcastResponseValidationError{}

// Validate checks the field values on PushMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushMessageMultiError, or
// nil if none found.
func (m *PushMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *PushMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBodies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PushMessageValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PushMessageValidationError{
						field:  fmt.Sprintf("Bodies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PushMessageValidationError{
					field:  fmt.Sprintf("Bodies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Broadcast

	// no validation rules for Color

	// no validation rules for Sid

	if len(errors) > 0 {
		return PushMessageMultiError(errors)
	}

	return nil
}

// PushMessageMultiError is an error wrapping multiple validation errors
// returned by PushMessage.ValidateAll() if the designated constraints aren't met.
type PushMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushMessageMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushMessageMultiError) AllErrors() []error { return m }

// PushMessageValidationError is the validation error returned by
// PushMessage.Validate if the designated constraints aren't met.
type PushMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushMessageValidationError) ErrorName() string { return "PushMessageValidationError" }

// Error satisfies the builtin error interface
func (e PushMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushMessageValidationError{}

// Validate checks the field values on PushBody with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PushBody) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PushBody with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PushBodyMultiError, or nil
// if none found.
func (m *PushBody) ValidateAll() error {
	return m.validate(true)
}

func (m *PushBody) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	// no validation rules for DataVersion

	// no validation rules for Obj

	// no validation rules for Mod

	// no validation rules for Seq

	if len(errors) > 0 {
		return PushBodyMultiError(errors)
	}

	return nil
}

// PushBodyMultiError is an error wrapping multiple validation errors returned
// by PushBody.ValidateAll() if the designated constraints aren't met.
type PushBodyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PushBodyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PushBodyMultiError) AllErrors() []error { return m }

// PushBodyValidationError is the validation error returned by
// PushBody.Validate if the designated constraints aren't met.
type PushBodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PushBodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PushBodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PushBodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PushBodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PushBodyValidationError) ErrorName() string { return "PushBodyValidationError" }

// Error satisfies the builtin error interface
func (e PushBodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPushBody.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PushBodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PushBodyValidationError{}
