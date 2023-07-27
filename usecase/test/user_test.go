package usecase_test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
	"github.com/takuya-okada-01/badminist-backend/usecase"
	"github.com/takuya-okada-01/badminist-backend/validator"
)

func Test_userUseCase_UpdateUser(t *testing.T) {
	type fields struct {
		userRepository domain.IUserRepository
	}
	type args struct {
		ctx  *gin.Context
		id   string
		user domain.User
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
					SelectUser("testID").
					Return(domain.User{}, fmt.Errorf("not exists user"))
			},
			args: args{
				id:   "testID",
				user: domain.User{Name: "testName", Email: "testEmail"},
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
			uv := validator.NewUserValidator()
			u := usecase.NewUserUseCase(ur, uv)
			resUser, err := u.UpdateUser(tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUseCase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if resUser.Name != tt.args.user.Name {
				t.Errorf("userUseCase.UpdateUser() name = %v, want %v", resUser.Name, tt.args.user.Name)
			}
			if resUser.Email != tt.args.user.Email {
				t.Errorf("userUseCase.UpdateUser() email = %v, want %v", resUser.Email, tt.args.user.Email)
			}
		})
	}
}
