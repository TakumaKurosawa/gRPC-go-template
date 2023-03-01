package userservice

import (
	"context"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user"
	"dataflow/pkg/terrors"
)

type Service interface {
	CreateNewUser(ctx context.Context, masterTx repository.MasterTx, entity *userentity.User) (*userentity.User, error)
	GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.User, error)
	GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*userentity.User, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (userentity.UserSlice, error)
}

type service struct {
	userRepository userrepository.Repository
}

func New(userRepository userrepository.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(ctx context.Context, masterTx repository.MasterTx, entity *userentity.User) (*userentity.User, error) {
	insertedUser, err := s.userRepository.InsertUser(ctx, masterTx, entity)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return insertedUser, nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.User, error) {
	userData, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*userentity.User, error) {
	userData, err := s.userRepository.SelectByUID(ctx, masterTx, uid)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (userentity.UserSlice, error) {
	userSlice, err := s.userRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userSlice, nil
}
