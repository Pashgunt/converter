## Why?

The library was invented at the moment when I faced the following problem. I'll try to describe it briefly.  We may have a different set of data for different operations, for example, to work with one entity (which is especially common in web development), but we do not want to create a new structure for each operation and deserialize it. To do this, the concept was invented to declare the structure once and regulate a specific set of fields through the group attributes of the fields. In short , this is so ...

## Using

#### Connect the library:

```bash
  go get github.com/Pashgunt/converter@v1.0.0
```

#### Usage example:

```go
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
```

This usage example provides examples of using groups as structures declared through the group comment, as well as a more complex approach that can be useful for complex structures that describes groups in a file and then simply use them.
Examples of using groups in files are described in `example/group` as well as in tests.

If you plan to use the groups described in the file, then you need to create a .env file in the project and write it in the GROUP_DIR variable to the folder with the group files and connect the .env file using the code:

```go
environment := serializer.Environment{}

if _, err := environment.Load("./.env.serializer.example"); err != nil {
    panic(err)
}
```

To get the groups that are generated from the file you described, you can use the code:

```go
groups, err := serializer.GetGroups([]string{})
```

## Example

Example of a JSON input structure

```json
{
  "name": "Test Name",
  "data": "Test Data",
  "inner": {
    "locate": "Test Inner Locate",
    "outer": {
      "ref": "Test Outer Ref"
    }
  }
}
```

Description of the JSON structure on Go, taking into account groups

```go
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
```

As a result, if only the `test__full` group is passed to `Convert`, then only those fields that have `group = test__full` are converted to the Test structure.
At the same time, the description of the JSON structure will remain complete and the filling of fields for different tasks can be adjusted using `group`.

````go
	_ = serializer.Convert(
		`{"name": "Test Name", "data": "Test Data", "inner": {"locate": "Test Inner Locate","outer": {"ref": "Test Outer Ref"}}}`,
		&test,
		map[string][]string{
			serializer.ContextGroup:       {"test__full"},
			serializer.ContextEnvironment: {environment.GetGroupDir()},
		},
	)
````