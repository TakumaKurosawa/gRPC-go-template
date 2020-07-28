package middleware

import (
	"context"
	"dataflow/pkg/terrors"
	"dataflow/pkg/tlog"
	"fmt"

	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
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
			uid, ok := ctx.Value(AuthedUserKey).(string)
			if !ok {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]:%v>", err))
			} else {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[%v]:%v>", uid, err))
			}

			// エラーレスポンスの送信
			var dataflowError *terrors.DataflowError
			if ok := xerrors.As(err, &dataflowError); ok {
				st := status.New(dataflowError.ErrorCode, "some error occurred")
				errResJP := &errdetails.LocalizedMessage{
					Locale:  "ja-JP",
					Message: dataflowError.ErrorMessageJP,
				}
				errResEN := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: dataflowError.ErrorMessageEN,
				}
				detailsErr, _ := st.WithDetails(errResJP, errResEN)
				return nil, detailsErr.Err()
			}
			return nil, err

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
			uid, ok := ctx.Value(AuthedUserKey).(string)
			if !ok {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[Unknown]:%v>", err))
			} else {
				tlog.GetAppLogger().Error(fmt.Sprintf("<[%v]:%v>", uid, err))
			}

			// エラーレスポンスの送信
			var dataflowError *terrors.DataflowError
			if ok := xerrors.As(err, &dataflowError); ok {
				st := status.New(dataflowError.ErrorCode, "some error occurred")
				errResJP := &errdetails.LocalizedMessage{
					Locale:  "ja-JP",
					Message: dataflowError.ErrorMessageJP,
				}
				errResEN := &errdetails.LocalizedMessage{
					Locale:  "en-US",
					Message: dataflowError.ErrorMessageEN,
				}
				detailsErr, _ := st.WithDetails(errResJP, errResEN)
				return detailsErr.Err()
			}
			return err
		}

		return nil
	}
}
