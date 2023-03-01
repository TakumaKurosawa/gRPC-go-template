package userservice

import (
	"context"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user/mock_user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID    = 1
	uid       = "uid"
	name      = "name"
	thumbnail = "thumbnail"
)

func TestService_CreateNewUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(ctx, masterTx, uid, name, thumbnail).Return(&user.User{
		ID:        userID,
		Name:      name,
		Thumbnail: thumbnail,
	}, nil).Times(1)

	service := New(userRepository)
	insertedUser, err := service.CreateNewUser(ctx, masterTx, uid, name, thumbnail)

	assert.NoError(t, err)
	assert.NotNil(t, insertedUser)
}

func TestService_GetByPK(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUser := &user.User{
		ID:        userID,
		Name:      name,
		Thumbnail: thumbnail,
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByPK(ctx, masterTx, userID).Return(existedUser, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetByPK(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestService_SelectAll(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUsers := user.UserSlice{
		{
			ID:        userID,
			Name:      name,
			Thumbnail: thumbnail,
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
