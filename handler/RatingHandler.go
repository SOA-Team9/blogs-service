package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
)

type RatingHandler struct {
	RatingService *service.RatingService
	BlogService *service.BlogService
}

func (handler *RatingHandler) Create(writer http.ResponseWriter, request *http.Request) {

	
	vars := mux.Vars(request)
	
	blogId := vars["blogId"]
	
	println(blogId)
	

	//parse rating
	var rating model.Rating
	err := json.NewDecoder(request.Body).Decode(&rating)
	
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//parse blogID and userID
	blogID, err := strconv.Atoi(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	rating.BlogID = uint(blogID)


	if handler.RatingService.FindByUserId(int(rating.UserId)) {
		//update rating
		err = handler.RatingService.Update(&rating)
		if err != nil {
			println("Error while updating a rating")
			writer.WriteHeader(http.StatusExpectationFailed)
			return
		}
		println("Rating updated! ")
		writer.WriteHeader(http.StatusCreated)
		writer.Header().Set("Content-Type", "application/json")
			//update blog rating
		// var number int = 0
		// if rating.RatingType == model.DOWNVOTE{
		// 	number = -2
		// }else{
		// 	number = 2
		// }
		handler.BlogService.UpdteBlogRating(blogID)
		return
	}
	
	//create new rating
	err = handler.RatingService.Create(&rating)
	if err != nil {
		println("Error while adding a new rating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}


	//update blog rating
	// var number int = 0
	// if rating.RatingType == model.DOWNVOTE{
	// 	number = -1
	// }else{
	// 	number = 1
	// }
	handler.BlogService.UpdteBlogRating(blogID)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


func (handler *RatingHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	blogId := vars["blogId"]

	
	userID, err := strconv.Atoi(userId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	blogID, err := strconv.Atoi(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}


	
	//Delete rating
	err = handler.RatingService.Delete(userID)

	if err != nil {
		println("Error while deleting a new rating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	//Update rating on blog
	handler.BlogService.UpdteBlogRating(blogID)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}