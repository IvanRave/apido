package apido

import (
	"strconv"
)

// InParam describes a single operation parameter
// http://json-schema.org/latest/json-schema-validation.html
// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#parameter-object
type InParam struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	// The value "type" MUST be one of
	// "string", "number", "integer", "boolean", "array" or "file".
	// If type is "file", the consumes MUST be either "multipart/form-data"
	// or " application/x-www-form-urlencoded" and the parameter MUST be in "formData".
	SwagType string   `json:"type,omitempty"`
	ArrItem  *InParam `json:"items,omitempty"`
	RefParam string   `json:"$ref,omitempty"`
	// Props and AddtProps doesnt work in UI
	//Props  map[string]InParam   `json:"properties,omitempty"`
	//AddtProps  map[string]string   `json:"additionalProperties,omitempty"`
	// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#dataTypeFormat
	// int32, int64, float, double, byte
	SwagFormat       string `json:"format,omitempty"`
	Maximum          int32  `json:"maximum,omitempty"`
	Minimum          int32  `json:"minimum,omitempty"`
	ExclusiveMaximum bool   `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum bool   `json:"exclusiveMinimum,omitempty"`
	MinLength        int32  `json:"minLength,omitempty"`
	MaxLength        int32  `json:"maxLength,omitempty"`
	Pattern          string `json:"pattern,omitempty"`
}

// calcIntegerFormat returns 8, 16, 32, 64
func calcIntegerFormat(swagFormat string) int8 {
	// by default: 32
	switch swagFormat {
	case "int8":
		return 8
	case "int16":
		return 16
	case "int32":
		return 32
	case "int64":
		return 64
	default:
		return 32
	}
}

// transormInteger to required format
func transformInteger(val int64, integerFormat int8) interface{} {
	switch integerFormat {
	case 8:
		return int8(val)
	case 16:
		return int16(val)
	case 32:
		return int32(val)
	case 64:
		return int64(val)
	default:
		return int32(val)
	}
}

// IsMatchValue checks a value from a request (from url or body)
// Return unmatched properties, like
// { unmatched: {
// valType: { our: string, yours: integer },
// maxLength: { our: 10, yours: 20 },
// minLength: { our: 2, yours: 0 },
// required: { our: true, yours: false }}
// Rerun ourValue, converted to required SwagType
// If val is not exists - empty string ""
// You can not use outValue if ValidConditions is not empty
func (inp *InParam) IsMatchValue(val string, isValExists bool) (interface{}, ValidCond) {

	validCond := ValidCond{
		UnMatched: map[string]string{},
	}

	if inp.Required == true && isValExists == false {
		validCond.UnMatched["paramRequired"] = "true"
		return nil, validCond
	}

	if inp.Required == false && isValExists == false {
		return nil, validCond
	}

	if inp.SwagType == "string" {
		// no conversion
		return val, validCond
	}

	if inp.SwagType == "boolean" {
		outBool, errOutBool := strconv.ParseBool(val)
		if errOutBool != nil {
			validCond.UnMatched["paramType"] = "boolean"
			return nil, validCond
		}

		return outBool, validCond
	}

	// if number exists: check it and transform from string to integer
	if inp.SwagType == "integer" {
		// int8, 16, 32, 64 or nil
		integerFormat := calcIntegerFormat(inp.SwagFormat)
		//fmt.Printf("integerFormat: %v", integerFormat)
		i, e := strconv.ParseInt(val, 10, int(integerFormat))
		if e != nil {
			validCond.UnMatched["paramType"] = "integer"
			return nil, validCond
		} else {
			transValue := transformInteger(i, integerFormat)
			return transValue, validCond
		}
	}

	// if no cases: return unknownError
	validCond.UnMatched["unknownType"] = ""
	return nil, validCond
}
