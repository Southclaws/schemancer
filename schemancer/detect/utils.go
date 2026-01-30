package detect

import (
	"fmt"
	"reflect"
)

func getConstValue(v interface{}) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return ""
		}
		rv = rv.Elem()
	}

	return fmt.Sprintf("%v", rv.Interface())
}
