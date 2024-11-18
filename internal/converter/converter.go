package converter

import (
	"encoding/json"
	"fmt"
	"github.com/Pashgunt/converter/internal/closure"
	"github.com/Pashgunt/converter/internal/entity"
	"github.com/Pashgunt/converter/internal/enum"
	"github.com/Pashgunt/converter/internal/exception"
	"github.com/Pashgunt/converter/internal/helper"
	"github.com/Pashgunt/converter/internal/infrastructure"
	"github.com/Pashgunt/converter/internal/reflect"
	"github.com/Pashgunt/converter/internal/resolver/alias/group"
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
		if err := json.Unmarshal(helper.PrepareData(data), object); err != nil {
			return err
		}
	}

	sliceOfGroups, err := group.GetGroups(helper.PrepareGroups(context[enum.ContextGroup]))

	if err != nil {
		return err
	}

	if err = process(data, object, sliceOfGroups); err != nil {
		return err
	}

	return nil
}

func process[TData helper.DataConstraint](data TData, object interface{}, sliceOfGroups []string) error {
	rawData, err := unmarshal.Decode(helper.PrepareData(data))

	if err != nil {
		return err
	}

	reflectObject, reflectType := reflect.PrepareValType(object)

	param, okConvert := infrastructure.ParamPool.Get().(*entity.Param)

	defer infrastructure.ParamPool.Put(param)

	if !okConvert {
		return fmt.Errorf(exception.ParamPoolException)
	}

	param.Init(
		sliceOfGroups,
		rawData,
		reflectObject,
		reflectType,
	)

	for numField := 0; numField < reflectType.NumField(); numField++ {
		if isClosure(*param, numField) {
			closure.InitStructure(*param, numField)

			if err = process(
				closure.GetInData(*param, numField),
				reflectObject.Field(numField).Interface(),
				param.InGroups(),
			); err != nil {
				return err
			}
		}

		field := reflectObject.Field(numField)

		if value, isSet := isSetValue(*param, numField); isSet &&
			field.IsValid() &&
			field.CanSet() &&
			corereflect.ValueOf(value).Type().ConvertibleTo(field.Type()) {
			field.Set(corereflect.ValueOf(value).Convert(field.Type()))
		}
	}

	return nil
}

func isSetValue(param entity.Param, index int) (interface{}, bool) {
	if slices.Contains(param.InGroups(), reflect.GetGroupTag(param.ReflectType().Field(index))) == false {
		return nil, false
	}

	if jsonValue, exists := param.RawData()[reflect.GetJsonTag(param.ReflectType().Field(index))]; exists {
		return jsonValue, true
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
