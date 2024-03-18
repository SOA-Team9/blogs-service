package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) Create(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blogId := vars["blogId"]

	var comment model.Comment
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	blogID, err := strconv.Atoi(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
	comment.BlogID = uint(blogID)
	err = handler.CommentService.Create(&comment)
	if err != nil {
		println("Error while adding a new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


func (handler *CommentHandler) Update(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	commentId := vars["commentId"]

	var comment model.Comment
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		println("error parsing json: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	commentID, err := strconv.Atoi(commentId)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
	comment.BlogID = uint(blogID)
	err = handler.CommentService.Create(&comment)
	if err != nil {
		println("Error while adding a new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}