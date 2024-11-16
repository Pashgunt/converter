package closure

import (
	"encoding/json"
	"github.com/Pashgunt/converter/internal/entity"
	"github.com/Pashgunt/converter/internal/reflect"
	corereflect "reflect"
)

func InitStructure(param entity.Param, index int) {
	if !param.ReflectObject().Field(index).IsNil() {
		return
	}

	param.ReflectObject().Field(index).Set(corereflect.New(param.ReflectType().Field(index).Type.Elem()))
}

func GetInData(param entity.Param, index int) []byte {
	inData, _ := json.Marshal(param.RawData()[reflect.GetJsonTag(param.ReflectType().Field(index))])

	return inData
}
