// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/v3/path.proto

package envoy_type_matcher_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _path_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on PathMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PathMatcher) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *PathMatcher_Path:

		if m.GetPath() == nil {
			return PathMatcherValidationError{
				field:  "Path",
				reason: "value is required",
			}
		}

		{
			tmp := m.GetPath()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return PathMatcherValidationError{
						field:  "Path",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return PathMatcherValidationError{
			field:  "Rule",
			reason: "value is required",
		}

	}

	return nil
}

// PathMatcherValidationError is the validation error returned by
// PathMatcher.Validate if the designated constraints aren't met.
type PathMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PathMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PathMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PathMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PathMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PathMatcherValidationError) ErrorName() string { return "PathMatcherValidationError" }

// Error satisfies the builtin error interface
func (e PathMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPathMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PathMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PathMatcherValidationError{}
