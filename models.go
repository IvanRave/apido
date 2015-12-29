package apido

// ApiSchema - definition of input and output data types
// "type": "array",items": {"$ref": "#/definitions/pet" }, "$ref": "someref"
// http://json-schema.org/example1.html
type ApiSchema struct {
	// Title           string      `json:"title,omitempty"`
	// Description     string      `json:"description,omitempty"`

	// All titles and description - in RefStr
	RefStr string `json:"$ref,omitempty"`
}

// ApiErr defines error type for clients
// Full list of errors generated in code (not in database)
// type ApiErr struct {
//     // till 32000
//     Id int16
//     Message string
//     Description string
// }

type ApiResponse struct {
	Description string    `json:"description"`
	Schema      ApiSchema `json:"schema,omitempty"`
}

type ApiMethod struct {
	Tags []string `json:"tags,omitempty"`
	// this field SHOULD be less than 120 characters.
	Summary     string                 `json:"summary,omitempty"`
	Description string                 `json:"description"`
	OperationId string                 `json:"operationId,omitempty"`
	Consumes    []string               `json:"consumes,omitempty"`
	Produces    []string               `json:"produces,omitempty"`
	Parameters  []InParam              `json:"parameters,omitempty"`
	Responses   map[string]ApiResponse `json:"responses,omitempty"`
	Deprecated  bool                   `json:"deprecated,omitempty"`
	Security    []ScrRequirement       `json:"security,omitempty"`
}

// ApiVerb: { "get": {...}, "post":...}
type ApiVerbs map[string]ApiMethod

// ApiPath: { "/master/getById": { ...}, "/rubric/getNext": {..}}
type ApiPaths map[string]ApiVerbs

// ApiInfo
// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#info-object
type ApiInfo struct {
	Title          string            `json:"title"`
	Description    string            `json:"description,omitempty"`
	TermsOfService string            `json:"termsOfService,omitempty"`
	Contact        map[string]string `json:"contact,omitempty"`
	Version        string            `json:"version"`
	License        map[string]string `json:"license,omitempty"`
}

type ApiDefinition struct {
	// required properties as one array
	Title      string              `json:"title,omitempty"`
	Required   []string            `json:"required,omitempty"`
	Properties map[string]*InParam `json:"properties,omitempty"`
}

// SecurityScheme
// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#securityDefinitionsObject
// Supported schemes are basic authentication, an API key
//   (either as a header or as a query parameter) and OAuth2's common flows
//   (implicit, password, application and access code).
type SecurityScheme struct {
	ScrType string `json:"type"`
	Name    string `json:"name,omitempty"`
	In      string `json:"in,omitempty"`
}

// ScrRequirement - security requirement.
// https://github.com/swagger-api/swagger-spec/blob/master/versions/2.0.md#securityRequirementObject
type ScrRequirement map[string][]string

type ApiSpec struct {
	Swagger  string   `json:"swagger"`
	Host     string   `json:"host"`
	Info     ApiInfo  `json:"info"`
	BasePath string   `json:"basePath"`
	Schemes  []string `json:"schemes"`
	Consumes []string `json:"consumes"`
	Produces []string `json:"produces"`

	Paths               ApiPaths                  `json:"paths""`
	Definitions         map[string]ApiDefinition  `json:"definitions,omitempty"`
	SecurityDefinitions map[string]SecurityScheme `json:"securityDefinitions,omitempty"`
	// A declaration of which security schemes are applied for the API as a whole.
	// The list of values describes alternative security schemes
	//   that can be used (that is, there is a logical OR between the security
	//   requirements). Individual operations can override this definition.
	// "security": [{"petstore_auth": ["write:pets","read:pets"]}]
	Security []ScrRequirement `json:"security,omitempty"`
}

func (tmpApiSpec *ApiSpec) AppendDef(myDefinition string,
	myTitle string,
	myObj interface{}){
	
	tmpApiSpec.Definitions[myDefinition] = ApiDefinition{
		Title:      myTitle,
		Properties: ToSwag(myObj),
	}

	// Definitions: map[string]apido.ApiDefinition{
	// 	"demo_prop": apido.ApiDefinition{
	// 		Title: "Some property",
	// 		//Required: []string{"id", "name"},
	// 		Properties: ToSwag(demoType{}),
	// 	},
}

func (tmpApiSpec *ApiSpec) AppendPath(apiPath string,
	reqType string, // GET, POST only
	summary string,
	description string,
	apiTags []string,
	consumes []string,
	produces []string,
	inParamArr []InParam,
	respMap map[string]ApiResponse){

	// register handler in API docs
	tmpApiSpec.Paths[apiPath] = ApiVerbs{
		reqType: ApiMethod{
			Tags: apiTags,
			Summary: summary,
			Description: description,
			// The id MUST be unique among all operations described in the API.
			// Tools and libraries MAY use the operation id
			//   to uniquely identify an operation.
			// convert my-obj + get-by-ids = myObjGetByIds
			//OperationId: operationId,
			Consumes: consumes,
			Produces: produces,
			// []apido.InParam
			Parameters: inParamArr,
			Responses:  respMap,
			//Deprecated: true,
			// This tag already in global settings
			// Redefine it, if neccessary
			// Security: []mddt.ScrRequirement {
			//     mddt.ScrRequirement{
			//         "api_key": []string{},
			//     },
			// },
		},
	}

}
