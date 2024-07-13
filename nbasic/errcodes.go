package nbasic

import "errors"

var ErrInternalError = errors.New("internal error")

// JSON Errors
var ErrJSONParsing = errors.New("error when parsing JSON")
var ErrJSONValidate = errors.New("error when validate JSON")
