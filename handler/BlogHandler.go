package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"blogs-service.xws.com/model"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	logger      *log.Logger
	BlogService *service.BlogService
}

type KeyProduct struct{}

func NewBlogsHandler(l *log.Logger, s *service.BlogService) *BlogHandler {
	return &BlogHandler{l, s}
}

func (handler *BlogHandler) Create(rw http.ResponseWriter, h *http.Request) {
	blog := h.Context().Value(KeyProduct{}).(*model.Blog)
	blog.Status = model.DRAFT
	blog.Rating = 0
	handler.BlogService.Create(blog)
	rw.WriteHeader(http.StatusCreated)
}

func (handler *BlogHandler) GetBlog(rw http.ResponseWriter, h *http.Request) {

	vars := mux.Vars(h)
	id := vars["blogId"]

	patient, err := handler.BlogService.GetBlog(id)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if patient == nil {
		http.Error(rw, "Blog with given id not found", http.StatusNotFound)
		handler.logger.Printf("Blog with id: '%s' not found", id)
		return
	}

	err = patient.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *BlogHandler) GetBlogs(rw http.ResponseWriter, h *http.Request) {

	blogs, err := handler.BlogService.GetBlogs()
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if blogs == nil {
		return
	}

	err = blogs.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (handler *BlogHandler) GetBlogsFollowing(rw http.ResponseWriter, h *http.Request) {

	vars := mux.Vars(h)
	id := vars["user_id"]
	idInt, err := strconv.Atoi(id)
	blogs, err := handler.BlogService.GetBlogsFollowing(idInt)
	if err != nil {
		handler.logger.Print("Database exception: ", err)
	}

	if blogs == nil {
		http.Error(rw, "Blog with given id not found", http.StatusNotFound)
		handler.logger.Printf("Blog with id: '%s' not found", id)
		return
	}

	err = blogs.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		handler.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (m *BlogHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		m.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (f *BlogHandler) MiddlewareBlogDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		blog := &model.Blog{}
		err := blog.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, blog)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
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
