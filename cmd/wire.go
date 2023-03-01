//go:build wireinject
// +build wireinject

package main

import (
	userHandler "dataflow/pkg/api/handler/user"
	userInteractor "dataflow/pkg/api/usecase/user"
	"dataflow/pkg/domain/repository"
	userSvc "dataflow/pkg/domain/service/user"
	userRepo "dataflow/pkg/infrastructure/dynamo/user"

	"github.com/google/wire"
)

func InitUserAPI(masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
