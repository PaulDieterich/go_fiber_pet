package queries

import (
	"pauldieterich/go_fiber_pet/pkg/database"
	"pauldieterich/go_fiber_pet/pkg/models"
)

func DeletePet(id int) error {
	petModel := models.Pet{}

	res := database.Db.Delete(petModel, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetPetById(id int) (*models.Pet, error) {
	petModel := models.Pet{}

	result := database.Db.Where("id=?", id).Find(&petModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return &petModel, nil
}

func GetPets() ([]models.Pet, error) {
	petModels := &[]models.Pet{}

	result := database.Db.Find(&petModels)

	if result.Error != nil {
		return nil, result.Error
	}
	return *petModels, nil

}

func Save(p *models.Pet) error {

	result := database.Db.Create(&p)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
