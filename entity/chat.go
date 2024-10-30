package entity

import "time"

type Chat struct {
	ID       string    `gorm:"primary key" json:"id"`
	Members  string    `json:"members"`  // members([]User.ID) of chat room
	Contents string    `json:"contents"` // contents([]Message.ID) of chat room
	CreateAt time.Time `json:"create_at"`
}
