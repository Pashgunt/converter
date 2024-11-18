package main

import (
	"fmt"
	"github.com/Pashgunt/converter/internal/converter"
	"github.com/Pashgunt/converter/internal/enum"
)

type Inner struct {
	Locate string `json:"locate" group:"inner__short"`
	Outer  *Outer `json:"outer" group:"inner__outer"`
}

type Outer struct {
	Abc string `json:"abc" group:"inner__short"`
}

type Test struct {
	Name  string `json:"name" group:"test__full"`
	Data  string `json:"data" group:"test__full"`
	Inner *Inner `json:"inner" group:"test__inner"`
}

func main() {
	jsonData := `{"name": "Test Name", "data": "Detailed Data", "inner": {"locate": "SPB","outer": {"abc": "SPB2"}}}`
	var tShort Test
	converter.Convert(jsonData, &tShort, map[string][]string{enum.ContextGroup: []string{"test_data_group"}})
	fmt.Println(tShort)
}
