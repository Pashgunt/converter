package reflect

import (
	"github.com/Pashgunt/converter/internal/enum"
	"reflect"
)

func PrepareValType(value interface{}) (reflect.Value, reflect.Type) {
	reflectValue := reflect.ValueOf(value).Elem()

	return reflectValue, reflectValue.Type()
}

func GetGroupTag(field reflect.StructField) string {
	return field.Tag.Get(enum.Group)
}

func GetJsonTag(field reflect.StructField) string {
	return field.Tag.Get(enum.Json)
}
