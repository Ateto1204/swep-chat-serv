package repository

import (
	"encoding/json"
	"time"

	"github.com/Ateto1204/swep-chat-serv/entity"
	"github.com/Ateto1204/swep-chat-serv/internal/domain"
	"gorm.io/gorm"
)

type ChatRepository interface {
	Save(chatID string, name string, membersID []string, t time.Time) (*domain.Chat, error)
	GetByID(id string) (*domain.Chat, error)
	UpdContentsByID(chat *domain.Chat) (*domain.Chat, error)
}

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db}
}

func (r *chatRepository) Save(chatID string, name string, membersID []string, t time.Time) (*domain.Chat, error) {
	chatModel := &domain.Chat{
		ID:       chatID,
		Name:     name,
		Members:  membersID,
		Contents: []string{},
		CreateAt: t,
	}
	chatEntity, err := parseToEntity(chatModel)
	if err != nil {
		return nil, err
	}
	err = r.db.Create(chatEntity).Error
	if err != nil {
		return nil, err
	}
	return chatModel, nil
}

func (r *chatRepository) GetByID(chatID string) (*domain.Chat, error) {
	var chat *entity.Chat
	err := r.db.Where("id = ?", chatID).Order("id").First(&chat).Error
	if err != nil {
		return nil, err
	}
	chatModel, err := parseToModel(chat)
	if err != nil {
		return nil, err
	}
	return chatModel, nil
}

func (r *chatRepository) UpdContentsByID(chat *domain.Chat) (*domain.Chat, error) {
	chatEntity, err := parseToEntity(chat)
	if err != nil {
		return nil, err
	}
	if err := r.db.Model(chatEntity).Update("Contents", chatEntity.Contents).Error; err != nil {
		return nil, err
	}
	return r.GetByID(chat.ID)
}

func parseToEntity(chat *domain.Chat) (*entity.Chat, error) {
	contentsStr, err := strSerialize(chat.Contents)
	if err != nil {
		return nil, err
	}
	membersStr, err := strSerialize(chat.Members)
	if err != nil {
		return nil, err
	}
	chatEntity := &entity.Chat{
		ID:       chat.ID,
		Name:     chat.Name,
		Members:  membersStr,
		Contents: contentsStr,
		CreateAt: chat.CreateAt,
	}
	return chatEntity, nil
}

func parseToModel(chat *entity.Chat) (*domain.Chat, error) {
	contentsData, err := strUnserialize(chat.Contents)
	if err != nil {
		return nil, err
	}
	membersData, err := strUnserialize(chat.Members)
	if err != nil {
		return nil, err
	}
	chatModel := &domain.Chat{
		ID:       chat.ID,
		Name:     chat.Name,
		Members:  membersData,
		Contents: contentsData,
		CreateAt: chat.CreateAt,
	}
	return chatModel, nil
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
