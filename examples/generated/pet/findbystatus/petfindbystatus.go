package petfindbystatus

import (
	"encoding/json"
	"github.com/lefelys/apibara/examples/generated/components"
	"net/http"
)

type (
	// Ok
	Ok struct {
		Body []components.Pet
	}
	// Invalid request
	BadRequest components.BadRequest
)

// Add a new pet to the store
type FindPetsByStatusResponder interface {
	FindPetsByStatusResp(w http.ResponseWriter)
}

func (resp *Ok) FindPetsByStatusResp(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Body)
}
func (resp *BadRequest) FindPetsByStatusResp(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}
