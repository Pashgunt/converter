package main

import (
	"github.com/Pashgunt/converter/internal/converter"
	"github.com/Pashgunt/converter/internal/enum"
)

type Inner struct {
	Locate string `json:"locate" group:"inner__short"`
	Outer  *Outer `json:"outer" group:"inner__outer"`
}

type Outer struct {
	Ref string `json:"ref" group:"inner__short"`
}

type Test struct {
	Name  string `json:"name" group:"test__full"`
	Data  string `json:"data" group:"test__full"`
	Inner *Inner `json:"inner" group:"test__inner"`
}

func main() {
	var test Test

	_ = converter.Convert(
		`{"name": "Test Name", "data": "Test Data", "inner": {"locate": "Test Inner Locate","outer": {"ref": "Test Outer Ref"}}}`,
		test,
		map[string][]string{enum.ContextGroup: {"test__full", "example1", "example6", "inner__outer"}},
	)
}
