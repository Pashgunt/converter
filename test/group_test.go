package test

import (
	"github.com/Pashgunt/converter"
	"reflect"
	"testing"
)

func TestGetGroups(t *testing.T) {
	type args struct {
		groupAlias []string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Example1",
			args:    args{groupAlias: []string{"example1"}},
			want:    []string{"example1", "test__short", "test__full", "test__inner", "test__outer"},
			wantErr: false,
		},
		{
			name:    "Example2",
			args:    args{groupAlias: []string{"example2"}},
			want:    []string{"example2", "inner__short", "inner__full"},
			wantErr: false,
		},
		{
			name:    "Example3",
			args:    args{groupAlias: []string{"example3"}},
			want:    []string{"example3", "test__short", "test__full", "test__inner", "test__outer", "inner__short", "inner__full"},
			wantErr: false,
		},
		{
			name:    "Example4",
			args:    args{groupAlias: []string{"example4"}},
			want:    []string{"example4", "test__short", "test__full", "test__inner", "test__outer", "outer__short", "outer__full", "inner__short", "inner__full"},
			wantErr: false,
		},
		{
			name:    "Example5",
			args:    args{groupAlias: []string{"example5"}},
			want:    []string{"example5", "test__short", "test__full", "test__inner", "test__outer", "inner__short", "inner__full"},
			wantErr: false,
		},
		{
			name:    "Example7",
			args:    args{groupAlias: []string{"example6"}},
			want:    []string{"example6", "test__short", "test__full", "test__inner", "test__outer", "outer__short", "outer__full", "inner__short", "inner__full"},
			wantErr: false,
		},
	}

	environment := serializer.Environment{}

	if _, err := environment.Load("./.env.test.serializer.example"); err != nil {
		panic(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := serializer.GetGroups(tt.args.groupAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroups() got = %v, want %v", got, tt.want)
			}
		})
	}
}
