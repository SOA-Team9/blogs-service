package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status int

const (
	DRAFT Status = iota
	PUBLISHED
	ACTIVE
	CLOSED
)

type Blog struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name,omitempty" json:"name"`
	Description         string                 `bson:"description,omitempty" json:"description"`
	DateCreated         time.Time              `bson:"creationDate,omitempty" json:"creationDate"`
	Images               string                 `bson:"images,omitempty" json:"images"`
	AuthorId            int32                 `bson:"authorId,omitempty" json:"authorId"`
	Status              Status                  `bson:"status,omitempty" json:"status"`
	Rating            int32                  `bson:"rating,omitempty" json:"rating"`
    Comments []Comment    `bson:"comments,omitempty" json:"comments"`
    Ratings  []Rating     `bson:"ratings,omitempty" json:"ratings"`
}


type Blogs []*Blog

type Comment struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    BlogID    uint `bson:"blogId,omitempty" json:"blogId"`
	Context        string   `bson:"context,omitempty" json:"context"`
	CreationTime   time.Time  `bson:"creationTime,omitempty" json:"creationTime"`
	LastUpdateTime time.Time `bson:"lastUpdateTime,omitempty" json:"lastUpdateTime"`
	UserId         int32     `bson:"userId,omitempty" json:"userId"`
}

type RatingType int

const (
	UPVOTE RatingType = iota
	DOWNVOTE
)

type Rating struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID     uint 	  `bson:"blogId,omitempty" json:"blogId"`
	UserId     int32        `bson:"userId,omitempty" json:"userId"`
	RatingType RatingType    `bson:"ratingType,omitempty" json:"ratingType"`
}


func (p *Blogs) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (b *Blog) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}


func (p *Blog) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Comment) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Comment) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Rating) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Rating) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}