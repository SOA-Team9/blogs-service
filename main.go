package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"blogs-service.xws.com/handler"
	"blogs-service.xws.com/model"
	"blogs-service.xws.com/repo"
	"blogs-service.xws.com/service"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func startServer(blogHandler *handler.BlogHandler, log *log.Logger) {

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/blogs", blogHandler.MiddlewareContentTypeSet(blogHandler.MiddlewareBlogDeserialization(http.HandlerFunc(blogHandler.Create)))).Methods("POST")
	router.Handle("/blogs", blogHandler.MiddlewareContentTypeSet(http.HandlerFunc(blogHandler.GetBlogs))).Methods("GET")
	router.Handle("/blogs-following/{user_id}", blogHandler.MiddlewareContentTypeSet(http.HandlerFunc(blogHandler.GetBlogsFollowing))).Methods("GET")
	router.Handle("/blogs/{blogId}", blogHandler.MiddlewareContentTypeSet(http.HandlerFunc(blogHandler.GetBlog))).Methods("GET")

	//router.HandleFunc("/blogs", blogHandler.Create).Methods("POST")5
	// router.HandleFunc("/blogs/{blogId}/comments", commentHandler.Create).Methods("POST")
	// router.HandleFunc("/blogs/{blogId}/ratings", ratingHandler.Create).Methods("POST")
	// router.HandleFunc("/blogs/{blogId}/ratings/{userId}", ratingHandler.Delete).Methods("DELETE")

	// router.HandleFunc("/add-equipment/{tourId}", tourHandler.AddEquipment).Methods("POST")
	// router.HandleFunc("/get-tour/{tourId}", tourHandler.GetTourById).Methods("GET")
	// router.HandleFunc("/get-checkpoints/{tourId}", checkpointHandler.GetCheckpoints).Methods("GET")
	// router.HandleFunc("/get-equipment", equipmentHandler.GetEquipment).Methods("GET")

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	println("Server started")
	server := http.Server{
		Addr:         ":" + "8083",
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Server listening on port", ":8083")

	serverError := server.ListenAndServe()

	if serverError != nil {
		log.Fatal(serverError)
	}
}

func main() {
	/*
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
	*/

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[blogs-api] ", log.LstdFlags)
	blogLogger := log.New(os.Stdout, "[blogs-store] ", log.LstdFlags)

	// NoSQL: Initialize Repository stores
	store, err := repo.New(timeoutContext, blogLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.Disconnect(timeoutContext)

	// NoSQL: Checking if the connection was established
	store.Ping()

	blogService := service.NewBlogService(blogLogger, store)
	blogHandler := handler.NewBlogsHandler(blogLogger, blogService)

	startServer(blogHandler, logger)
}
