package user

import (
	"context"
	"dataflow/pkg/domain/entity"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/service/user/mock_user"
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

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(ctx, masterTx, uid, name, thumbnail).Return(&entity.User{
		ID:        userID,
		Name:      name,
		Thumbnail: thumbnail,
	}, nil).Times(1)

	interactor := New(masterTxManager, userService)
	insertedUser, err := interactor.CreateNewUser(ctx, uid, name, thumbnail)

	assert.NoError(t, err)
	assert.NotNil(t, insertedUser)
}

func TestIntereractor_GetUserProfile(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUser := &entity.User{
		ID:        userID,
		Name:      name,
		Thumbnail: thumbnail,
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByUID(ctx, masterTx, uid).Return(existedUser, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetUserProfile(ctx, uid)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := entity.UserSlice{
		{
			ID:        userID,
			Name:      name,
			Thumbnail: thumbnail,
		},
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
