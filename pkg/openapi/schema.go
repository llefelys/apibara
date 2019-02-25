package main

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	"reflect"
)

var z = `{
"openapi" : "sdsd",
"servers" : ["sdsd", "sdsd2"],
"hi" : "sdsd"
}`

func main() {
	var o OpenAPI
	err := json.Unmarshal([]byte(z), &o)
	pp.Println(err)
	pp.Println(o)
	fmt.Println(validate(o))
}

type OpenAPI struct {
	Openapi      string        `json:"openapi" required:"true"`
	Servers      []string      `json:"servers"`
	Tags         []Tags        `json:"tags"`
	Info         *Info         `json:"info" required:"true"`
	Paths        []Path        `json:"paths" required:"true"`
	ExternalDocs *ExternalDocs `json:"externalDocs"`
	Components   *Components   `json:"components"`
	Extensions   `json:"-"`
}

// ToDo 1: wait generics.
// ToDo 2: any more effective ways without double unmarshalling?
// ToDo 3: reflect to get tags?
func (o *OpenAPI) UnmarshalJSON(b []byte) (err error) {
	type _OpenAPI OpenAPI
	_o := _OpenAPI{}
	err = json.Unmarshal(b, &_o)
	*o = OpenAPI(_o)
	if err != nil {
		return err
	}
	if o.Extensions == nil {
		o.Extensions = make(map[string]interface{}, 0)
	}
	err = json.Unmarshal(b, &o.Extensions)
	if err != nil {
		return err
	}
	for _, key := range []string{"openapi", "servers", "tags", "info", "paths", "externalDocs", "components"} {
		delete(o.Extensions, key)
	}
	return
}

type ValidationErr struct {
	fields []string
}

func (e *ValidationErr) Error() string {
	return fmt.Sprintf("missed required fields: %+v", e.fields)
}

func NewValidationErr(fields []string) error {
	return &ValidationErr{
		fields: fields,
	}
}

func validate(i interface{}) (err error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		required := t.Field(i).Tag.Get("required")
		zeroField := reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface())
		if required == "true" && zeroField {
			fields = append(fields, t.Field(i).Name)
		}
	}
	if len(fields) > 0 {
		err = NewValidationErr(fields)
	}
	return
}

type Extensions map[string]interface{}

type Path struct {
	Ref         string    `json:"$ref"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Get         Operation `json:"get"`
	Put         Operation `json:"put"`
	Post        Operation `json:"post"`
	Delete      Operation `json:"delete"`
	Options     Operation `json:"options"`
	Head        Operation `json:"head"`
	Patch       Operation `json:"patch"`
	Trace       Operation `json:"trace"`
	Servers     []Server  `json:"servers"`
	Parameters  []struct {
		*Parameter
		*ReferenceObject
	} `json:"parameters"`
}

type Operation struct {
	Tags         []string     `json:"tags"`
	Summary      string       `json:"summary"`
	Description  string       `json:"description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
	OperationId  string       `json:"operationId"`
	Parameters   struct {
		*Parameter
		*ReferenceObject
	} `json:"parameters"`
	RequestBody struct {
		*RequestBody
		*ReferenceObject
	} `json:"requestBody"`
	Responses map[string]Response `json:"responses"` // required
	Callbacks map[string]struct {
		*Parameter
		*ReferenceObject
	} `json:"callbacks"`
	Deprecated bool     `json:"deprecated"`
	Security   Security `json:"security"`
	Servers    Server   `json:"servers"`
}

type Response struct {
	Description string `json:"description"`
	Headers     map[string]struct {
		*Parameter // read Header Object in docs, not all fields
		*ReferenceObject
	} `json:"headers"`
	Content map[string]MediaType `json:"content"`
	Links   map[string]struct {
		*Link
		*ReferenceObject
	} `json:"links"`
}

type Link struct {
	OperationRef string                 `json:"operationRef"`
	OperationId  string                 `json:"operationId"`
	Parameters   map[string]interface{} `json:"parameters"`
	RequestBody  interface{}            `json:"requestBody"`
	Description  string                 `json:"description"`
	Server       Server                 `json:"server"`
}

type RequestBody struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content"`
	Required    bool                 `json:"required"`
}

type MediaType struct {
	Schema struct {
		*Schema
		*ReferenceObject
	} `json:"schema"`
	Example  interface{} `json:"example"`
	Examples map[string]struct {
		*Example
		*ReferenceObject
	} `json:"examples"`
	Encoding map[string]Encoding `json:"encoding"`
}

type Encoding struct {
	ContentType string `json:"contentType"`
	Headers     map[string]struct {
		*Parameter // read Header Object in docs, not all fields
		*ReferenceObject
	} `json:"headers"`
	Style         string `json:"style"`
	Explode       bool   `json:"explode"`
	AllowReserved bool   `json:"allowReserved"`
}

type Example struct {
	Summary       string      `json:"summary"`
	Description   string      `json:"description"`
	Value         interface{} `json:"value"`
	ExternalValue string      `json:"externalValue"`
}

type ReferenceObject struct {
	Ref string `json:"$ref"`
}

type Parameter struct {
	Name            string `json:"name"` // required
	In              string `json:"in"`   // required
	Description     string `json:"description"`
	Required        bool   `json:"required"`
	Deprecated      bool   `json:"deprecated"`
	AllowEmptyValue bool   `json:"allowEmptyValue"`

	Style         string `json:"style"`
	Explode       bool   `json:"explode"`
	AllowReserved bool   `json:"allowReserved"`
	Schema        *struct {
		*Schema
		*ReferenceObject
	} `json:"schema"`
	Content  *map[string]MediaType `json:"content"` // Schema || Content required
	Example  interface{}           `json:"example"`
	Examples map[string]struct {
		*Example
		*ReferenceObject
	} `json:"examples"`
}

type Schema struct {
	Type                 string
	AllOf                []SchemaOrRef     `json:"allOf"`
	OneOf                []SchemaOrRef     `json:"oneOf"`
	AnyOf                []SchemaOrRef     `json:"anyOf"`
	Not                  []SchemaOrRef     `json:"not"`
	Items                []SchemaOrRef     `json:"items"`
	Properties           map[string]Schema `json:"properties"`
	AdditionalProperties []struct {
		SchemaOrRef
		*bool
	} `json:"additionalProperties"`
	Description string      `json:"description"`
	Format      string      `json:"format"`
	Default     interface{} `json:"default"`

	Title            string        `json:"title"`
	MultipleOf       int64         `json:"multipleOf"`
	Maximum          int64         `json:"maximum"`
	ExclusiveMaximum int64         `json:"exclusiveMaximum"`
	Minimum          int64         `json:"minimum"`
	ExclusiveMinimum int64         `json:"exclusiveMinimum"`
	MaxLength        int64         `json:"maxLength"`
	MinLength        int64         `json:"minLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         uint64        `json:"maxItems"`
	MinItems         uint64        `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	MaxProperties    uint64        `json:"maxProperties"`
	MinProperties    uint64        `json:"minProperties"`
	Required         bool          `json:"required"`
	Enum             []interface{} `json:"enum"`

	Nullable      bool          `json:"nullable"`
	Discriminator Discriminator `json:"discriminator"`
	ReadOnly      bool          `json:"readOnly"`
	WriteOnly     bool          `json:"writeOnly"`
	XML           XML           `json:"xml"`
	ExternalDocs  ExternalDocs  `json:"externalDocs"`
	Example       interface{}   `json:"example"`
	Deprecated    bool          `json:"deprecated"`
}

type SchemaOrRef struct {
	*Schema
	*ReferenceObject
}

type Discriminator struct {
	PropertyName string            `json:"propertyName"` //required
	Mapping      map[string]string `json:"mapping"`
}

type XML struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Prefix    string `json:"prefix"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}

type Info struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	TermsOfService string `json:"termsOfService"`
	Contact        struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"licence"`
	Version string `json:"version"`
}

type Server struct {
	URL         string                    `json:"url"`
	Description string                    `json:"description"`
	Variables   map[string]ServerVariable `json:"variables"`
}

type ServerVariable struct {
	Enum        []string `json:"enum"`
	Default     string   `json:"default"`
	Description string   `json:"description"`
}

type Tags struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
}

type ExternalDocs struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}
type Security struct {
	Type             string     `json:"type"`             // required
	Description      string     `json:"description"`      //
	Name             string     `json:"name"`             // required
	In               string     `json:"in"`               // required
	Scheme           string     `json:"scheme"`           // required
	BearerFormat     string     `json:"bearerFormat"`     //
	Flows            OAuthFlows `json:"flows"`            // required
	OpenIdConnectUrl string     `json:"openIdConnectUrl"` // required
}

type OAuthFlows struct {
	Implicit          OAuthFlow `json:"implicit"`
	Password          OAuthFlow `json:"password"`
	ClientCredentials OAuthFlow `json:"clientCredentials"`
	AuthorizationCode OAuthFlow `json:"authorizationCode"`
}

type OAuthFlow struct {
	AuthorizationUrl string            `json:"authorizationUrl"`
	TokenUrl         string            `json:"tokenUrl"`
	RefreshUrl       string            `json:"refreshUrl"`
	Scopes           map[string]string `json:"scopes"`
}

type Components struct {
	Schemas map[string]struct {
		*Schema
		*ReferenceObject
	} `json:"schemas"`
	Responses map[string]struct {
		*Response
		*ReferenceObject
	} `json:"responses"`
	Parameters map[string]struct {
		*Parameter
		*ReferenceObject
	} `json:"parameters"`
	Examples map[string]struct {
		*Example
		*ReferenceObject
	} `json:"examples"`
	RequestBodies map[string]struct {
		*RequestBody
		*ReferenceObject
	} `json:"requestBodies"`
	Headers map[string]struct {
		*Schema
		*ReferenceObject
	} `json:"headers"`
	SecuritySchemes map[string]struct {
		*Security
		*ReferenceObject
	} `json:"securitySchemes"`
	Links map[string]struct {
		*Link
		*ReferenceObject
	} `json:"links"`
	Callbacks map[string]struct {
		//*Callback
		*ReferenceObject
	} `json:"callbacks"`
}
