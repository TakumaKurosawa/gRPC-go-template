package planrepository

import (
	"context"
	"dataflow/pkg/domain/entity/plan"
	"dataflow/pkg/domain/repository"
)

type Repository interface {
	InsertPlan(ctx context.Context, masterTx repository.MasterTx, plan *planentity.Plan) (*planentity.Plan, error)
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*planentity.Plan, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) ([]*planentity.Plan, error)
}
