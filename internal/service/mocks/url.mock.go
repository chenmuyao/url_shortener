// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/url.go
//
// Generated by this command:
//
//	mockgen -source=./internal/service/url.go -package=urlsvcmock -destination=./internal/service/mocks/url.mock.go
//

// Package urlsvcmock is a generated GoMock package.
package urlsvcmock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUrlShortenerSvc is a mock of UrlShortenerSvc interface.
type MockUrlShortenerSvc struct {
	ctrl     *gomock.Controller
	recorder *MockUrlShortenerSvcMockRecorder
	isgomock struct{}
}

// MockUrlShortenerSvcMockRecorder is the mock recorder for MockUrlShortenerSvc.
type MockUrlShortenerSvcMockRecorder struct {
	mock *MockUrlShortenerSvc
}

// NewMockUrlShortenerSvc creates a new mock instance.
func NewMockUrlShortenerSvc(ctrl *gomock.Controller) *MockUrlShortenerSvc {
	mock := &MockUrlShortenerSvc{ctrl: ctrl}
	mock.recorder = &MockUrlShortenerSvcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUrlShortenerSvc) EXPECT() *MockUrlShortenerSvcMockRecorder {
	return m.recorder
}

// GetFull mocks base method.
func (m *MockUrlShortenerSvc) GetFull(ctx context.Context, shortID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFull", ctx, shortID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFull indicates an expected call of GetFull.
func (mr *MockUrlShortenerSvcMockRecorder) GetFull(ctx, shortID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFull", reflect.TypeOf((*MockUrlShortenerSvc)(nil).GetFull), ctx, shortID)
}

// Shorten mocks base method.
func (m *MockUrlShortenerSvc) Shorten(ctx context.Context, fullUrl string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shorten", ctx, fullUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shorten indicates an expected call of Shorten.
func (mr *MockUrlShortenerSvcMockRecorder) Shorten(ctx, fullUrl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shorten", reflect.TypeOf((*MockUrlShortenerSvc)(nil).Shorten), ctx, fullUrl)
}
