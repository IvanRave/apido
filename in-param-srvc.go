package apido

// CheckReq check input params of request (url or body) and params of method
// If params doesnt match - return errors as a map
// methodParams: [{in:form, name:my_param, description:super, format:int16}]
// reqParams {my_param:req_value1, other_param: req_value2}
func CheckReq(methodParams []InParam, reqParams map[string]string) (map[string]interface{}, map[string]ValidCond) {

	validCondParams := make(map[string]ValidCond)

	var fixedParams map[string]interface{} = make(map[string]interface{})

	for _, mp := range methodParams {

		// Send req value to inparam
		// IsMatchValue("superdatafromurl")
		// If reqParams["some"]  doesnt exists - send empty string ""
		outValue, validCond := mp.IsMatchValue(reqParams[mp.Name])
		if validCond.IsValidated() == false {
			//result = false
			validCondParams[mp.Name] = validCond
		} else {
			fixedParams[mp.Name] = outValue
			//fmt.Printf("outvalue: %v %T", outValue, outValue);
			//fmt.Println()
		}
	}

	return fixedParams, validCondParams
}
