package usecase

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
)

func Test_userUseCase_UpdateUser(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx  *gin.Context
		id   string
		name string
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock_domain.MockIUserRepository)
		args          args
		wantErr       bool
	}{
		{
			name: "not exists user",
			prepareMockFn: func(m *mock_domain.MockIUserRepository) {
				m.EXPECT().
					SelectUser(domain.UserCriteria{ID: "testID", IDIsNotNull: true}).
					Return(domain.User{}, fmt.Errorf("not exists user"))
			},
			args: args{
				ctx:  &gin.Context{},
				id:   "testID",
				name: "testName",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ur := mock_domain.NewMockIUserRepository(mockCtrl)
			tt.prepareMockFn(ur)
			u := NewUserUseCase(ur)

			if err := u.UpdateUser(tt.args.ctx, tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
