package domain

import (
	"time"
)

type Chat struct {
	ID        string
	Name      string
	Members   []string // members(User.ID) of chat room
	Contents  []string // contents(Message.ID) of chat room
	CreateAt  time.Time
	IsDeleted bool
}
