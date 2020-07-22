package middleware

import (
	"context"
	"dataflow/pkg/terrors"
	"dataflow/pkg/tlog"
	"fmt"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func UnaryErrorHandling() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)

		if err != nil {
			// エラーログ出力
			uid := ctx.Value(string(AuthedUserKey))
			if uid == "" {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]:%v>", err))
			} else {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[%s]:%v>", uid, err))
			}

			// エラーレスポンスの送信
			if err != nil {
				var dataflowError *terrors.DataflowError
				if ok := xerrors.As(err, &dataflowError); ok {
					return nil, dataflowError
				}
				return nil, err
			}
		}

		return resp, nil
	}
}

func StreamErrorHandling() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		err := handler(srv, ss)
		ctx := ss.Context()

		if err != nil {
			// エラーログ出力
			uid := ctx.Value(string(AuthedUserKey))
			if uid == "" {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]:%v>", err))
			} else {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[%s]:%v>", uid, err))
			}

			// エラーレスポンスの送信
			var dataflowError *terrors.DataflowError
			if ok := xerrors.As(err, &dataflowError); ok {
				return dataflowError
			}
			return err
		}

		return nil
	}
}
