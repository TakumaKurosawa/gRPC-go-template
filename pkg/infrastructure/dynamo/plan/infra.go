package planinfra

import (
	"context"
	"dataflow/pkg/domain/entity/plan"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/domain/repository/plan"
	"dataflow/pkg/infrastructure/dynamo"
	"dataflow/pkg/terrors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"google.golang.org/grpc/codes"
)

type planRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) planrepository.Repository {
	return &planRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

const tableName = "plans"

func (u *planRepositoryImpliment) InsertPlan(ctx context.Context, masterTx repository.MasterTx, entity *planentity.Plan) (*planentity.Plan, error) {
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

	return convertToPlanEntity(dto), nil
}

func (u *planRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*planentity.Plan, error) {
	// TODO implement me
	panic("implement me")
}

func (u *planRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) ([]*planentity.Plan, error) {
	// TODO implement me
	panic("implement me")
}
