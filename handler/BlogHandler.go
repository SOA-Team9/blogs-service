package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (handler *BlogHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var blog model.Blog
	err := json.NewDecoder(request.Body).Decode(&blog)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	blog.Status = model.DRAFT
	blog.Rating = 0
	err = handler.BlogService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) GetBlog(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blogId := vars["blogId"]
	
	fmt.Println(blogId)

	if blogId == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	blogID, err := strconv.Atoi(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(blogID)

	blogs := handler.BlogService.GetBlog(int32(blogID))
	json.NewEncoder(writer).Encode(blogs)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogHandler) GetBlogs(writer http.ResponseWriter, request *http.Request) {
	blog := handler.BlogService.GetBlogs()
	json.NewEncoder(writer).Encode(blog)
	writer.Header().Set("Content-Type", "application/json")
}

// func (handler *BlogHandler) Update(writer http.ResponseWriter, request *http.Request) {
// 	var tour model.Tour
// 	err := json.NewDecoder(request.Body).Decode(&tour)
// 	fmt.Println("Equipment: ", tour.TourEquipment)
// 	if err != nil {
// 		println("error parsing json: ", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	updatedTour, err := handler.BlogService.Update(&tour)
// 	if err != nil {
// 		println("Error while updating the tour")
// 		writer.WriteHeader(http.StatusExpectationFailed)
// 		return
// 	}

// 	updatedTourJSON, err := json.Marshal(updatedTour)
// 	if err != nil {
// 		println("Error encoding updated tour as JSON: ", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusOK)
// 	writer.Header().Set("Content-Type", "application/json")
// 	_, err = writer.Write(updatedTourJSON)
// 	if err != nil {
// 		println("Error writing response: ", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// func (handler *BlogHandler) AddEquipment(writer http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	tourIDStr := vars["tourId"]
// 	tourID, err := strconv.Atoi(tourIDStr)
// 	if err != nil {
// 		fmt.Println("error parsing tour ID:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var equipment []model.Equipment
// 	if err := json.NewDecoder(request.Body).Decode(&equipment); err != nil {
// 		fmt.Println("error parsing JSON:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	if err := handler.BlogService.AddEquipment(int32(tourID), equipment); err != nil {
// 		fmt.Println("error adding equipment to tour:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	writer.WriteHeader(http.StatusCreated)
// }

// func (handler *BlogHandler) GetTourById(writer http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	tourIDstr := vars["tourId"]
// 	tourID, err := strconv.Atoi(tourIDstr)
// 	if err != nil {
// 		fmt.Println("error parsing tour ID:", err)
// 		writer.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	var tour *model.Tour
// 	tour, err = handler.BlogService.GetTourById(int32(tourID))
// 	if err != nil {
// 		fmt.Println("error getting tour by id:", err)
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(writer).Encode(tour)
// 	writer.Header().Set("Content-Type", "application/json")
// }
