package mocks

import (
	"context"
	mainpb "github.com/auto-check/protocol-buffer/golang/main"
	gomock "github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockMainClient struct {
	ctrl *gomock.Controller
	recorder *mockMainClientRecorder
}

type mockMainClientRecorder struct {
	mock *MockMainClient
}

func NewMockMainClient(ctrl *gomock.Controller) *MockMainClient{
	mock := &MockMainClient{ctrl: ctrl}
	mock.recorder = &mockMainClientRecorder{mock}
	return mock
}

func (m *MockMainClient) EXPECT() *mockMainClientRecorder {
	return m.recorder
}

func (m *MockMainClient) CreateMacro(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	s := []interface{}{ctx, in}
	for _, opt:= range opts {
		s = append(s, opt)
	}

	ret := m.ctrl.Call(m, "CreateMacro", s...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder) CreateMacro(arg0, arg1 interface{}) *gomock.Call {
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateMacro", s...)
}


func (m *MockMainClient) GetLogListWithID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*mainpb.GetLogListWithIDResponse, error){
	s := []interface{}{ctx, in}
	for _, opt:= range opts {
		s = append(s, opt)
	}

	ret := m.ctrl.Call(m, "GetLogListWithID", s...)
	ret0, _ := ret[0].(*mainpb.GetLogListWithIDResponse)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder) GetLogListWithID(arg0, arg1 interface{}) *gomock.Call {
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "GetLogListWithID", s...)
}

func (m *MockMainClient) GetMacroStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*mainpb.GetMacroStatusResponse, error){
	s := []interface{}{ctx, in}
	for _, opt:= range opts {
		s = append(s, opt)
	}

	ret := m.ctrl.Call(m, "GetMacroStatus", s...)
	ret0, _ := ret[0].(*mainpb.GetMacroStatusResponse)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder) GetMacroStatus(arg0, arg1 interface{}) *gomock.Call {
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "GetMacroStatus", s...)
}

func (m *MockMainClient)  GetMacroSecret(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*mainpb.GetMacroSecretResponse, error){
	s := []interface{}{ctx, in}
	for _, opt:= range opts {
		s = append(s, opt)
	}

	ret := m.ctrl.Call(m, "GetMacroSecret", s...)
	ret0, _ := ret[0].(*mainpb.GetMacroSecretResponse)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder)  GetMacroSecret(arg0, arg1 interface{}) *gomock.Call{
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "GetMacroSecret", s...)
}

func (m *MockMainClient) ControlMacro(ctx context.Context, in *mainpb.ControlMacroRequest, opts ...grpc.CallOption) (*emptypb.Empty, error){
	s := []interface{}{ctx, in}
	for _, opt:= range opts {
		s = append(s, opt)
	}

	ret := m.ctrl.Call(m, "ControlMacro", s...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *mockMainClientRecorder) ControlMacro(arg0, arg1 interface{}) *gomock.Call {
	s := []interface{}{arg0, arg1}
	return mr.mock.ctrl.RecordCall(mr.mock, "ControlMacro", s...)
}