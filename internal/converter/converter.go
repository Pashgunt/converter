package converter

import (
	"encoding/json"
	"github.com/Pashgunt/converter/internal/closure"
	"github.com/Pashgunt/converter/internal/entity"
	"github.com/Pashgunt/converter/internal/enum"
	"github.com/Pashgunt/converter/internal/helper"
	"github.com/Pashgunt/converter/internal/infrastructure"
	"github.com/Pashgunt/converter/internal/reflect"
	"github.com/Pashgunt/converter/internal/unmarshal"
	corereflect "reflect"
	"slices"
)

func Convert[TData helper.DataConstraint, TGroups helper.GroupConstraint](
	data TData,
	object interface{},
	context map[string]TGroups,
) error {
	if _, isset := context[enum.ContextGroup]; !isset {
		err := json.Unmarshal(helper.PrepareData(data), object)

		if err != nil {
			return err
		}
	}

	rawData, err := unmarshal.Decode(helper.PrepareData(data))

	if err != nil {
		return err
	}

	reflectObject, reflectType := reflect.PrepareValType(object)

	param := infrastructure.ParamPool.
		Get().(*entity.Param).
		Init(
			helper.PrepareGroups(context[enum.ContextGroup]),
			rawData,
			reflectObject,
			reflectType,
		)

	for numField := 0; numField < reflectType.NumField(); numField++ {
		if isClosure(*param, numField) {
			closure.InitStructure(*param, numField)

			return Convert(
				closure.GetInData(*param, numField),
				reflectObject.Field(numField).Interface(),
				map[string][]string{
					enum.ContextGroup: param.InGroups(),
				},
			)
		}

		if value, isSet := isSetValue(*param, numField); isSet {
			reflectObject.Field(numField).Set(corereflect.ValueOf(value).Convert(reflectType.Field(numField).Type))
		}
	}

	infrastructure.ParamPool.Put(param)

	return nil
}

func isSetValue(param entity.Param, index int) (interface{}, bool) {
	if slices.Contains(param.InGroups(), reflect.GetGroupTag(param.ReflectType().Field(index))) {
		if jsonValue, exists := param.RawData()[reflect.GetJsonTag(param.ReflectType().Field(index))]; exists {
			return jsonValue, true
		}
	}

	return nil, false
}

func isClosure(param entity.Param, index int) bool {
	if param.ReflectType().Field(index).Type.Kind() != corereflect.Ptr ||
		param.ReflectType().Field(index).Type.Elem().Kind() != corereflect.Struct {
		return false
	}

	if _, isSet := isSetValue(param, index); !isSet {
		return false
	}

	return true
}

type Abc struct {
	Test string
}
