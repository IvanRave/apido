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
	OptProp      *string      `json:"opt_prop" summary:"Optional property"`

	OptInt          *int32       `json:"opt_int"`
	ItemUnderStruct *UnderStruct `json:"under_struct"`
}

type UnderStruct struct {
	VeryProp string `json:"very_prop" summary:"Very"`
}

func ExampleToSwag() {

	result := ToSwag(DemoStruct{})

	fmt.Println(result["id"].SwagType)
	fmt.Println(result["id"].SwagFormat)
	fmt.Println(result["name"].SwagType)
	fmt.Println(result["arr_some_thing"].SwagType)
	// name convention for child elements: arr_child_name
	fmt.Println(result["arr_some_thing"].ArrItem.RefParam)

	fmt.Println("==opt_prop==")
	fmt.Println(result["opt_prop"].SwagType)

	fmt.Println("==opt_int==")
	fmt.Println(result["opt_int"].SwagType)
	fmt.Println(result["opt_int"].SwagFormat)

	fmt.Println("==under_struct==")
	fmt.Println(result["under_struct"].RefParam)

	// OUTPUT: integer
	// int32
	// string
	// array
	// some_thing
	// ==opt_prop==
	// string
	// ==opt_int==
	// integer
	// int32
	// ==under_struct==
	// under_struct
}
