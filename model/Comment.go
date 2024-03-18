package model

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
    BlogID    uint 
	Context        string    `json:"Context"`
	CreationTime   time.Time `json:"CreationTime"`
	LastUpdateTime time.Time `json:"LastUpdateTime"`
	UserId         int32     `json:"UserId"`
}
