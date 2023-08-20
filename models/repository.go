package models

import "context"

type PetRepository interface {
	DeletePet(ctx context.Context, id int) error
	AllPets(ctx context.Context) ([]Pet, error)
	Save(ctx context.Context) error
	FindPetById(ctx context.Context, id int) (*Pet, error)
}
