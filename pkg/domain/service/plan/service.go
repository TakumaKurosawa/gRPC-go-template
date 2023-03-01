package planservice

import (
	"context"
	"dataflow/pkg/domain/entity/plan"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/plan"
	"dataflow/pkg/terrors"
)

type Service interface {
	CreateNewPlan(ctx context.Context, masterTx repository.MasterTx, entity *planentity.Plan) (*planentity.Plan, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*planentity.Plan, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) ([]*planentity.Plan, error)
}

type service struct {
	planRepository planrepository.Repository
}

func New(planRepository planrepository.Repository) Service {
	return &service{
		planRepository: planRepository,
	}
}

func (s *service) CreateNewPlan(ctx context.Context, masterTx repository.MasterTx, entity *planentity.Plan) (*planentity.Plan, error) {
	// 予約が1件も登録なしの場合はステータスを非公開にする
	if len(entity.Reserves) == 0 {
		entity.Status = planentity.PlanStatusClose
	}

	plan, err := s.planRepository.InsertPlan(ctx, masterTx, entity)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return plan, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*planentity.Plan, error) {
	result, err := s.planRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return result, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) ([]*planentity.Plan, error) {
	list, err := s.planRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return list, nil
}
