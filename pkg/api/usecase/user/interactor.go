package user

import (
	"context"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
	userservice "dataflow/pkg/domain/service/user"
	"dataflow/pkg/terrors"
)

type Interactor interface {
	CreateNewUser(ctx context.Context, uid, name, thumbnail string) (*userentity.User, error)
	GetUserProfile(ctx context.Context, uid string) (*userentity.User, error)
	GetAll(ctx context.Context) (userentity.UserSlice, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	userService     userservice.Service
}

func New(masterTxManager repository.MasterTxManager, userService userservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		userService:     userService,
	}
}

func (i *intereractor) CreateNewUser(ctx context.Context, uid, name, thumbnail string) (*userentity.User, error) {
	var userData *userentity.User
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// 新規ユーザ作成
		userData, err = i.userService.CreateNewUser(ctx, masterTx, &userentity.User{
			UID:       uid,
			Name:      name,
			Thumbnail: thumbnail,
		})
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (i *intereractor) GetUserProfile(ctx context.Context, uid string) (*userentity.User, error) {
	var userData *userentity.User
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのプロフィール情報取得
		userData, err = i.userService.GetByUID(ctx, masterTx, uid)
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userData, nil
}

func (i *intereractor) GetAll(ctx context.Context) (userentity.UserSlice, error) {
	var userSlice userentity.UserSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// (管理者用)ユーザ全件取得
		userSlice, err = i.userService.GetAll(ctx, masterTx)
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return userSlice, nil
}
