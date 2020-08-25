// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_profile is a generated GoMock package.
package mock_profile

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	profilepb "pixielabs.ai/pixielabs/src/cloud/profile/profilepb"
	proto "pixielabs.ai/pixielabs/src/common/uuid/proto"
	reflect "reflect"
)

// MockProfileServiceClient is a mock of ProfileServiceClient interface
type MockProfileServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceClientMockRecorder
}

// MockProfileServiceClientMockRecorder is the mock recorder for MockProfileServiceClient
type MockProfileServiceClientMockRecorder struct {
	mock *MockProfileServiceClient
}

// NewMockProfileServiceClient creates a new mock instance
func NewMockProfileServiceClient(ctrl *gomock.Controller) *MockProfileServiceClient {
	mock := &MockProfileServiceClient{ctrl: ctrl}
	mock.recorder = &MockProfileServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProfileServiceClient) EXPECT() *MockProfileServiceClientMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockProfileServiceClient) CreateUser(ctx context.Context, in *profilepb.CreateUserRequest, opts ...grpc.CallOption) (*proto.UUID, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockProfileServiceClientMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateUser), varargs...)
}

// GetUser mocks base method
func (m *MockProfileServiceClient) GetUser(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockProfileServiceClientMockRecorder) GetUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUser), varargs...)
}

// GetUserByEmail mocks base method
func (m *MockProfileServiceClient) GetUserByEmail(ctx context.Context, in *profilepb.GetUserByEmailRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserByEmail", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockProfileServiceClientMockRecorder) GetUserByEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceClient)(nil).GetUserByEmail), varargs...)
}

// CreateOrgAndUser mocks base method
func (m *MockProfileServiceClient) CreateOrgAndUser(ctx context.Context, in *profilepb.CreateOrgAndUserRequest, opts ...grpc.CallOption) (*profilepb.CreateOrgAndUserResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrgAndUser", varargs...)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser
func (mr *MockProfileServiceClientMockRecorder) CreateOrgAndUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceClient)(nil).CreateOrgAndUser), varargs...)
}

// GetOrg mocks base method
func (m *MockProfileServiceClient) GetOrg(ctx context.Context, in *proto.UUID, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrg", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg
func (mr *MockProfileServiceClientMockRecorder) GetOrg(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrg), varargs...)
}

// GetOrgByDomain mocks base method
func (m *MockProfileServiceClient) GetOrgByDomain(ctx context.Context, in *profilepb.GetOrgByDomainRequest, opts ...grpc.CallOption) (*profilepb.OrgInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgByDomain", varargs...)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain
func (mr *MockProfileServiceClientMockRecorder) GetOrgByDomain(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrgByDomain), varargs...)
}

// UpdateUser mocks base method
func (m *MockProfileServiceClient) UpdateUser(ctx context.Context, in *profilepb.UpdateUserRequest, opts ...grpc.CallOption) (*profilepb.UserInfo, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUser", varargs...)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockProfileServiceClientMockRecorder) UpdateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceClient)(nil).UpdateUser), varargs...)
}

// GetOrgs mocks base method
func (m *MockProfileServiceClient) GetOrgs(ctx context.Context, in *profilepb.GetOrgsRequest, opts ...grpc.CallOption) (*profilepb.GetOrgsResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrgs", varargs...)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs
func (mr *MockProfileServiceClientMockRecorder) GetOrgs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockProfileServiceClient)(nil).GetOrgs), varargs...)
}

// MockProfileServiceServer is a mock of ProfileServiceServer interface
type MockProfileServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServiceServerMockRecorder
}

// MockProfileServiceServerMockRecorder is the mock recorder for MockProfileServiceServer
type MockProfileServiceServerMockRecorder struct {
	mock *MockProfileServiceServer
}

// NewMockProfileServiceServer creates a new mock instance
func NewMockProfileServiceServer(ctrl *gomock.Controller) *MockProfileServiceServer {
	mock := &MockProfileServiceServer{ctrl: ctrl}
	mock.recorder = &MockProfileServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProfileServiceServer) EXPECT() *MockProfileServiceServerMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockProfileServiceServer) CreateUser(arg0 context.Context, arg1 *profilepb.CreateUserRequest) (*proto.UUID, error) {
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*proto.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockProfileServiceServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateUser), arg0, arg1)
}

// GetUser mocks base method
func (m *MockProfileServiceServer) GetUser(arg0 context.Context, arg1 *proto.UUID) (*profilepb.UserInfo, error) {
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockProfileServiceServerMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUser), arg0, arg1)
}

// GetUserByEmail mocks base method
func (m *MockProfileServiceServer) GetUserByEmail(arg0 context.Context, arg1 *profilepb.GetUserByEmailRequest) (*profilepb.UserInfo, error) {
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockProfileServiceServerMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockProfileServiceServer)(nil).GetUserByEmail), arg0, arg1)
}

// CreateOrgAndUser mocks base method
func (m *MockProfileServiceServer) CreateOrgAndUser(arg0 context.Context, arg1 *profilepb.CreateOrgAndUserRequest) (*profilepb.CreateOrgAndUserResponse, error) {
	ret := m.ctrl.Call(m, "CreateOrgAndUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.CreateOrgAndUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrgAndUser indicates an expected call of CreateOrgAndUser
func (mr *MockProfileServiceServerMockRecorder) CreateOrgAndUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrgAndUser", reflect.TypeOf((*MockProfileServiceServer)(nil).CreateOrgAndUser), arg0, arg1)
}

// GetOrg mocks base method
func (m *MockProfileServiceServer) GetOrg(arg0 context.Context, arg1 *proto.UUID) (*profilepb.OrgInfo, error) {
	ret := m.ctrl.Call(m, "GetOrg", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrg indicates an expected call of GetOrg
func (mr *MockProfileServiceServerMockRecorder) GetOrg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrg", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrg), arg0, arg1)
}

// GetOrgByDomain mocks base method
func (m *MockProfileServiceServer) GetOrgByDomain(arg0 context.Context, arg1 *profilepb.GetOrgByDomainRequest) (*profilepb.OrgInfo, error) {
	ret := m.ctrl.Call(m, "GetOrgByDomain", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.OrgInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgByDomain indicates an expected call of GetOrgByDomain
func (mr *MockProfileServiceServerMockRecorder) GetOrgByDomain(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgByDomain", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrgByDomain), arg0, arg1)
}

// UpdateUser mocks base method
func (m *MockProfileServiceServer) UpdateUser(arg0 context.Context, arg1 *profilepb.UpdateUserRequest) (*profilepb.UserInfo, error) {
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockProfileServiceServerMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockProfileServiceServer)(nil).UpdateUser), arg0, arg1)
}

// GetOrgs mocks base method
func (m *MockProfileServiceServer) GetOrgs(arg0 context.Context, arg1 *profilepb.GetOrgsRequest) (*profilepb.GetOrgsResponse, error) {
	ret := m.ctrl.Call(m, "GetOrgs", arg0, arg1)
	ret0, _ := ret[0].(*profilepb.GetOrgsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgs indicates an expected call of GetOrgs
func (mr *MockProfileServiceServerMockRecorder) GetOrgs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgs", reflect.TypeOf((*MockProfileServiceServer)(nil).GetOrgs), arg0, arg1)
}
