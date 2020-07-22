package user

import (
	"context"
	"database/sql"
	"dataflow/db/mysql/model"
	"dataflow/pkg/domain/entity"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user"
	"dataflow/pkg/infrastructure/mysql"
	"dataflow/pkg/terrors"
	"log"
	"net/http"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) user.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u *userRepositoryImpliment) InsertUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) (*entity.User, error) {
	newUserData := &model.User{
		UID:       uid,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	if err := newUserData.Insert(ctx, exec, boil.Infer()); err != nil {
		return nil, terrors.Stack(err)
	}

	return ConvertToUserEntity(newUserData), nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	userData, err := model.FindUser(ctx, exec, userID)
	if err == sql.ErrNoRows {
		messageJP := "ユーザが見つかりませんでした。ユーザ登録されているか確認してください。"
		messageEN := "User not found. Please make sure signup."
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToUserEntity(userData), nil
}

func (u *userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err == sql.ErrNoRows {
		messageJP := "不正なユーザです。"
		messageEN := "Invalid user."
		return nil, terrors.Newf(http.StatusUnauthorized, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToUserEntity(userData), nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(ctx, exec)
	if err == sql.ErrNoRows {
		messageJP := "ユーザは1人も登録されていません。"
		messageEN := "User doesn't exists."
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToUserSliceEntity(users), nil
}

func ConvertToUserEntity(userData *model.User) *entity.User {
	return &entity.User{
		ID:        userData.ID,
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail.String,
	}
}

func ConvertToUserSliceEntity(userSlice model.UserSlice) entity.UserSlice {
	res := make(entity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}
