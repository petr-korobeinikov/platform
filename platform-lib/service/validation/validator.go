package validation

import (
	"errors"
	"fmt"
	"strings"
)

func EnsureServiceNameValid(serviceName string) error {
	// Use some validation library here.
	if (len(serviceName)) > 63 {
		return ErrServiceNameTooLong
	}

	if strings.Contains(serviceName, "_") {
		return ErrServiceNameContainsUnderscore
	}

	if strings.Contains(serviceName, ".") {
		return ErrServiceNameContainsDot
	}

	// starts alphabetic character
	// ends with valid character

	return nil
}

var (
	ErrServiceNameTooLong            = errors.New(fmt.Sprintf("service name too long, maximum length is %d characters", serviceNameMaxLength))
	ErrServiceNameContainsUnderscore = errors.New("service name contains underscore")
	ErrServiceNameContainsDot        = errors.New("service name contains dot")
)

const (
	serviceNameMaxLength = 63
)
