package userinfra

import (
	"context"
	"dataflow/pkg/domain/entity/user"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/user"
	"dataflow/pkg/infrastructure/dynamo"
	"dataflow/pkg/terrors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

const tableName = "users"

func (u *userRepositoryImpliment) InsertUser(ctx context.Context, masterTx repository.MasterTx, entity *userentity.User) (*userentity.User, error) {
	dto := convertToDto(entity)
	exec, err := dynamo.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Wrapf(err, codes.Internal, "サーバーでエラーが起きました", "server error occurred.")
	}

	results, err := dynamodbattribute.MarshalMap(dto)
	if _, err := exec.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      results,
	}); err != nil {
		return nil, terrors.Wrapf(err, codes.Internal, "サーバーでエラーが起きました", "server error occurred.")
	}

	return convertToUserEntity(dto), nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*userentity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*userentity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (userentity.UserSlice, error) {
	// TODO implement me
	panic("implement me")
}
