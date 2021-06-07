package handler

import (
	"auth/model"
	"auth/module/client"
	authpb "auth/protocol-buffer/golang/auth"
	"auth/usecase"
	"context"
	"google.golang.org/grpc"
)

type StudentHandler struct {
	authpb.AuthServer
	su usecase.StudentUsecase
}

func NewStudentHandler(gserver *grpc.Server, us usecase.StudentUsecase) *StudentHandler {
	handler := &StudentHandler{
		su: us,
	}
	authpb.RegisterAuthServer(gserver, handler)
	return handler
}

func (sh *StudentHandler) LoginAuth(ctx context.Context, r *authpb.LoginAuthRequest) (*authpb.LoginAuthResponse, error) {
	oauthToken, err := client.GetOauthAccessToken(r.Code)
	if err != nil {
		return nil, err
	}

	s, err := client.GetOauthInfo(oauthToken)
	if err != nil {
		return nil, err
	}

	at, rt, err := sh.su.LoginAuth(&model.Student{
		Name: s.Name,
		Gcn: s.Gcn,
		Email: s.Email,
	})
	if err != nil {
		return nil, err
	}

	return &authpb.LoginAuthResponse{AccessToken: at, RefreshToken: rt}, nil
}


func (sh *StudentHandler) GetStudentWithId(ctx context.Context, req *authpb.GetStudentWithIdRequest) (*authpb.GetStudentWithIdResponse, error) {

	return nil, nil
}

