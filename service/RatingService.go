package service

import (
	"fmt"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
)

type RatingService struct {
	RatingRepo *repo.RatingRepository
}

func (service *RatingService) Create(rating *model.Rating) error {
	err := service.RatingRepo.AddRating(rating)
	if err != nil {
		fmt.Println("Error creating rating: ", err)
		return err
	}
	return nil
}

func (service *RatingService) Delete(ratingId int) error {
	err := service.RatingRepo.DeleteRating(ratingId)
	if err != nil {
		fmt.Println("Error creating rating: ", err)
		return err
	}
	return nil
}

func (service *RatingService) Update(rating *model.Rating) error {
	err := service.RatingRepo.UpdateRating(rating)
	if err != nil {
		fmt.Println("Error creating rating: ", err)
		return err
	}
	return nil
}

func (service *RatingService) FindByUserId(userId int) bool {
	return service.RatingRepo.FindByUserId(userId)
}