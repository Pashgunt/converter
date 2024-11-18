package main

import (
	serializer "github.com/Pashgunt/converter"
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

	environment := serializer.Environment{}

	if _, err := environment.Load("./.env.serializer.example"); err != nil {
		panic(err)
	}

	_ = serializer.Convert(
		`{"name": "Test Name", "data": "Test Data", "inner": {"locate": "Test Inner Locate","outer": {"ref": "Test Outer Ref"}}}`,
		&test,
		map[string][]string{
			serializer.ContextGroup:       {"test__full", "example1", "example6", "inner__outer"},
			serializer.ContextEnvironment: {environment.GetGroupDir()},
		},
	)
}
