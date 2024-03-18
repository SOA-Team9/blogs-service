package model

type RatingType int

const (
	UPVOTE RatingType = iota
	DOWNVOTE
)

type Rating struct {
	ID         uint       `gorm:"primaryKey"`
	BlogID     uint 
	UserId     int32      `json:"UserId"`
	RatingType RatingType `json:"RatingType"`
}
