package main

import (
	"fmt"
	"serializer/internal/converter"
)

type Inner struct {
	Locate string `json:"locate" group:"short"`
}

type Test struct {
	Name  string `json:"name" group:"full"`
	Data  string `json:"data" group:"full"`
	Inner *Inner `json:"inner" group:"full"`
}

func main() {
	jsonData := `{"name": "Test Name", "data": "Detailed Data", "inner": {"locate": "SPB"}}`

	var tShort Test
	converter.Convert(jsonData, &tShort, "full")
	fmt.Println(tShort)
}
