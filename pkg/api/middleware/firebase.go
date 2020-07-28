package middleware

import (
	"context"
	"dataflow/pkg/tlog"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FirebaseAuth interface {
	MiddlewareFunc() grpc_auth.AuthFunc
}

type firebaseAuth struct {
	client *auth.Client
}

type contextKey string

const (
	AuthedUserKey contextKey = "AUTHED_UID"
)

func CreateFirebaseInstance() FirebaseAuth {
	ctx := context.Background()

	// get credential of firebase
	opt := option.WithCredentialsFile("pluslab-dataflow-firebase-adminsdk-gsdof-34cb5964c2.json")

	// firebase appの作成
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Panic(fmt.Errorf("error initializing firebase app: %v", err))
	}

	// firebase admin clientの作成
	client, err := app.Auth(ctx)
	if err != nil {
		log.Panic(fmt.Errorf("error initialize firebase instance. %v", err))
	}

	return &firebaseAuth{
		client: client,
	}
}

func (fa *firebaseAuth) MiddlewareFunc() grpc_auth.AuthFunc {
	return fa.middlewareImpl()
}

func (fa *firebaseAuth) middlewareImpl() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		// Authorizationヘッダーからjwtトークンを取得
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			tlog.GetAppLogger().Error(err.Error())
			log.Println(err)
			st := status.New(codes.Unauthenticated, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "認証情報が見つかりませんでした。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Authorization header is empty",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// JWT の検証
		authedUserToken, err := fa.client.VerifyIDToken(ctx, token)
		if err != nil {
			tlog.GetAppLogger().Error(err.Error())
			st := status.New(codes.Unauthenticated, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "トークンが有効ではありませんでした。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Your token is invalid.",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// ユーザデータの取得
		userData, err := fa.client.GetUser(ctx, authedUserToken.UID)
		if err != nil {
			tlog.GetAppLogger().Error(err.Error())
			st := status.New(codes.Internal, "authentication error occurred")
			errResJP := &errdetails.LocalizedMessage{
				Locale:  "ja-JP",
				Message: "サーバでエラーが発生しました。。",
			}
			errResEN := &errdetails.LocalizedMessage{
				Locale:  "en-US",
				Message: "Error occurred in server.",
			}
			detailsErr, _ := st.WithDetails(errResJP, errResEN)
			return nil, detailsErr.Err()
		}

		// contextにuidを格納
		ctx = context.WithValue(ctx, AuthedUserKey, userData.UID)
		return ctx, nil
	}
}
