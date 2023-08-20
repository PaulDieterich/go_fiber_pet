package database

import (
	"gorm.io/gorm"
	"pauldieterich/go_fiber_pet/models"
)

type PetRepository struct {
	DB *gorm.DB
}

func (r *PetRepository) DeletePet(id int) error {
	petModel := models.Pet{}

	res := r.DB.Delete(petModel, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *PetRepository) GetPetById(id int) (*models.Pet, error) {
	petModel := models.Pet{}

	result := r.DB.Where("id: ?", id).Find(petModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &petModel, nil
}

func (r *PetRepository) GetPets() ([]models.Pet, error) {
	petModels := &[]models.Pet{}

	result := r.DB.Find(petModels)

	if result.Error != nil {
		return nil, result.Error
	}
	return *petModels, nil

}

func (r *PetRepository) Save(p *models.Pet) error {

	result := r.DB.Create(&p)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
