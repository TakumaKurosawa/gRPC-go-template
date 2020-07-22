package user

import (
	"context"
	"dataflow/pkg/domain/entity"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user"
	"dataflow/pkg/terrors"
)

type Service interface {
	CreateNewUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) (*entity.User, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) (*entity.User, error) {
	insertedUser, err := s.userRepository.InsertUser(ctx, masterTx, uid, name, thumbnail)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return insertedUser, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	userData, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	userData, err := s.userRepository.SelectByUID(ctx, masterTx, uid)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	userSlice, err := s.userRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userSlice, nil
}
