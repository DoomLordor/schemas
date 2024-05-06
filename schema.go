package schemas

import (
	"net/url"
	"reflect"
)

const (
	tagSchema        = "schema"
	tagSchemaDefault = "default"
	tagRequired      = "required"
	tagSkip          = "-"
	emptyTagSchema   = ""
	emptyValue       = ""
	requiredTrue     = "true"
)

type field interface {
	Parse(raw, def string) error
}

func Parse(obj any, query url.Values) error {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	} else {
		return PointerError
	}

	if v.Kind() != reflect.Struct {
		return TypeError
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldStruct := t.Field(i)
		name := fieldStruct.Tag.Get(tagSchema)
		def := fieldStruct.Tag.Get(tagSchemaDefault)
		required := fieldStruct.Tag.Get(tagRequired)

		if name == tagSkip {
			continue
		}

		fieldValue := v.Field(i)

		if name == emptyTagSchema {
			if fieldValue.Kind() == reflect.Pointer && fieldValue.Elem().Kind() == reflect.Struct {
				err := Parse(fieldValue.Interface(), query)
				if err != nil {
					return err
				}
			}
			continue
		}

		if fieldValue.Kind() == reflect.Pointer && fieldValue.IsNil() {
			newField := reflect.New(fieldValue.Type().Elem())
			fieldValue.Set(newField)
		}

		if fieldValue.Kind() != reflect.Pointer {
			fieldValue = fieldValue.Addr()
		}

		interfaceValue := fieldValue.Interface()
		f, ok := interfaceValue.(field)

		if !ok {
			return FieldError
		}

		val := query.Get(name)
		if val == emptyValue && required == requiredTrue {
			return NewRequiredError(name)
		}

		err := f.Parse(val, def)
		if err != nil {
			return NewParseError(name, val)
		}
	}
	return nil
}
