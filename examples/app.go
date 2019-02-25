package examples

import (
	"context"
	"github.com/lefelys/apibara/examples/generated/components"
	"github.com/lefelys/apibara/examples/generated/pet"
	"github.com/lefelys/apibara/examples/generated/pet/findbystatus"
)

type s struct{}

func (s *s) AddPetCatOrDog(ctx context.Context, pet func() ([]byte, error)) (resp pet.AddPetResponder) {
	panic("implement me")
}

func (s *s) AddPet(ctx context.Context, pet func() (components.Pet, error)) (resp pet.AddPetResponder) {
	panic("implement me")
}

func (s *s) UpdatePet(ctx context.Context, pet components.Pet) (resp pet.UpdatePetResponder) {
	panic("implement me")
}

func (s *s) findPetsByStatus(ctx context.Context, statuses func() ([]string, error)) (resp petfindbystatus.FindPetsByStatusResponder) {
	panic("implement me")
}
