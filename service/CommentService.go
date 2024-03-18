package service

import (
	"fmt"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
)

type CommentService struct {
	CommentRepo *repo.CommentRepository
}

func (service *CommentService) Create(comment *model.Comment) error {
	err := service.CommentRepo.CreateComment(comment)
	if err != nil {
		fmt.Println("Error creating comment: ", err)
		return err
	}
	return nil
}