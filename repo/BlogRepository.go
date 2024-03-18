package repo

import (
	"fmt"

	"blogs-service.xws.com/model"
	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) error {
	dbResult := repo.DatabaseConnection.Create(blog)
	if dbResult.Error != nil {
		panic(dbResult.Error)
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *BlogRepository) GetBlog(id int32) []model.Blog {
	var blog []model.Blog
	repo.DatabaseConnection.Where("ID = ?", id).Preload("Comments").Preload("Ratings").Find(&blog)
	fmt.Println(id)
	return blog
}

func (repo *BlogRepository) GetBlogs() []model.Blog {
	var blogs []model.Blog
	repo.DatabaseConnection.Preload("Comments").Preload("Ratings").Find(&blogs)
	return blogs
}

func (repo *BlogRepository) UpdateBlogRating(id int, number int) error {
	var blog model.Blog
	repo.DatabaseConnection.Where("ID = ?", id).Find(&blog)
	blog.Rating += int32(number)

	if err := repo.DatabaseConnection.Save(&blog).Error; err != nil {
        return err // Return error if save operation fails
    }
	return nil
}

// func (repo *BlogRepository) UpdateTour(tour *model.Tour) (*model.Tour, error) {
// 	tx := repo.DatabaseConnection.Begin()

// 	if err := tx.Save(tour).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	var foundTour model.Tour
// 	if err := tx.Preload("TourEquipment").Find(&foundTour).Error; err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	var existingEquipmentIds []int32
// 	for _, equipment := range foundTour.TourEquipment {
// 		existingEquipmentIds = append(existingEquipmentIds, equipment.Id)
// 	}

// 	var newEquipmentIds []int32
// 	for _, equipment := range tour.TourEquipment {
// 		newEquipmentIds = append(newEquipmentIds, equipment.Id)
// 	}

// 	var removedEquipmentIds []int32
// 	for _, existingEquipmentId := range existingEquipmentIds {
// 		found := false
// 		for _, newEquipmentId := range newEquipmentIds {
// 			if existingEquipmentId == newEquipmentId {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			removedEquipmentIds = append(removedEquipmentIds, existingEquipmentId)
// 		}
// 	}

// 	if len(removedEquipmentIds) > 0 {
// 		tx.Where("tour_id = ? AND equipment_id IN (?)", tour.Id, removedEquipmentIds).Delete(&model.TourEquipment{})
// 	}

// 	tx.Commit()

// 	return tour, nil
// }

// func (repo *BlogRepository) AddEquipment(tourId int32, newEquipment []model.Equipment) error {

// 	var tour model.Tour
// 	if err := repo.DatabaseConnection.Where("id = ?", tourId).First(&tour).Error; err != nil {
// 		return err
// 	}
// 	tour.TourEquipment = append(tour.TourEquipment, newEquipment...)

// 	if err := repo.DatabaseConnection.Save(&tour).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (repo *BlogRepository) GetTourById(tourId int32) (*model.Tour, error) {
// 	var tour model.Tour
// 	if err := repo.DatabaseConnection.Where("id = ?", tourId).First(&tour).Error; err != nil {
// 		return nil, err
// 	}
// 	return &tour, nil
// }
