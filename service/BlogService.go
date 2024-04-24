package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
)

type BlogService struct {
	logger   *log.Logger
	BlogRepo *repo.BlogRepository
}

func NewBlogService(l *log.Logger, r *repo.BlogRepository) *BlogService {
	return &BlogService{l, r}
}

func (service *BlogService) Create(blog *model.Blog) error {
	err := service.BlogRepo.CreateBlog(blog)
	if err != nil {
		fmt.Println("Error creating blog: ", err)
		return err
	}
	return nil
}

func (service *BlogService) GetBlog(id string) (*model.Blog, error) {
	blog, err := service.BlogRepo.GetBlog(id)
	if err != nil {
		service.logger.Print("Database exception: ", err)
	}
	return blog, err
}

func (service *BlogService) GetBlogs() (model.Blogs, error) {
	blogs, err := service.BlogRepo.GetBlogs()
	if err != nil {
		service.logger.Print("Database exception: ", err)
	}
	return blogs, err
}

func (service *BlogService) UpdteBlogRating(id int) error {
	blogs := service.BlogRepo.UpdateBlogRating(id)
	return blogs
}

func (service *BlogService) GetBlogsFollowing(id int) (model.Blogs, error) {
	url := fmt.Sprintf("http://followers:8086/user/following-ids/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		service.logger.Fatal("Failed to get recommendation from the microservice "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	defer resp.Body.Close()

	var userIds []int
	err = json.NewDecoder(resp.Body).Decode(&userIds)

	if err != nil {
		service.logger.Fatal("Failed to decode users IDs", http.StatusInternalServerError)
		return nil, err
	}

	log.Println("Recommendation IDs: ", userIds)

	blogs, err := service.BlogRepo.GetBlogs()
	if err != nil {
		return nil, err
	}

	var filteredBlogs model.Blogs
	for _, blog := range blogs {
		for _, userId := range userIds {
			if blog.AuthorId == int32(userId) {
				filteredBlogs = append(filteredBlogs, blog)
				break
			}
		}
	}

	return filteredBlogs, nil
}

// func (service *BlogService) UpdteBlogRating(id int, number int) error {
// 	blogs := service.BlogRepo.UpdateBlogRating(id, number)
// 	return blogs
// }

// func (service *BlogService) Update(tourToUpdate *model.Tour) (*model.Tour, error) {
// 	updatedTour, err := service.BlogRepo.UpdateTour(tourToUpdate)
// 	if err != nil {
// 		fmt.Println("Error updating tour: ", err)
// 		return nil, err
// 	}
// 	return updatedTour, nil
// }

// func (service *BlogService) AddEquipment(tourId int32, newEquipment []model.Equipment) error {
// 	err := service.BlogRepo.AddEquipment(tourId, newEquipment)
// 	if err != nil {
// 		fmt.Println("Error adding equipment to tour: ", err)
// 		return err
// 	}
// 	return nil
// }

// func (service *BlogService) GetTourById(tourId int32) (*model.Tour, error) {
// 	tour, err := service.BlogRepo.GetTourById(tourId)
// 	if err != nil {
// 		fmt.Println("Error getting tour by id: ", err)
// 		return nil, err
// 	}
// 	return tour, nil
// }
