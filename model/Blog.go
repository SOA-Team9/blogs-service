package model

import (
	"time"
)

type Status int

const (
	DRAFT Status = iota
	PUBLISHED
	ACTIVE
	CLOSED
)

type Blog struct {
	ID       uint       `gorm:"primaryKey"`
	Name                string                `json:"Name"`
	Description         string                `json:"Description"`
	DateCreated         time.Time             `json:"DateCreated"`
	Images               string               `json:"Images"`
	AuthorId            int32                 `json:"AuthorId"`
	Status              Status                `json:"Status"`
	Rating            int32                 `json:"Rating"`
    Comments []Comment  `gorm:"foreignKey:BlogID"`
    Ratings  []Rating   `gorm:"foreignKey:BlogID"`
}
