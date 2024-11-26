package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-chat-serv/internal/domain"
	"github.com/Ateto1204/swep-chat-serv/internal/repository"
)

type ChatUseCase interface {
	SaveChat(name string, membersID []string) (*domain.Chat, error)
	GetChat(id string) (*domain.Chat, error)
	AddMsgToChat(msgID, chatID string) (*domain.Chat, error)
	ModifyChatName(chatID, newName string) (*domain.Chat, error)
	AddMemberToChat(chatID, memberID string) (*domain.Chat, error)
	RemoveMembersFromChat(chatID, memberID string) (*domain.Chat, error)
}

type chatUseCase struct {
	repository repository.ChatRepository
}

func NewMsgUseCase(repo repository.ChatRepository) ChatUseCase {
	return &chatUseCase{
		repository: repo,
	}
}

func (uc *chatUseCase) SaveChat(name string, membersID []string) (*domain.Chat, error) {
	if len(membersID) < 1 {
		return nil, errors.New("chat room cannot be less than 2 people")
	}
	t := time.Now()
	chatID := GenerateID()
	if name == "" {
		name = "group"
	}
	chat, err := uc.repository.Save(chatID, name, membersID, t)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (uc *chatUseCase) GetChat(chatID string) (*domain.Chat, error) {
	chat, err := uc.repository.GetByID(chatID)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (uc *chatUseCase) AddMsgToChat(chatID, msgID string) (*domain.Chat, error) {
	chat, err := uc.repository.GetByID(chatID)
	if err != nil {
		return nil, err
	}
	chat.Contents = append(chat.Contents, msgID)
	field := "Contents"
	chat, err = uc.repository.UpdByID(field, chat)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (uc *chatUseCase) AddMemberToChat(chatID, memberID string) (*domain.Chat, error) {
	chat, err := uc.repository.GetByID(chatID)
	if chat == nil || err != nil {
		return nil, err
	}
	chat.Members = append(chat.Members, memberID)
	field := "Members"
	return uc.repository.UpdByID(field, chat)
}

func (uc *chatUseCase) RemoveMembersFromChat(chatID, memberID string) (*domain.Chat, error) {
	chat, err := uc.repository.GetByID(chatID)
	if chat == nil || err != nil {
		return nil, err
	}
	chat.Members = removeFromSlice(chat.Members, memberID)
	field := "Members"
	return uc.repository.UpdByID(field, chat)
}

func (uc *chatUseCase) ModifyChatName(chatID, newName string) (*domain.Chat, error) {
	chat, err := uc.repository.GetByID(chatID)
	if err != nil {
		return nil, err
	}
	chat.Name = newName
	field := "Name"
	return uc.repository.UpdByID(field, chat)
}

func removeFromSlice(slice []string, target string) []string {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func GenerateID() string {
	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(input))
	hashID := hex.EncodeToString(hash.Sum(nil))

	return hashID
}
