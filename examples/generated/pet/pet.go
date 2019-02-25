package pet

import (
	"github.com/lefelys/apibara/examples/generated/codes"
	"github.com/lefelys/apibara/examples/generated/components"
	"net/http"
)

type (
	Ok               struct{ codes.StatusOK }               // Ok
	MethodNotAllowed struct{ codes.StatusMethodNotAllowed } // Invalid input
	NotFound         struct{ codes.StatusNotFound }         // Pet not found
	BadRequest       components.BadRequest                  // Invalid request
)

// Add a new pet to the store
type AddPetResponder interface{ AddPetResp(w http.ResponseWriter) }

func (resp *Ok) AddPetResp(r http.Request, w http.ResponseWriter) { w.WriteHeader(http.StatusOK) }
func (resp *MethodNotAllowed) AddPetResp(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Update an existing pet
type UpdatePetResponder interface{ UpdatePetResp(w http.ResponseWriter) }

func (resp *Ok) UpdatePetResp(w http.ResponseWriter) { w.WriteHeader(http.StatusOK) }
func (resp *MethodNotAllowed) UpdatePetResp(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (resp *NotFound) UpdatePetResp(w http.ResponseWriter)   { w.WriteHeader(http.StatusNotFound) }
func (resp *BadRequest) UpdatePetResp(w http.ResponseWriter) { components.BadRequest.Respond(resp, w) } // ToDo: check if it realy works that way, try (*components.BadRequest) if not
