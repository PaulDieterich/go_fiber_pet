package queries

import (
	"pauldieterich/go_fiber_pet/pkg/database"
	"pauldieterich/go_fiber_pet/pkg/models"
)

func GetAllUsers() ([]models.User, error) {
	userModels := []models.User{}

	result := database.Db.Find(&userModels)

	if result.Error != nil {
		return nil, result.Error
	}
	return userModels, nil
}
func GetUserById(id int) (*models.User, error) {
	user := models.User{}

	result := database.Db.Where("id=?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func FilterUserByRole(roleName models.Role) ([]models.User, error) {
	usermodels := []models.User{}

	result := database.Db.Where("role=?", roleName).Find(usermodels)

	if result.Error != nil {
		return nil, result.Error
	}
	return usermodels, nil

}

func DeleteUser(id int) error {
	user := models.User{}

	result := database.Db.Delete(user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func SaveUser(user *models.User) error {

	//TODO: implement
	return nil
}
