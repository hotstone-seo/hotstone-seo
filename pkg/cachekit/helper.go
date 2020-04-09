package cachekit

import (
	"strings"
)

// NotModifiedError return true if error is not modified error
func NotModifiedError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "not modified")
}
