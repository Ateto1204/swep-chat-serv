package repository

import (
	"encoding/json"
	"time"

	"github.com/Ateto1204/swep-chat-serv/entity"
	"github.com/Ateto1204/swep-chat-serv/internal/model"
	"gorm.io/gorm"
)

type ChatRepository interface {
	Save(id string, ids []string, t time.Time) (*entity.Chat, error)
	GetByID(id string) (*model.Chat, error)
	UpdByID(id string) error
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db}
}

func (r *chatRepository) Save(id string, ids []string, t time.Time) (*entity.Chat, error) {
	members, err := strSerialize(ids)
	if err != nil {
		return nil, err
	}
	chat := &entity.Chat{
		ID:       id,
		Members:  members,
		Contents: "[]",
		CreateAt: t,
	}
	err = r.db.Create(chat).Error
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *chatRepository) GetByID(id string) (*model.Chat, error) {
	var chat entity.Chat
	err := r.db.Where("id = ?", id).Order("id").First(&chat).Error
	if err != nil {
		return nil, err
	}
	members, err := strUnserialize(chat.Members)
	if err != nil {
		return nil, err
	}
	contents, err := strUnserialize(chat.Contents)
	if err != nil {
		return nil, err
	}
	model := &model.Chat{
		ID:       chat.ID,
		Members:  members,
		Contents: contents,
		CreateAt: chat.CreateAt,
	}
	return model, err
}

func (r *chatRepository) UpdByID(id string) error {
	return nil
}

func strSerialize(sa []string) (string, error) {
	s, err := json.Marshal(sa)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func strUnserialize(s string) ([]string, error) {
	var sa []string
	err := json.Unmarshal([]byte(s), &sa)
	return sa, err
}
