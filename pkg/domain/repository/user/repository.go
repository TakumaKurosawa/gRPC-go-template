package userrepository

import (
	"context"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
)

type Repository interface {
	InsertUser(ctx context.Context, masterTx repository.MasterTx, user *userentity.User) (*userentity.User, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.User, error)
	SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*userentity.User, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (userentity.UserSlice, error)
}
