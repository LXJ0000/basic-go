// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/cache/article.go
//
// Generated by this command:
//
//	mockgen -source=./internal/repository/cache/article.go -package=cachemock -destination=./internal/repository/cache/mocks/article_gen.go
//

// Package cachemock is a generated GoMock package.
package cachemock

import (
	gomock "go.uber.org/mock/gomock"
)

// MockArticleCACHE is a mock of ArticleCACHE interface.
type MockArticleCACHE struct {
	ctrl     *gomock.Controller
	recorder *MockArticleCACHEMockRecorder
}

// MockArticleCACHEMockRecorder is the mock recorder for MockArticleCACHE.
type MockArticleCACHEMockRecorder struct {
	mock *MockArticleCACHE
}

// NewMockArticleCACHE creates a new mock instance.
func NewMockArticleCACHE(ctrl *gomock.Controller) *MockArticleCACHE {
	mock := &MockArticleCACHE{ctrl: ctrl}
	mock.recorder = &MockArticleCACHEMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleCACHE) EXPECT() *MockArticleCACHEMockRecorder {
	return m.recorder
}
