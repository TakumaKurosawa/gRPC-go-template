package middleware

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/api/option"
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
	opt := option.WithCredentialsFile("credential.json")

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
			log.Println(err)
			return nil, err
		}

		// JWT の検証
		authedUserToken, err := fa.client.VerifyIDToken(ctx, token)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// ユーザデータの取得
		userData, err := fa.client.GetUser(ctx, authedUserToken.UID)
		if err != nil {
			return nil, err
		}

		// contextにuidを格納
		ctx = context.WithValue(ctx, AuthedUserKey, userData.UID)
		return ctx, nil
	}
}
