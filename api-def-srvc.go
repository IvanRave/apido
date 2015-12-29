package apido

import (
	"fmt"
	"reflect"
	"strings"
)

// calcInParam - calculate inner parameter from reflect values
func calcInParam(f reflect.StructField) (*InParam, string) {
	// http://golang.org/pkg/reflect/#StructTag
	// By convention, tag strings are a concatenation of
	// optionally space-separated key:"value" pairs.
	// Each key is a non-empty string consisting of non-control characters
	// other than space (U+0020 ' '), quote (U+0022 '"'), and colon (U+003A ':').
	// Each value is quoted using U+0022 '"' characters and Go string literal syntax.
	var st reflect.StructTag = f.Tag
	//qwer := st.Get("rqr")
	//fmt.Printf("%v", qwer)

	prm := &InParam{
		//SwagType: swagType,
		Description: st.Get("summary"),
	}

	// json="myval,string,omitempty"
	prmKey := st.Get("json")

	// only myval
	prmKey = strings.Split(prmKey, ",")[0]

	// It returns an empty string for unnamed types.
	//    like arrays of structs
	switch f.Type.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		prm.SwagType = "integer"
		prm.SwagFormat = f.Type.String()
	case reflect.String:
		prm.SwagType = "string"
	case reflect.Bool:
		prm.SwagType = "boolean"
	case reflect.Slice:
		prm.SwagType = "array"
		prm.ArrItem = &InParam{
			RefParam: prmKey[4:len(prmKey)], // remove arr_
		}
	case reflect.Ptr:

		// f.Type.Name() == "" empty for unnamed types
		// Elem() returns a type's element type.
		// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.

		switch f.Type.Elem().Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			prm.SwagType = "integer"
			prm.SwagFormat = f.Type.Elem().String()
		case reflect.String:
			prm.SwagType = "string"
		case reflect.Bool:
			prm.SwagType = "boolean"
		default:
			//fmt.Println(f.Type)

			//prm.SwagType = "object"
			prm.RefParam = prmKey // name of object, like object (in json
		}

	default:
		fmt.Println("warning: no type")
		fmt.Println(f.Type.Kind())
		// reflect.Int - doesnt supported
	}

	return prm, prmKey
}

// ToSwag converts params to swag format
func ToSwag(val interface{}) map[string]*InParam {
	// apd := Agglo{}

	// // TypeOf returns the reflection Type of the value in the interface{}.
	// // TypeOf(nil) returns nil.
	// t := reflect.TypeOf(&apd)
	// fmt.Println(t) //*mddt.Agglo

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	// Real type (not a pointer) - can be setted directly in previous function
	s := reflect.TypeOf(val) // t.Elem()

	//fmt.Println(s)
	// Field returns a struct type's i'th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	// http://golang.org/pkg/reflect/#StructField
	//var f reflect.StructField = s.Field(5) // 5

	result := map[string]*InParam{}

	for i := 0; i < s.NumField(); i++ {
		prm, prmKey := calcInParam(s.Field(i))

		result[prmKey] = prm
	}

	return result
}
