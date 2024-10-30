package main

import (
	"log"

	"github.com/Ateto1204/swep-chat-serv/internal/infrastructure"
	"github.com/Ateto1204/swep-chat-serv/internal/repository"
	"github.com/Ateto1204/swep-chat-serv/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := infrastructure.NewDatabase()
	if err != nil {
		panic(err)
	}

	repo := repository.NewChatRepository(db)
	chatUseCase := usecase.NewMsgUseCase(repo)

	router := infrastructure.NewRouter(chatUseCase)
	log.Println("Server Start:")
	router.Run(":8080")
}
