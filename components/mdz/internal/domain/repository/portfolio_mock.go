// Code generated by MockGen. DO NOT EDIT.
// Source: /home/max/Workspace/midaz/components/mdz/internal/domain/repository/portfolio.go
//
// Generated by this command:
//
//	mockgen -source=/home/max/Workspace/midaz/components/mdz/internal/domain/repository/portfolio.go -destination=/home/max/Workspace/midaz/components/mdz/internal/domain/repository/portfolio_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	mmodel "github.com/LerianStudio/midaz/common/mmodel"
	gomock "go.uber.org/mock/gomock"
)

// MockPortfolio is a mock of Portfolio interface.
type MockPortfolio struct {
	ctrl     *gomock.Controller
	recorder *MockPortfolioMockRecorder
	isgomock struct{}
}

// MockPortfolioMockRecorder is the mock recorder for MockPortfolio.
type MockPortfolioMockRecorder struct {
	mock *MockPortfolio
}

// NewMockPortfolio creates a new mock instance.
func NewMockPortfolio(ctrl *gomock.Controller) *MockPortfolio {
	mock := &MockPortfolio{ctrl: ctrl}
	mock.recorder = &MockPortfolioMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortfolio) EXPECT() *MockPortfolioMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPortfolio) Create(organizationID, ledgerID string, inp mmodel.CreatePortfolioInput) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", organizationID, ledgerID, inp)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPortfolioMockRecorder) Create(organizationID, ledgerID, inp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPortfolio)(nil).Create), organizationID, ledgerID, inp)
}

// Get mocks base method.
func (m *MockPortfolio) Get(organizationID, ledgerID string, limit, page int) (*mmodel.Portfolios, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", organizationID, ledgerID, limit, page)
	ret0, _ := ret[0].(*mmodel.Portfolios)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPortfolioMockRecorder) Get(organizationID, ledgerID, limit, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPortfolio)(nil).Get), organizationID, ledgerID, limit, page)
}

// GetByID mocks base method.
func (m *MockPortfolio) GetByID(organizationID, ledgerID, portfolioID string) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", organizationID, ledgerID, portfolioID)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockPortfolioMockRecorder) GetByID(organizationID, ledgerID, portfolioID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPortfolio)(nil).GetByID), organizationID, ledgerID, portfolioID)
}

// Update mocks base method.
func (m *MockPortfolio) Update(organizationID, ledgerID, portfolioID string, inp mmodel.UpdatePortfolioInput) (*mmodel.Portfolio, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", organizationID, ledgerID, portfolioID, inp)
	ret0, _ := ret[0].(*mmodel.Portfolio)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPortfolioMockRecorder) Update(organizationID, ledgerID, portfolioID, inp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPortfolio)(nil).Update), organizationID, ledgerID, portfolioID, inp)
}
