// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/recognize/v1/recognize-api.proto

package recognize_api

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

// Validate checks the field values on RecognizePhotoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RecognizePhotoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecognizePhotoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RecognizePhotoRequestMultiError, or nil if none found.
func (m *RecognizePhotoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RecognizePhotoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Image

	if len(errors) > 0 {
		return RecognizePhotoRequestMultiError(errors)
	}
	return nil
}

// RecognizePhotoRequestMultiError is an error wrapping multiple validation
// errors returned by RecognizePhotoRequest.ValidateAll() if the designated
// constraints aren't met.
type RecognizePhotoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecognizePhotoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecognizePhotoRequestMultiError) AllErrors() []error { return m }

// RecognizePhotoRequestValidationError is the validation error returned by
// RecognizePhotoRequest.Validate if the designated constraints aren't met.
type RecognizePhotoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecognizePhotoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecognizePhotoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecognizePhotoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecognizePhotoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecognizePhotoRequestValidationError) ErrorName() string {
	return "RecognizePhotoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RecognizePhotoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecognizePhotoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecognizePhotoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecognizePhotoRequestValidationError{}

// Validate checks the field values on RecognizePhotoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RecognizePhotoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RecognizePhotoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RecognizePhotoResponseMultiError, or nil if none found.
func (m *RecognizePhotoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RecognizePhotoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Category

	if len(errors) > 0 {
		return RecognizePhotoResponseMultiError(errors)
	}
	return nil
}

// RecognizePhotoResponseMultiError is an error wrapping multiple validation
// errors returned by RecognizePhotoResponse.ValidateAll() if the designated
// constraints aren't met.
type RecognizePhotoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecognizePhotoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecognizePhotoResponseMultiError) AllErrors() []error { return m }

// RecognizePhotoResponseValidationError is the validation error returned by
// RecognizePhotoResponse.Validate if the designated constraints aren't met.
type RecognizePhotoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecognizePhotoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecognizePhotoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecognizePhotoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecognizePhotoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecognizePhotoResponseValidationError) ErrorName() string {
	return "RecognizePhotoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RecognizePhotoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecognizePhotoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecognizePhotoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecognizePhotoResponseValidationError{}