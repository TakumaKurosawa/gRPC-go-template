package reserveinfra

import (
	"context"
	"dataflow/pkg/domain/entity/reserve"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/reserve"
	"dataflow/pkg/infrastructure/dynamo"
	"dataflow/pkg/terrors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"google.golang.org/grpc/codes"
)

type reserveRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) reserverepository.Repository {
	return &reserveRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

const tableName = "reserves"

func (u *reserveRepositoryImpliment) InsertReserve(ctx context.Context, masterTx repository.MasterTx, entity *reserveentity.Reserve) (*reserveentity.Reserve, error) {
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
	reserve, err := convertToReserveEntity(dto)
	if err != nil {
		return nil, terrors.Wrapf(err, codes.Internal, "サーバーでエラーが起きました", "server error occurred.")
	}

	return reserve, nil
}

func (u *reserveRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*reserveentity.Reserve, error) {
	// TODO implement me
	panic("implement me")
}

func (u *reserveRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) ([]*reserveentity.Reserve, error) {
	// TODO implement me
	panic("implement me")
}
