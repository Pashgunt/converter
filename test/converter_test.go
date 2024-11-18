package test

import (
	"github.com/Pashgunt/converter"
	"github.com/Pashgunt/converter/internal/helper"
	"testing"
)

func TestConvert(t *testing.T) {
	type Outer struct {
		Ref string `json:"ref" group:"inner__short"`
	}

	type Inner struct {
		Locate string `json:"locate" group:"inner__short"`
		Outer  *Outer `json:"outer" group:"inner__outer"`
	}

	type Test struct {
		Name  string `json:"name" group:"test__full"`
		Data  string `json:"data" group:"test__full"`
		Inner *Inner `json:"inner" group:"test__inner"`
	}

	var testData Test

	type args[TData helper.DataConstraint, TGroups helper.GroupConstraint] struct {
		data    TData
		object  interface{}
		context map[string]TGroups
	}
	type testCase[TData helper.DataConstraint, TGroups helper.GroupConstraint] struct {
		name    string
		args    args[TData, TGroups]
		wantErr bool
	}

	environment := serializer.Environment{}

	if _, err := environment.Load("./.env.test.serializer.example"); err != nil {
		panic(err)
	}

	tests := []testCase[string, []string]{
		{
			name: "Converter",
			args: args[string, []string]{
				data:   `{"name": "Test Name", "data": "Test Data", "inner": {"locate": "Test Inner Locate", "outer": {"ref": "Test Outer Ref"}}}`,
				object: &testData,
				context: map[string][]string{
					serializer.ContextGroup: {"example6"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := serializer.Convert(tt.args.data, tt.args.object, tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
