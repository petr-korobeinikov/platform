package spec

import (
	"errors"
	"fmt"
)

var (
	ErrSpecFileDoesNotExists = errors.New(fmt.Sprintf("%s does not found in project directory", File))
	ErrSpecReading           = errors.New(fmt.Sprintf("can't read %s", File))
	ErrSpecInvalid           = errors.New(fmt.Sprintf("%s contains errors", File))
	ErrTaskNotDefined        = errors.New(fmt.Sprintf("task not defined in file %s", File))
)
