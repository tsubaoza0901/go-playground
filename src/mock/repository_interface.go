// Code generated by MockGen. DO NOT EDIT.
// Source: repository_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	dto "go-playground/m/v1/src/usecase/repository/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBalanceRepository is a mock of IBalanceRepository interface.
type MockIBalanceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIBalanceRepositoryMockRecorder
}

// MockIBalanceRepositoryMockRecorder is the mock recorder for MockIBalanceRepository.
type MockIBalanceRepositoryMockRecorder struct {
	mock *MockIBalanceRepository
}

// NewMockIBalanceRepository creates a new mock instance.
func NewMockIBalanceRepository(ctrl *gomock.Controller) *MockIBalanceRepository {
	mock := &MockIBalanceRepository{ctrl: ctrl}
	mock.recorder = &MockIBalanceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBalanceRepository) EXPECT() *MockIBalanceRepositoryMockRecorder {
	return m.recorder
}

// FetchBalanceByUserID mocks base method.
func (m *MockIBalanceRepository) FetchBalanceByUserID(ctx context.Context, userID uint) (*dto.FetchBlanceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBalanceByUserID", ctx, userID)
	ret0, _ := ret[0].(*dto.FetchBlanceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBalanceByUserID indicates an expected call of FetchBalanceByUserID.
func (mr *MockIBalanceRepositoryMockRecorder) FetchBalanceByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBalanceByUserID", reflect.TypeOf((*MockIBalanceRepository)(nil).FetchBalanceByUserID), ctx, userID)
}

// RegisterBalance mocks base method.
func (m *MockIBalanceRepository) RegisterBalance(ctx context.Context, createBalanceDTO dto.RegisterBalance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBalance", ctx, createBalanceDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterBalance indicates an expected call of RegisterBalance.
func (mr *MockIBalanceRepositoryMockRecorder) RegisterBalance(ctx, createBalanceDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBalance", reflect.TypeOf((*MockIBalanceRepository)(nil).RegisterBalance), ctx, createBalanceDTO)
}

// UpdateBalance mocks base method.
func (m *MockIBalanceRepository) UpdateBalance(ctx context.Context, updateBalanceDTO dto.UpdateBalance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", ctx, updateBalanceDTO)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockIBalanceRepositoryMockRecorder) UpdateBalance(ctx, updateBalanceDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockIBalanceRepository)(nil).UpdateBalance), ctx, updateBalanceDTO)
}

// MockIGradeRepository is a mock of IGradeRepository interface.
type MockIGradeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIGradeRepositoryMockRecorder
}

// MockIGradeRepositoryMockRecorder is the mock recorder for MockIGradeRepository.
type MockIGradeRepositoryMockRecorder struct {
	mock *MockIGradeRepository
}

// NewMockIGradeRepository creates a new mock instance.
func NewMockIGradeRepository(ctrl *gomock.Controller) *MockIGradeRepository {
	mock := &MockIGradeRepository{ctrl: ctrl}
	mock.recorder = &MockIGradeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGradeRepository) EXPECT() *MockIGradeRepositoryMockRecorder {
	return m.recorder
}

// FetchGradeList mocks base method.
func (m *MockIGradeRepository) FetchGradeList(arg0 context.Context) (*dto.FetchGradeListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchGradeList", arg0)
	ret0, _ := ret[0].(*dto.FetchGradeListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchGradeList indicates an expected call of FetchGradeList.
func (mr *MockIGradeRepositoryMockRecorder) FetchGradeList(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchGradeList", reflect.TypeOf((*MockIGradeRepository)(nil).FetchGradeList), arg0)
}

// MockIDealHistoryRepository is a mock of IDealHistoryRepository interface.
type MockIDealHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIDealHistoryRepositoryMockRecorder
}

// MockIDealHistoryRepositoryMockRecorder is the mock recorder for MockIDealHistoryRepository.
type MockIDealHistoryRepositoryMockRecorder struct {
	mock *MockIDealHistoryRepository
}

// NewMockIDealHistoryRepository creates a new mock instance.
func NewMockIDealHistoryRepository(ctrl *gomock.Controller) *MockIDealHistoryRepository {
	mock := &MockIDealHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockIDealHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDealHistoryRepository) EXPECT() *MockIDealHistoryRepositoryMockRecorder {
	return m.recorder
}

// FetchDealHistoriesByUserID mocks base method.
func (m *MockIDealHistoryRepository) FetchDealHistoriesByUserID(ctx context.Context, userID uint) (*dto.FetchDealHistoryListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDealHistoriesByUserID", ctx, userID)
	ret0, _ := ret[0].(*dto.FetchDealHistoryListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDealHistoriesByUserID indicates an expected call of FetchDealHistoriesByUserID.
func (mr *MockIDealHistoryRepositoryMockRecorder) FetchDealHistoriesByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDealHistoriesByUserID", reflect.TypeOf((*MockIDealHistoryRepository)(nil).FetchDealHistoriesByUserID), ctx, userID)
}

// RegisterDealHistory mocks base method.
func (m *MockIDealHistoryRepository) RegisterDealHistory(arg0 context.Context, arg1 dto.RegisterDealHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterDealHistory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterDealHistory indicates an expected call of RegisterDealHistory.
func (mr *MockIDealHistoryRepositoryMockRecorder) RegisterDealHistory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDealHistory", reflect.TypeOf((*MockIDealHistoryRepository)(nil).RegisterDealHistory), arg0, arg1)
}

// MockIUserManagementRepository is a mock of IUserManagementRepository interface.
type MockIUserManagementRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserManagementRepositoryMockRecorder
}

// MockIUserManagementRepositoryMockRecorder is the mock recorder for MockIUserManagementRepository.
type MockIUserManagementRepositoryMockRecorder struct {
	mock *MockIUserManagementRepository
}

// NewMockIUserManagementRepository creates a new mock instance.
func NewMockIUserManagementRepository(ctrl *gomock.Controller) *MockIUserManagementRepository {
	mock := &MockIUserManagementRepository{ctrl: ctrl}
	mock.recorder = &MockIUserManagementRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserManagementRepository) EXPECT() *MockIUserManagementRepositoryMockRecorder {
	return m.recorder
}

// FetchUserByEmail mocks base method.
func (m *MockIUserManagementRepository) FetchUserByEmail(ctx context.Context, email string) (*dto.FetchUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserByEmail", ctx, email)
	ret0, _ := ret[0].(*dto.FetchUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserByEmail indicates an expected call of FetchUserByEmail.
func (mr *MockIUserManagementRepositoryMockRecorder) FetchUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserByEmail", reflect.TypeOf((*MockIUserManagementRepository)(nil).FetchUserByEmail), ctx, email)
}

// FetchUserByID mocks base method.
func (m *MockIUserManagementRepository) FetchUserByID(ctx context.Context, id uint) (*dto.FetchUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserByID", ctx, id)
	ret0, _ := ret[0].(*dto.FetchUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserByID indicates an expected call of FetchUserByID.
func (mr *MockIUserManagementRepositoryMockRecorder) FetchUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserByID", reflect.TypeOf((*MockIUserManagementRepository)(nil).FetchUserByID), ctx, id)
}

// FetchUserList mocks base method.
func (m *MockIUserManagementRepository) FetchUserList(ctx context.Context) (*dto.FetchUserListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserList", ctx)
	ret0, _ := ret[0].(*dto.FetchUserListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserList indicates an expected call of FetchUserList.
func (mr *MockIUserManagementRepositoryMockRecorder) FetchUserList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserList", reflect.TypeOf((*MockIUserManagementRepository)(nil).FetchUserList), ctx)
}

// RegisterUser mocks base method.
func (m *MockIUserManagementRepository) RegisterUser(arg0 context.Context, arg1 dto.RegisterUser) (*dto.FetchUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0, arg1)
	ret0, _ := ret[0].(*dto.FetchUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserManagementRepositoryMockRecorder) RegisterUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserManagementRepository)(nil).RegisterUser), arg0, arg1)
}

// MockITransactionManagementRepository is a mock of ITransactionManagementRepository interface.
type MockITransactionManagementRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionManagementRepositoryMockRecorder
}

// MockITransactionManagementRepositoryMockRecorder is the mock recorder for MockITransactionManagementRepository.
type MockITransactionManagementRepositoryMockRecorder struct {
	mock *MockITransactionManagementRepository
}

// NewMockITransactionManagementRepository creates a new mock instance.
func NewMockITransactionManagementRepository(ctrl *gomock.Controller) *MockITransactionManagementRepository {
	mock := &MockITransactionManagementRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionManagementRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionManagementRepository) EXPECT() *MockITransactionManagementRepositoryMockRecorder {
	return m.recorder
}

// Transaction mocks base method.
func (m *MockITransactionManagementRepository) Transaction(ctx context.Context, fc func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transaction", ctx, fc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transaction indicates an expected call of Transaction.
func (mr *MockITransactionManagementRepositoryMockRecorder) Transaction(ctx, fc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transaction", reflect.TypeOf((*MockITransactionManagementRepository)(nil).Transaction), ctx, fc)
}
