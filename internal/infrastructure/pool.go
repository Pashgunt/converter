package infrastructure

import (
	"github.com/Pashgunt/converter/internal/entity"
	"sync"
)

var ParamPool = sync.Pool{
	New: func() interface{} { return &entity.Param{} },
}
