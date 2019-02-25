package generated

import (
	"encoding/json"
	"github.com/lefelys/apibara/examples/generated/components"
	"github.com/lefelys/apibara/examples/generated/errors"
	"net/http"
	"strings"
)

func (s *Service) x() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.addPet)
}

func (s *Service) addPet(w http.ResponseWriter, r *http.Request) { s.O.AddPet(w, r, decodePet) }
func (s *Service) findPetsByStatus(w http.ResponseWriter, r *http.Request) {
	s.O.findPetsByStatus(w, r, decodeStatus)
}

func decodePet(r *http.Request) (pet components.Pet, err error) {
	err = json.NewDecoder(r.Body).Decode(&pet)
	return pet, err
}

func decodeStatus(r *http.Request) (statuses []string, err error) {
	statusesStr := r.URL.Query().Get("status")
	statuses = strings.Split(statusesStr, ",")
	if len(statuses) == 0 {
		statuses = []string{"available"} // default
		return
	}
	// intersection
	var enum = []string{"available", "pending", "sold"}
	contains := func(s string) (contains bool) {
		for _, i := range enum {
			if s == i {
				contains = true
				return
			}
		}
		return
	}
	for _, i := range statuses {
		if !contains(i) {
			return nil, errors.NewPosibleValuesValidationError(i)
		}
	}
	return
}
