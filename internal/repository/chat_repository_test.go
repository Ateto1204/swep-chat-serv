package repository_test

import (
	"testing"
	"time"

	"github.com/Ateto1204/swep-chat-serv/entity"
	"github.com/Ateto1204/swep-chat-serv/internal/domain"
	"github.com/Ateto1204/swep-chat-serv/internal/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	db.AutoMigrate(&entity.Chat{})
	testDB = db
}

func TestSave(t *testing.T) {
	setupTestDB()
	repo := repository.NewChatRepository(testDB)

	chatID := "group123"
	name := "apt"
	membersID := []string{"user1", "user2", "user3"}
	contents := []string{}
	now := time.Now()

	chat, err := repo.Save(chatID, name, membersID, now)
	assert.NoError(t, err)
	assert.Equal(t, chatID, chat.ID)
	assert.Equal(t, name, chat.Name)
	assert.Equal(t, membersID, chat.Members)
	assert.Equal(t, contents, chat.Contents)
	assert.Equal(t, now, chat.CreateAt)
}

func TestGetByID(t *testing.T) {
	setupTestDB()
	repo := repository.NewChatRepository(testDB)

	chatID := "group123"
	name := "apt"
	membersID := []string{"user1", "user2", "user3"}
	contents := []string{}
	now := time.Now()
	repo.Save(chatID, name, membersID, now)

	chat, err := repo.GetByID(chatID)
	assert.NoError(t, err)
	assert.Equal(t, chatID, chat.ID)
	assert.Equal(t, name, chat.Name)
	assert.Equal(t, membersID, chat.Members)
	assert.Equal(t, contents, chat.Contents)

	assert.True(t, chat.CreateAt.Equal(now), "CreateAt should match")
}

func TestUpdByID(t *testing.T) {
	setupTestDB()
	repo := repository.NewChatRepository(testDB)

	chatID := "group123"
	name := "apt"
	membersID := []string{"user1", "user2", "user3"}
	now := time.Now()
	chat, _ := repo.Save(chatID, name, membersID, now)

	chatModel := &domain.Chat{
		ID:       chat.ID,
		Name:     chat.Name,
		Members:  chat.Members,
		Contents: []string{"demo_msg_id"},
		CreateAt: chat.CreateAt,
	}

	field := "Contents"
	updatedChat, err := repo.UpdByID(field, chatModel)

	assert.NoError(t, err)
	assert.NotNil(t, updatedChat)
	assert.Equal(t, chatID, updatedChat.ID)
	assert.Equal(t, name, updatedChat.Name)
	assert.Equal(t, membersID, updatedChat.Members)
	assert.Equal(t, []string{"demo_msg_id"}, updatedChat.Contents)
}
