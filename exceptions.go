package schemas

import (
	"errors"
	"fmt"
)

var TypeError = errors.New("schema not a struct")
var PointerError = errors.New("schema not pointer")
var FieldError = errors.New("schema field not implementation interface `Field`")

type ParseFieldError struct {
	name  string
	value string
}

func (p *ParseFieldError) Error() string {
	return fmt.Sprintf(`parse field "%s" error from value "%s"`, p.name, p.value)
}

func NewParseError(name, value string) error {
	return &ParseFieldError{
		name:  name,
		value: value,
	}
}

type RequiredError struct {
	name string
}

func (r *RequiredError) Error() string {
	return fmt.Sprintf(`"%s" required`, r.name)
}

func NewRequiredError(name string) error {
	return &RequiredError{name: name}
}
