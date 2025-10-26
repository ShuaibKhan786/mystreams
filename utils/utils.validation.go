package utils

import (
	"reflect"
	"slices"
)

// it will automatically validate
// if struct field have validation tagging with "required"
// e.g `validation:"required"`
// or you can specify optionals and remove "required" tagging
// e.g `validation:"optional"` and just specify in the args
// and make sure struct field is pointer
func Validate(body interface{}, optionals ...string) bool {
	t := reflect.TypeOf(body)
	k := t.Kind()
	v := reflect.ValueOf(body)

	if k == reflect.Pointer {
		t = t.Elem()
		k = t.Kind()
		v = v.Elem()
		if k != reflect.Struct {
			return false
		}
	}

	if k != reflect.Struct {
		return false
	}

	n := t.NumField()
	shouldCheck := false
	for i := range n - 1 {
		sft := t.Field(i)
		sfv := v.Field(i)

		tag := sft.Tag.Get("validation")
		if tag == "" {
			continue
		}
		shouldCheck = tag == "required" || slices.Contains(optionals, tag)

		if shouldCheck && sfv.Kind() == reflect.Pointer {
			switch sfv.IsNil() {
			case true:
				return false
			case false:
				fv := sfv.Elem()
				if fv.Kind() == reflect.String && fv.Len() == 0 {
					return false
				}
			}
		}

		shouldCheck = false
	}

	return true
}
