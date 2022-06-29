package services

import (
	"github.com/mrinjamul/go-secret/api/controllers"
	"github.com/mrinjamul/go-secret/database"
	"github.com/mrinjamul/go-secret/repository"
)

type Services interface {
	MessageService() controllers.Message
	ViewService() controllers.Views
	HealthCheckService() controllers.HealthCheck
}

type services struct {
	message     controllers.Message
	views       controllers.Views
	healthCheck controllers.HealthCheck
}

func (svc *services) MessageService() controllers.Message {
	return svc.message
}

func (svc *services) ViewService() controllers.Views {
	return svc.views
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// NewServices initializes services
func NewServices() Services {
	db := database.GetDB()
	return &services{
		message: controllers.NewMessage(
			repository.NewMessageRepo(db),
		),
		views: controllers.NewViews(
			repository.NewMessageRepo(db),
		),
		healthCheck: controllers.NewHealthCheck(),
	}
}
