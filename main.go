package main

import (
	"fmt"
	"log"
	"net/http"

	"blogs-service.xws.com/handler"
	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	
	connectionStr := "user=postgres password=super dbname=blogs host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil
	}

	db.AutoMigrate(&model.Comment{})
	db.AutoMigrate(&model.Rating{})
	db.AutoMigrate(&model.Blog{})
	fmt.Println("Successfully connected!")

	return db
}



func startServer(blogHandler *handler.BlogHandler,	commentHandler *handler.CommentHandler, ratingHandler *handler.RatingHandler){

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", blogHandler.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", blogHandler.GetBlog).Methods("GET")
	router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")
	router.HandleFunc("/blogs/{blogId}/comments", commentHandler.Create).Methods("POST")
	router.HandleFunc("/blogs/{blogId}/ratings", ratingHandler.Create).Methods("POST")
	router.HandleFunc("/blogs/{blogId}/ratings/{userId}", ratingHandler.Delete).Methods("DELETE")

	// router.HandleFunc("/add-equipment/{tourId}", tourHandler.AddEquipment).Methods("POST")
	// router.HandleFunc("/get-tour/{tourId}", tourHandler.GetTourById).Methods("GET")
	// router.HandleFunc("/get-checkpoints/{tourId}", checkpointHandler.GetCheckpoints).Methods("GET")
	// router.HandleFunc("/get-equipment", equipmentHandler.GetEquipment).Methods("GET")

	println("Server started")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	commentRepo := &repo.CommentRepository{DatabaseConnection: database}
	commentService := &service.CommentService{CommentRepo: commentRepo}
	commentHandler := &handler.CommentHandler{CommentService: commentService}

	ratingRepo := &repo.RatingRepository{DatabaseConnection: database}
	ratingService := &service.RatingService{RatingRepo: ratingRepo}
	ratingHandler := &handler.RatingHandler{RatingService: ratingService, BlogService: blogService}

	startServer(blogHandler, commentHandler, ratingHandler)
}
