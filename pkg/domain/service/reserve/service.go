package reserveservice

import (
	"context"
	"dataflow/pkg/domain/entity/reserve"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/reserve"
	"dataflow/pkg/terrors"
)

type Service interface {
	CreateNewReserve(ctx context.Context, masterTx repository.MasterTx, entity *reserveentity.Reserve) (*reserveentity.Reserve, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*reserveentity.Reserve, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) ([]*reserveentity.Reserve, error)
}

type service struct {
	reserveRepository reserverepository.Repository
}

func New(reserveRepository reserverepository.Repository) Service {
	return &service{
		reserveRepository: reserveRepository,
	}
}

func (s *service) CreateNewReserve(ctx context.Context, masterTx repository.MasterTx, entity *reserveentity.Reserve) (*reserveentity.Reserve, error) {
	reserve, err := s.reserveRepository.InsertReserve(ctx, masterTx, entity)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return reserve, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*reserveentity.Reserve, error) {
	result, err := s.reserveRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return result, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) ([]*reserveentity.Reserve, error) {
	list, err := s.reserveRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return list, nil
}
