package schemas

import (
	"strconv"
	"strings"
	"time"
)

const (
	timeNow     = "now"
	layout      = "2006-01-02T15:04:05Z"
	layoutMilli = "2006-01-02T15:04:05.000Z"
)

type BoolFiled struct {
	value bool
}

func (f *BoolFiled) Parse(raw, def string) error {
	f.value = false

	if raw == emptyValue {
		raw = def
	}

	if raw == emptyValue {
		return nil
	}

	value, err := strconv.ParseBool(raw)
	if err != nil {
		return err
	}

	f.value = value
	return nil
}

func (f *BoolFiled) Value() bool {
	return f.value
}

type UintFiled struct {
	value uint64
}

func (f *UintFiled) Parse(raw, def string) error {
	f.value = 0

	if raw == emptyValue {
		raw = def
	}

	if raw == emptyValue {
		return nil
	}

	value, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return err
	}

	f.value = value
	return nil
}

func (f *UintFiled) Value() uint64 {
	return f.value
}

type IntFiled struct {
	value int64
}

func (f *IntFiled) Parse(raw, def string) error {
	f.value = 0

	if raw == emptyValue {
		raw = def
	}

	if raw == emptyValue {
		return nil
	}

	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return err
	}

	f.value = value
	return nil
}

func (f *IntFiled) Value() int64 {
	return f.value
}

type StringFiled struct {
	value string
}

func (f *StringFiled) Parse(raw, def string) error {
	f.value = raw

	if raw == emptyValue {
		f.value = def
	}

	return nil
}

func (f *StringFiled) Value() string {
	return f.value
}

type StringsFiled struct {
	value []string
}

func (f *StringsFiled) Parse(raw, def string) error {
	f.value = []string{}

	if raw == emptyValue {
		raw = def
	}

	raw = strings.ReplaceAll(raw, ",,", ",")

	f.value = strings.Split(raw, ",")

	return nil
}

func (f *StringsFiled) Value() []string {
	return f.value
}

type TimeFiled struct {
	value *time.Time
}

func (f *TimeFiled) Parse(raw, def string) error {
	f.value = nil

	if raw == emptyValue {
		if def == timeNow {
			now := time.Now().Truncate(time.Microsecond)
			f.value = &now
			return nil
		}
		raw = def
	}

	if raw == emptyValue {
		return nil
	}

	for _, format := range []string{layout, layoutMilli} {
		t, err := time.Parse(format, raw)
		if err == nil {
			f.value = &t
			return nil
		}
	}

	return nil
}

func (f *TimeFiled) SetValue(value *time.Time) {
	f.value = value
}

func (f *TimeFiled) Value() *time.Time {
	return f.value
}
