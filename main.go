package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"blogs-service.xws.com/handler"
	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
	"blogs-service.xws.com/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *gorm.DB {
	
	connectionStr := "user=postgres password=super dbname=blogs host=blog-db port=5432 sslmode=disable"
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

func initMongoDB() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Error connecting to database: ", err)
        return nil
    }
    fmt.Println("Successfully connected to MongoDB!")
    return client
}



func startServer(blogHandler *handler.BlogHandler){

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blogs", blogHandler.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{blogId}", blogHandler.GetBlog).Methods("GET")
	router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")
	// router.HandleFunc("/blogs/{blogId}/comments", commentHandler.Create).Methods("POST")
	// router.HandleFunc("/blogs/{blogId}/ratings", ratingHandler.Create).Methods("POST")
	// router.HandleFunc("/blogs/{blogId}/ratings/{userId}", ratingHandler.Delete).Methods("DELETE")

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

	mongoDb := initMongoDB()
	if mongoDb == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	logger := log.New(os.Stdout, "[blog-repo] ", log.LstdFlags)

	blogRepo := repo.NewBlogRepository(mongoDb, logger)
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)
}
