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
		InParam{
			In:          "query",
			Name:        "demobool",
			Description: "Demo boolean",
			SwagType:    "boolean",
			Required:    false,
		},
		InParam{
			In:          "query",
			Name:        "demostring",
			Description: "Demo string value",
			SwagType:    "string",
			Required:    false,
		},
	}

	demoReqParams := map[string]string{
		"demoparam":  "1234",
		"demobool":   "true",
		"demostring": "",
	}

	result, conds := CheckReq(demoMethodParams, demoReqParams)

	fmt.Println(conds)

	fmt.Println(result["demoparam"])
	fmt.Println(result["demobool"])

	//Output:
	// map[]
	// 1234
	// true
}

//methodParams []InParam, reqParams map[string]string) (map[string]interface{}, map[string]ValidCond) {
