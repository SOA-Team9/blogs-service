package repo

import (
	"blogs-service.xws.com/model"
	"gorm.io/gorm"
)

type RatingRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *RatingRepository) AddRating(rating *model.Rating) error {
	dbResult := repo.DatabaseConnection.Create(rating)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *RatingRepository) DeleteRating(userId int) error {

	var rating model.Rating
	if err := repo.DatabaseConnection.Where("user_id = ?", userId).First(&rating).Error; err != nil {
        return err // Return error if rating not found
    }

	if err := repo.DatabaseConnection.Delete(&rating).Error; err != nil {
        return err // Return error if deletion fails
    }

	return nil
}

func (repo *RatingRepository) UpdateRating(newRating *model.Rating) error {
	var rating model.Rating
    if err := repo.DatabaseConnection.Where("user_id = ?", newRating.UserId).First(&rating).Error; err != nil {
        return nil // Return false if rating is not found 
    }

	println("Old type:", rating.RatingType)
	rating.RatingType = newRating.RatingType
	println("New type:", rating.RatingType)

    if err := repo.DatabaseConnection.Save(&rating).Error; err != nil {
        return err // Return error if save operation fails
    }

    return nil // Return nil if update is successful
}


func (repo *RatingRepository) FindByUserId(userId int) bool {
    var rating model.Rating
    if err := repo.DatabaseConnection.Where("user_id = ?", userId).First(&rating).Error; err != nil {
        return false // Return false if rating is not found 
    }
    return true
}