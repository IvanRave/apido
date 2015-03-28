package apido

import (
	"fmt"
)

type SomeThing struct {
	IsImportant bool `json:"is_important" summary:"maybe"`
}

// serv_group with locale
type DemoStruct struct {
	Id           int32        `json:"id" summary:"unique number"`
	Name         string       `json:"name" summary:"name of object"`
	ArrSomeThing []*SomeThing `json:"arr_some_thing" summary:"array of some things"`
}

func ExampleToSwag() {

	result := ToSwag(DemoStruct{})

	fmt.Println(result["id"].SwagType)
	fmt.Println(result["id"].SwagFormat)
	fmt.Println(result["name"].SwagType)
	fmt.Println(result["arr_some_thing"].SwagType)
	// name convention for child elements: arr_child_name
	fmt.Println(result["arr_some_thing"].ArrItem.RefParam)

	// OUTPUT: integer
	// int32
	// string
	// array
	// some_thing
}
