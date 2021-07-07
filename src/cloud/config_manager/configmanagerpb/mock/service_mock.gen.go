// Code generated by MockGen. DO NOT EDIT.
// Source: service.pb.go

// Package mock_configmanagerpb is a generated GoMock package.
package mock_configmanagerpb

import (
	gomock "github.com/golang/mock/gomock"
)

// MockConfigManagerServiceClient is a mock of ConfigManagerServiceClient interface.
type MockConfigManagerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerServiceClientMockRecorder
}

// MockConfigManagerServiceClientMockRecorder is the mock recorder for MockConfigManagerServiceClient.
type MockConfigManagerServiceClientMockRecorder struct {
	mock *MockConfigManagerServiceClient
}

// NewMockConfigManagerServiceClient creates a new mock instance.
func NewMockConfigManagerServiceClient(ctrl *gomock.Controller) *MockConfigManagerServiceClient {
	mock := &MockConfigManagerServiceClient{ctrl: ctrl}
	mock.recorder = &MockConfigManagerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigManagerServiceClient) EXPECT() *MockConfigManagerServiceClientMockRecorder {
	return m.recorder
}

// MockConfigManagerServiceServer is a mock of ConfigManagerServiceServer interface.
type MockConfigManagerServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerServiceServerMockRecorder
}

// MockConfigManagerServiceServerMockRecorder is the mock recorder for MockConfigManagerServiceServer.
type MockConfigManagerServiceServerMockRecorder struct {
	mock *MockConfigManagerServiceServer
}

// NewMockConfigManagerServiceServer creates a new mock instance.
func NewMockConfigManagerServiceServer(ctrl *gomock.Controller) *MockConfigManagerServiceServer {
	mock := &MockConfigManagerServiceServer{ctrl: ctrl}
	mock.recorder = &MockConfigManagerServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigManagerServiceServer) EXPECT() *MockConfigManagerServiceServerMockRecorder {
	return m.recorder
}