package cachekit

import (
	"strings"
)

// NoModifiedError return true if error is not modified error
func NoModifiedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "not modified")
}
