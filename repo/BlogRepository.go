package repo

import (
	"fmt"

	"context"
	"log"
	"time"

	"blogs-service.xws.com/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BlogRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

func NewBlogRepository(client *mongo.Client, logger *log.Logger) *BlogRepository {
	return &BlogRepository{
		cli:    client,
		logger: logger,
	}
}

func New(ctx context.Context, logger *log.Logger) (*BlogRepository, error) {
	dburi := "mongodb://root:pass@blog-db:27017/"

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &BlogRepository{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (pr *BlogRepository) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (pr *BlogRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	// Print available databases
	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (pr *BlogRepository) CreateBlog(blog *model.Blog) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	blogsCollection := pr.getBlogsCollection()

	result, err := blogsCollection.InsertOne(ctx, &blog)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *BlogRepository) getBlogsCollection() *mongo.Collection {
	blogDatabase := pr.cli.Database("blogs")
	blogsCollection := blogDatabase.Collection("blogs")
	return blogsCollection
}

func (pr *BlogRepository) GetBlog(id string) (*model.Blog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blogsCollection := pr.getBlogsCollection()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		pr.logger.Println("Invalid ObjectId:", err)
		return nil, err
	}

	pr.logger.Println("ObjectID:", objectID)

	var blog model.Blog
	err = blogsCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blog)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &blog, nil
}

func (pr *BlogRepository) GetBlogs() (model.Blogs, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogsCollection := pr.getBlogsCollection()

	var blogs model.Blogs
	blogsCursor, err := blogsCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = blogsCursor.All(ctx, &blogs); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlogRating(id int) error {
	// var blog model.Blog
	// repo.DatabaseConnection.Where("ID = ?", id).Preload("Ratings").Find(&blog)
	// println("AAA")
	// blog.Rating = 0

	// for _, rating := range blog.Ratings {
	//     // Update blog based on the rating
	//     // For example, you can calculate an average rating
	//     // and update the blog's rating field
	//     // Here, we simply sum up all rating values
	// 	println(rating.RatingType)
	//     if rating.RatingType == model.DOWNVOTE{
	// 		blog.Rating--
	// 	}else{
	// 		blog.Rating++
	// 	}
	// }

	// if err := repo.DatabaseConnection.Save(&blog).Error; err != nil {
	//     return err // Return error if save operation fails
	// }
	return nil
}

// func (repo *BlogRepository) UpdateBlogRating(id int, number int) error {
// 	var blog model.Blog
// 	repo.DatabaseConnection.Where("ID = ?", id).Find(&blog)
// 	blog.Rating += int32(number)

// 	if err := repo.DatabaseConnection.Save(&blog).Error; err != nil {
//         return err // Return error if save operation fails
//     }
// 	return nil
// }
