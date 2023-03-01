package reserverepository

import (
	"context"
	"dataflow/pkg/domain/entity/reserve"
	"dataflow/pkg/domain/repository"
)

type Repository interface {
	InsertReserve(ctx context.Context, masterTx repository.MasterTx, reserve *reserveentity.Reserve) (*reserveentity.Reserve, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*reserveentity.Reserve, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) ([]*reserveentity.Reserve, error)
}
