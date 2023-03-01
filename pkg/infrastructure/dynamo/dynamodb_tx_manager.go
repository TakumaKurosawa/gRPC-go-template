package dynamo

import (
	"context"
	"dataflow/pkg/domain/repository"
	"dataflow/pkg/terrors"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc/codes"
)

type dbMasterTxManager struct {
	db *dynamodb.DynamoDB
}

func NewDBMasterTxManager(db *dynamodb.DynamoDB) repository.MasterTxManager {
	return &dbMasterTxManager{db}
}

type dbMasterTx struct {
	tx *dynamodb.DynamoDB
}

func (m *dbMasterTxManager) Transaction(ctx context.Context, f func(ctx context.Context, masterTx repository.MasterTx) error) error {
	if err := f(ctx, &dbMasterTx{m.db}); err != nil {
		return terrors.Stack(err)
	}
	return nil
}

func (m *dbMasterTx) Commit() error {
	// do nothing
	return nil
}

func (m *dbMasterTx) Rollback() error {
	// do nothing
	return nil
}

func ExtractExecutor(masterTx repository.MasterTx) (*dynamodb.DynamoDB, error) {
	return ExtractTx(masterTx)
}

func ExtractTx(masterTx repository.MasterTx) (*dynamodb.DynamoDB, error) {
	// キャストする
	tx, ok := masterTx.(*dbMasterTx)
	if !ok {
		return nil, terrors.Newf(codes.Internal, "masterTxからdbMasterTxへのキャストに失敗しました。", "masterTx cannot cast to dbMasterTx")
	}
	return tx.tx, nil
}
