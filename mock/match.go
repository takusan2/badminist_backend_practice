// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/okadatakuya/my_folder/dev/my_app/badminist/backend/domain/match.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/takuya-okada-01/badminist-backend/domain"
)

// MockIMatchRepository is a mock of IMatchRepository interface.
type MockIMatchRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIMatchRepositoryMockRecorder
}

// MockIMatchRepositoryMockRecorder is the mock recorder for MockIMatchRepository.
type MockIMatchRepositoryMockRecorder struct {
	mock *MockIMatchRepository
}

// NewMockIMatchRepository creates a new mock instance.
func NewMockIMatchRepository(ctrl *gomock.Controller) *MockIMatchRepository {
	mock := &MockIMatchRepository{ctrl: ctrl}
	mock.recorder = &MockIMatchRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMatchRepository) EXPECT() *MockIMatchRepositoryMockRecorder {
	return m.recorder
}

// DeleteMatch mocks base method.
func (m *MockIMatchRepository) DeleteMatch(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMatch", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMatch indicates an expected call of DeleteMatch.
func (mr *MockIMatchRepositoryMockRecorder) DeleteMatch(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMatch", reflect.TypeOf((*MockIMatchRepository)(nil).DeleteMatch), id)
}

// InsertMatch mocks base method.
func (m *MockIMatchRepository) InsertMatch(communityID string, match *domain.Match) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMatch", communityID, match)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertMatch indicates an expected call of InsertMatch.
func (mr *MockIMatchRepositoryMockRecorder) InsertMatch(communityID, match interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMatch", reflect.TypeOf((*MockIMatchRepository)(nil).InsertMatch), communityID, match)
}

// SelectMatch mocks base method.
func (m *MockIMatchRepository) SelectMatch(criteria domain.MatchCriteria) (domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatch", criteria)
	ret0, _ := ret[0].(domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatch indicates an expected call of SelectMatch.
func (mr *MockIMatchRepositoryMockRecorder) SelectMatch(criteria interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatch", reflect.TypeOf((*MockIMatchRepository)(nil).SelectMatch), criteria)
}

// SelectMatches mocks base method.
func (m *MockIMatchRepository) SelectMatches(criteria domain.MatchCriteria) ([]domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatches", criteria)
	ret0, _ := ret[0].([]domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatches indicates an expected call of SelectMatches.
func (mr *MockIMatchRepositoryMockRecorder) SelectMatches(criteria interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatches", reflect.TypeOf((*MockIMatchRepository)(nil).SelectMatches), criteria)
}

// UpdateMatch mocks base method.
func (m *MockIMatchRepository) UpdateMatch(match *domain.Match) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMatch", match)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMatch indicates an expected call of UpdateMatch.
func (mr *MockIMatchRepositoryMockRecorder) UpdateMatch(match interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMatch", reflect.TypeOf((*MockIMatchRepository)(nil).UpdateMatch), match)
}

// MockIMatchUseCase is a mock of IMatchUseCase interface.
type MockIMatchUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIMatchUseCaseMockRecorder
}

// MockIMatchUseCaseMockRecorder is the mock recorder for MockIMatchUseCase.
type MockIMatchUseCaseMockRecorder struct {
	mock *MockIMatchUseCase
}

// NewMockIMatchUseCase creates a new mock instance.
func NewMockIMatchUseCase(ctrl *gomock.Controller) *MockIMatchUseCase {
	mock := &MockIMatchUseCase{ctrl: ctrl}
	mock.recorder = &MockIMatchUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMatchUseCase) EXPECT() *MockIMatchUseCaseMockRecorder {
	return m.recorder
}

// DeleteMatch mocks base method.
func (m *MockIMatchUseCase) DeleteMatch(ctx *gin.Context, communityID string, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMatch", ctx, communityID, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMatch indicates an expected call of DeleteMatch.
func (mr *MockIMatchUseCaseMockRecorder) DeleteMatch(ctx, communityID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMatch", reflect.TypeOf((*MockIMatchUseCase)(nil).DeleteMatch), ctx, communityID, id)
}

// SelectMatch mocks base method.
func (m *MockIMatchUseCase) SelectMatch(ctx *gin.Context, communityID, id string) (domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatch", ctx, communityID, id)
	ret0, _ := ret[0].(domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatch indicates an expected call of SelectMatch.
func (mr *MockIMatchUseCaseMockRecorder) SelectMatch(ctx, communityID, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatch", reflect.TypeOf((*MockIMatchUseCase)(nil).SelectMatch), ctx, communityID, id)
}

// SelectMatchesByCommunityID mocks base method.
func (m *MockIMatchUseCase) SelectMatchesByCommunityID(ctx *gin.Context, communityId string) ([]domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatchesByCommunityID", ctx, communityId)
	ret0, _ := ret[0].([]domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatchesByCommunityID indicates an expected call of SelectMatchesByCommunityID.
func (mr *MockIMatchUseCaseMockRecorder) SelectMatchesByCommunityID(ctx, communityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatchesByCommunityID", reflect.TypeOf((*MockIMatchUseCase)(nil).SelectMatchesByCommunityID), ctx, communityId)
}

// SelectMatchesByCommunityIDAndDate mocks base method.
func (m *MockIMatchUseCase) SelectMatchesByCommunityIDAndDate(ctx *gin.Context, communityId, date string) ([]domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatchesByCommunityIDAndDate", ctx, communityId, date)
	ret0, _ := ret[0].([]domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatchesByCommunityIDAndDate indicates an expected call of SelectMatchesByCommunityIDAndDate.
func (mr *MockIMatchUseCaseMockRecorder) SelectMatchesByCommunityIDAndDate(ctx, communityId, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatchesByCommunityIDAndDate", reflect.TypeOf((*MockIMatchUseCase)(nil).SelectMatchesByCommunityIDAndDate), ctx, communityId, date)
}

// SelectMatchesByPlayerIDAndDate mocks base method.
func (m *MockIMatchUseCase) SelectMatchesByPlayerIDAndDate(ctx *gin.Context, playerID, date string) ([]domain.Match, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectMatchesByPlayerIDAndDate", ctx, playerID, date)
	ret0, _ := ret[0].([]domain.Match)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectMatchesByPlayerIDAndDate indicates an expected call of SelectMatchesByPlayerIDAndDate.
func (mr *MockIMatchUseCaseMockRecorder) SelectMatchesByPlayerIDAndDate(ctx, playerID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectMatchesByPlayerIDAndDate", reflect.TypeOf((*MockIMatchUseCase)(nil).SelectMatchesByPlayerIDAndDate), ctx, playerID, date)
}

// UpdateMatch mocks base method.
func (m *MockIMatchUseCase) UpdateMatch(ctx *gin.Context, communityID string, id int64, isSingles bool, playerID1, playerID2, playerID3, playerID4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMatch", ctx, communityID, id, isSingles, playerID1, playerID2, playerID3, playerID4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMatch indicates an expected call of UpdateMatch.
func (mr *MockIMatchUseCaseMockRecorder) UpdateMatch(ctx, communityID, id, isSingles, playerID1, playerID2, playerID3, playerID4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMatch", reflect.TypeOf((*MockIMatchUseCase)(nil).UpdateMatch), ctx, communityID, id, isSingles, playerID1, playerID2, playerID3, playerID4)
}
