package errkit

import (
	"fmt"
	"strings"
)

const (
	validationErrPrefix = "Validation: "
	validationSep       = "\n"
)

// ValidationErr to create validation error from messages
func ValidationErr(msgs ...string) error {
	return fmt.Errorf("%s%s", validationErrPrefix, strings.Join(msgs, validationSep))
}

// WrapValidation to wrap error to validation error
func WrapValidation(err error) error {
	return fmt.Errorf("%s%w", validationErrPrefix, err)
}

// ValidationMessages return message of validations
func ValidationMessages(err error) (msgs []string) {
	if !IsValidationErr(err) {
		return
	}

	errMsg := err.Error()
	errMsg = strings.TrimSpace(errMsg[len(validationErrPrefix):])
	return strings.Split(errMsg, validationSep)
}

// IsValidationErr return true if err is validation error
func IsValidationErr(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), validationErrPrefix)
}
