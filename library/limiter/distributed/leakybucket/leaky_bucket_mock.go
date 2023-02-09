// Code generated by MockGen. DO NOT EDIT.
// Source: leaky_bucket.go

// Package leakybucket is a generated GoMock package.
package leakybucket

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLeakyBucket is a mock of LeakyBucket interface.
type MockLeakyBucket struct {
	ctrl     *gomock.Controller
	recorder *MockLeakyBucketMockRecorder
}

// MockLeakyBucketMockRecorder is the mock recorder for MockLeakyBucket.
type MockLeakyBucketMockRecorder struct {
	mock *MockLeakyBucket
}

// NewMockLeakyBucket creates a new mock instance.
func NewMockLeakyBucket(ctrl *gomock.Controller) *MockLeakyBucket {
	mock := &MockLeakyBucket{ctrl: ctrl}
	mock.recorder = &MockLeakyBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeakyBucket) EXPECT() *MockLeakyBucketMockRecorder {
	return m.recorder
}

// Allow mocks base method.
func (m *MockLeakyBucket) Allow(ctx context.Context, key string, requestCount int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow", ctx, key, requestCount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allow indicates an expected call of Allow.
func (mr *MockLeakyBucketMockRecorder) Allow(ctx, key, requestCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockLeakyBucket)(nil).Allow), ctx, key, requestCount)
}

// SetBurst mocks base method.
func (m *MockLeakyBucket) SetBurst(burst int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBurst", burst)
}

// SetBurst indicates an expected call of SetBurst.
func (mr *MockLeakyBucketMockRecorder) SetBurst(burst interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBurst", reflect.TypeOf((*MockLeakyBucket)(nil).SetBurst), burst)
}

// SetLimit mocks base method.
func (m *MockLeakyBucket) SetLimit(rate int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLimit", rate)
}

// SetLimit indicates an expected call of SetLimit.
func (mr *MockLeakyBucketMockRecorder) SetLimit(rate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLimit", reflect.TypeOf((*MockLeakyBucket)(nil).SetLimit), rate)
}
