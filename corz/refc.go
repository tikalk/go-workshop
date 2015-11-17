// Reflection
//
// Author: Dmitri Krasnenko

package corz

import (
	"reflect"
	"strings"
)

func parseTag(tag string) string {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}

func namesForType(t reflect.Type) (n []string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			name := strings.ToLower(f.Name)

			if al := parseTag(f.Tag.Get("json")); al != "" && al != "-" {
				name = al
			}

			n = append(n, name)
			n = append(n, namesForType(f.Type)...)
		}
	}

	return
}

//
//The method prints out declared field names of any structure
//
func DeclaredFields(i interface{}) (n []string) {
	return namesForType(
		reflect.TypeOf(
			i,
		),
	)
}
