package infrastructure

import (
	"serializer/internal/entity"
	"sync"
)

var ParamPool = sync.Pool{
	New: func() interface{} { return &entity.Param{} },
}
