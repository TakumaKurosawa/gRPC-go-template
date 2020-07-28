package user

import (
	"context"
	"dataflow/pkg/api/middleware"
	userinteractor "dataflow/pkg/api/usecase/user"
	"dataflow/pkg/pb"
	"dataflow/pkg/terrors"

	"google.golang.org/grpc/codes"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
	uid, ok := ctx.Value(middleware.AuthedUserKey).(string)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		return nil, terrors.Newf(codes.Unauthenticated, errMessageJP, errMessageEN)
	}

	insertedUser, err := s.userInteractor.CreateNewUser(ctx, uid, req.Name, req.Thumbnail)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return &pb.UserInfo{
		Name:      insertedUser.Name,
		Thumbnail: insertedUser.Thumbnail,
	}, nil
}

//func (s *Server) GetUserProfile(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error) {
//	uid, ok := ctx.Value(middleware.AuthedUserKey).(string)
//	if !ok {
//		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
//		errMessageEN := "The content blocked because user is not certified."
//		return nil, terrors.Newf(codes.Unauthenticated, errMessageJP, errMessageEN)
//	}
//
//	userData, err := s.userInteractor.GetUserProfile(ctx, uid)
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.UserInfo{
//		Name:      userData.Name,
//		Thumbnail: userData.Thumbnail,
//	}, nil
//}
//
//func (s *Server) GetAllUsers(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserInfo, error){
//	users, err := s.userInteractor.GetAll(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	return
//}
