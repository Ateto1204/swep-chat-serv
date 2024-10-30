package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-chat-serv/entity"
	"github.com/Ateto1204/swep-chat-serv/internal/model"
	"github.com/Ateto1204/swep-chat-serv/internal/repository"
)

type ChatUseCase interface {
	SaveChat(ids []string) (*entity.Chat, error)
	GetChat(id string) (*model.Chat, error)
	GenerateID() string
}

type chatUseCase struct {
	repository repository.ChatRepository
}

func NewMsgUseCase(repo repository.ChatRepository) ChatUseCase {
	return &chatUseCase{
		repository: repo,
	}
}

func (uc *chatUseCase) SaveChat(ids []string) (*entity.Chat, error) {
	t := time.Now()
	id := uc.GenerateID()
	chat, err := uc.repository.Save(id, ids, t)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (uc *chatUseCase) GetChat(id string) (*model.Chat, error) {
	chat, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (uc *chatUseCase) GenerateID() string {
	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(input))
	hashID := hex.EncodeToString(hash.Sum(nil))

	return hashID
}
