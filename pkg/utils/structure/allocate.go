package structure

import (
	"reflect"
	"regexp"
	"strings"
)

// Allocate allocate memory for all struct fields.
// OBS: the structure passed must be the memory address of a variable: `var object Strcut > &object` or; the passed structure must be the variable itself which already an empty pointer: `object := &Struct{} > object`.
func Allocate(structure any, skips ...string) {
	object := reflect.ValueOf(structure).Elem()
	for index := 0; index < object.NumField(); index++ {
		if !regexp.MustCompile(object.Type().Field(index).Name).MatchString(strings.Join(skips, " ")) {
			field := object.Field(index)
			if field.CanSet() {
				field.Set(reflect.New(field.Type().Elem()))
			}
		}
	}
}
