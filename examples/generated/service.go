package generated

import (
	"github.com/lefelys/apibara/examples/generated/components"
	"net/http"
)

type Middlewear func(http.Handler) http.Handler

type PetImplementer interface {
	// OneOf returns []bytes
	// Code generation (and open api 3) don't know
	// on what criteria you are desiding the needed type

	// maybe tests??? How can we generate tests to check
	// if choosed type is one of posible by manifest?

	// test for example sends wrong structures to implemented interface
	// and checks method responses. Programmer has to implemennt mocking
	// for inner logick, but for the wrong request data always must be
	// returner error

	// Same with array length. Generate tests. Or make []interface{},
	// decode in it and check explisictly? On request ok, but on response - tests.
	// Same uniqueness
	//
	// Enums - make type and constants to it.

	// Fixed Keys in dict - struct with fixed keys and map in it
	// Discriminator HELPS A LOT to check type. Decode it in structure, get
	// discriminator type name, and switch then to decode in proper type

	//
	// marks 02.02.2019
	// 1) when it's posible: genereated code validates everything in request and gives you
	// result and error (predefined type for every parameter validation to make type switch checks)
	// 2) when it's not possible to deside what type request-body must be unmarshalled to:
	// generated code gives you []byte and content-type header (annd other parameters)
	// 3) if it is posible to unnmarshall json/xml - generted code will do that, if there is
	// extra possible type text/plain or others that are not unmarshallable - it gives you
	// []byte and content-type
	// 4) content-types like "image/*" gives you []byte
	// 5) responses are not strict - all methods gives ResponseWriter, and app can answer
	// with any combination of http code, body, headers. Can we make sure that response follows
	// the manifest? I think no. Quote:
	//  > Note that an API specification does not necessarily need
	// 	> to cover all possible HTTP response codes, since they
	//  > may not be known in advance. However, it is expected
	//  > to cover successful responses and any known errors.
	//

	// Add a new pet to the store'
	AddPetCatOrDog(w http.ResponseWriter, r *http.Request,
		pet func(r *http.Request) ([]byte, error),
		headerParams func(r *http.Request) (string, error))

	AddPet(w http.ResponseWriter, r *http.Request,
		pet func(r *http.Request) (components.Pet, error))

	// Update an existing pet
	UpdatePet(w http.ResponseWriter, r *http.Request,
		pet func(r *http.Request) (components.Pet, error))
}

type PetFindByStatusImplementer interface {
	// Finds Pets by status
	findPetsByStatus(w http.ResponseWriter, r *http.Request,
		statuses func(r *http.Request) ([]string, error))
}

type BearerAuth interface {
	CheckBearer(bearer string) bool
}

type Auth interface {
	BearerAuth
}

type Operations interface {
	PetImplementer
	PetFindByStatusImplementer
}

type Service struct {
	O      Operations   //  implement these interfaces
	A      Auth         //  and call init() to initialise
	Server *http.Server //  all handlers
}
