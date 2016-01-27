package common

import (
	"errors"
)

var (

	ErrInvalidParameter = errors.New(`Invalid parameter string encountered.`)

	ErrMandatoryParameterMissing = errors.New(`Mandatory Parameter missing.`)

	ErrGeneral = errors.New(`Something went wrong`)


)
