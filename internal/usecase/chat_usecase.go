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
	if len(membersID) < 3 {
		return nil, errors.New("chat room cannot be less than 3 people")
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

func GenerateID() string {
	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(input))
	hashID := hex.EncodeToString(hash.Sum(nil))

	return hashID
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

func (uc *chatUseCase) AddMemberToChat(memberID []string, chatID string) error {
	return nil
}

func (uc *chatUseCase) RemoveMembersFromChat(memberID []string, chatID string) error {
	return nil
}

func (uc *chatUseCase) ModifyChatName(newName string) error {
	return nil
}
