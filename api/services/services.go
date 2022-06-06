package services

import (
	"github.com/mrinjamul/go-secret/api/controllers"
	"github.com/mrinjamul/go-secret/database"
	"github.com/mrinjamul/go-secret/repository"
)

type Services interface {
	MessageService() controllers.Message
}

type services struct {
	message controllers.Message
}

func (svc *services) MessageService() controllers.Message {
	return svc.message
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		message: controllers.NewMessage(
			repository.NewMessageRepo(db),
		),
	}
}
