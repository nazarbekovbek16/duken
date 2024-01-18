package handlers

import (
	"archi/config"
	"archi/service"
	"archi/storage"
	"go.uber.org/zap"
)

type Handlers struct {
	User *UserHandler
	log  *zap.Logger
}

func NewHandlers(l *zap.Logger, config *config.Config, storage *storage.Storage, service *service.Service) *Handlers {
	return &Handlers{
		User: NewUserHandler(l, config, storage, service),
		log:  l,
	}
}
