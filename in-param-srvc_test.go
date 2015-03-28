package apido

import "fmt"

// methodParams: [{in:form, name:my_param, description:super, format:int16}]
// reqParams {my_param:req_value1, other_param: req_value2}
func ExampleCheckReq() {
	demoMethodParams := []InParam{
		InParam{
			In:          "formData",
			Name:        "demoparam",
			Description: "Demo parameter",
			SwagType:    "integer",
			SwagFormat:  "int16",
			Required:    false,
		},
	}

	demoReqParams := map[string]string{
	//		"demoparam": "1234",
	}

	result, conds := CheckReq(demoMethodParams, demoReqParams)

	fmt.Println(conds)

	fmt.Println(result)

	//Output:
	// map[]
	// map[demoparam:<nil>]
}

//methodParams []InParam, reqParams map[string]string) (map[string]interface{}, map[string]ValidCond) {
