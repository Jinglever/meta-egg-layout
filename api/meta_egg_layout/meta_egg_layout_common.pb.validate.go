// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: meta_egg_layout_common.proto

package meta_egg_layout

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

// Validate checks the field values on Pagination with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pagination with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginationMultiError, or
// nil if none found.
func (m *Pagination) ValidateAll() error {
	return m.validate(true)
}

func (m *Pagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PaginationValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val >= 100 {
		err := PaginationValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 100)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PaginationMultiError(errors)
	}

	return nil
}

// PaginationMultiError is an error wrapping multiple validation errors
// returned by Pagination.ValidateAll() if the designated constraints aren't met.
type PaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationMultiError) AllErrors() []error { return m }

// PaginationValidationError is the validation error returned by
// Pagination.Validate if the designated constraints aren't met.
type PaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationValidationError) ErrorName() string { return "PaginationValidationError" }

// Error satisfies the builtin error interface
func (e PaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationValidationError{}