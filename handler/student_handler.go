package handler

import (
	"context"
	"github.com/auto-check/auth-service/usecase"
	"github.com/auto-check/common-module/client"
	"github.com/auto-check/common-module/model"
	authpb "github.com/auto-check/protocol-buffer/golang/auth"
	mainpb "github.com/auto-check/protocol-buffer/golang/main"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StudentHandler struct {
	authpb.AuthServer
	mc mainpb.MainClient
	su usecase.StudentUsecase
}

func NewStudentHandler(gserver *grpc.Server, us usecase.StudentUsecase, mc mainpb.MainClient) *StudentHandler {
	handler := &StudentHandler{
		su: us,
		mc: mc,
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

	ctx = metadata.NewOutgoingContext(ctx, metadata.MD{"authorization": []string{"bearer", at}})
	log.Println(metautils.ExtractOutgoing(ctx))
	_, err = sh.mc.CreateMacro(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &authpb.LoginAuthResponse{AccessToken: at, RefreshToken: rt}, nil
}

func (sh *StudentHandler) GetStudentWithId(ctx context.Context, req *authpb.GetStudentWithIdRequest) (*authpb.GetStudentWithIdResponse, error) {
	id, _ := ctx.Value("student_id").(int64)
	log.Println(id)

	return nil, nil
}

