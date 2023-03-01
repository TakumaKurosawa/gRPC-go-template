package userinfra

import (
	"context"
	"database/sql"
	"dataflow/db/mysql/model"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user"
	"dataflow/pkg/infrastructure/mysql"
	"dataflow/pkg/terrors"
	"log"

	"github.com/VividCortex/mysqlerr"
	driver "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"google.golang.org/grpc/codes"
)

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) userrepository.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u *userRepositoryImpliment) InsertUser(ctx context.Context, masterTx repository.MasterTx, user *userentity.User) (*userentity.User, error) {
	newUserData := convertToDto(user)
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Wrapf(err, codes.Internal, "サーバーでエラーが起きました", "server error occurred.")
	}

	if err := newUserData.Insert(ctx, exec, boil.Infer()); err != nil {
		if driverErr, ok := err.(*driver.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
				return nil, terrors.Wrapf(err, codes.Internal, "すでに存在しているユーザです。", "User is already exists.")
			}
		}
		return nil, terrors.Wrapf(err, codes.Internal, "サーバーでエラーが起きました", "server error occurred.")
	}

	return convertToUserEntity(newUserData), nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	dto, err := model.FindUser(ctx, exec, userID)
	if err == sql.ErrNoRows {
		messageJP := "ユーザが見つかりませんでした。ユーザ登録されているか確認してください。"
		messageEN := "User not found. Please make sure signup."
		return nil, terrors.Newf(codes.Internal, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, codes.Internal, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return convertToUserEntity(dto), nil
}

func (u *userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*userentity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	dto, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err == sql.ErrNoRows {
		messageJP := "不正なユーザです。"
		messageEN := "Invalid user."
		return nil, terrors.Newf(codes.Unauthenticated, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Stack(err)
	}

	return convertToUserEntity(dto), nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (userentity.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	var queries []qm.QueryMod
	users, err := model.Users(queries...).All(ctx, exec)
	if err == sql.ErrNoRows {
		messageJP := "ユーザは1人も登録されていません。"
		messageEN := "User doesn't exists."
		return nil, terrors.Newf(codes.Internal, messageJP, messageEN)
	}
	if err != nil {
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, codes.Internal, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return convertToUserSliceEntity(users), nil
}
