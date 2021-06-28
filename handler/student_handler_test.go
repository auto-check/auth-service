package handler

import (
	"auth/usecase/mocks"
	"bytes"
	"context"
	"encoding/json"
	"github.com/auto-check/common-module/jwt"
	"github.com/auto-check/common-module/model"
	mcmocks "github.com/auto-check/main-service/client/mocks"
	authpb "github.com/auto-check/protocol-buffer/golang/auth"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"os"
	"strings"
	"testing"
)

func init() {
	err := 	godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

type DsmOAuthRequest struct {
	ID string `json:"id"`
	Password string `json:"password"`
	RedirectURL string `json:"redirect_url"`
	ClientID string `json:"client_id"`
}

type DsmOAuthResponse struct {
	Location string `json:"location"`
}

func TestLoginAuth(t *testing.T) {
	at, rt, err := jwt.GenerateToken(1)
	mockStudent := new(mocks.StudentUsecase)
	mockStudent.On("LoginAuth", &model.Student{
		Name: "조호원",
		Gcn: "2318",
		Email: "201118jhw@dsm.hs.kr",
	}).Return(at, rt, err)

	/* OAuth로 부터 authorization code 가져오기*/
	var reqByte []byte
	reqBody := bytes.NewBuffer(reqByte)
	err = json.NewEncoder(reqBody).Encode(DsmOAuthRequest{
		ID: 		 os.Getenv("TEST_ID"),
		Password:    os.Getenv("TEST_PASS"),
		RedirectURL: os.Getenv("REDIRECT_URL"),
		ClientID:    os.Getenv("CLIENT_ID"),
	})

	resp, err := http.Post("https://developer-api.dsmkr.com/dsmauth/login",
		"application/json", reqBody)

	var respBody DsmOAuthResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	code := strings.Split(respBody.Location, "=")[1]

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	/* handler mock 객체 생성 */
	mcMock := mcmocks.NewMockMainClient(ctrl)
	ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{"student_id": []string{"1"}})
	mcMock.EXPECT().CreateMacro(
		ctx,
		&emptypb.Empty{},
	).Return(&emptypb.Empty{}, nil)
	sh := NewStudentHandler(grpc.NewServer(), mockStudent, mcMock)

	/* authorization code로 LoginAuth 테스트*/
	loginAuthResponse, err := sh.LoginAuth(ctx, &authpb.LoginAuthRequest{
		Code: code,
	})

	studentID,  err := jwt.ParseStudentIDFromToken(loginAuthResponse.AccessToken)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), studentID)
}
