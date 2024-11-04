package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-chat-serv/entity"
	"github.com/Ateto1204/swep-chat-serv/internal/domain"
	"github.com/Ateto1204/swep-chat-serv/internal/repository"
)

type ChatUseCase interface {
	SaveChat(name string, membersID []string) (*entity.Chat, error)
	GetChat(id string) (*domain.Chat, error)
}

type chatUseCase struct {
	repository repository.ChatRepository
}

func NewMsgUseCase(repo repository.ChatRepository) ChatUseCase {
	return &chatUseCase{
		repository: repo,
	}
}

func (uc *chatUseCase) SaveChat(name string, membersID []string) (*entity.Chat, error) {
	if len(membersID) < 3 {
		return nil, errors.New("chat room cannot be less than 3 people")
	}
	t := time.Now()
	chatID := GenerateID()
	if name == "" {
		name = "gruop"
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

func AddMsgToChat(msgId, chatID string) error {
	return nil
}

func AddMemberToChat(memberID []string, chatID string) error {
	return nil
}

func RemoveMembersFromChat(memberID []string, chatID string) error {
	return nil
}
