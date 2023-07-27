package usecase_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

func Test_playerUseCase_SelectPlayersByCommunityID(t *testing.T) {
	type fields struct {
		playerRepository domain.IPlayerRepository
	}
	type args struct {
		ctx         *gin.Context
		communityID string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(mpr mock_domain.MockIPlayerRepository)
		want          []domain.Player
		wantErr       bool
	}{
		{
			name: "Not found the communityID",
			args: args{
				ctx:         &gin.Context{},
				communityID: "testCommunityID",
			},
			prepareMockFn: func(mpr mock_domain.MockIPlayerRepository) {
				mpr.EXPECT().SelectPlayersByCommunityID("testCommunity").
					Return([]domain.Player{}, fmt.Errorf("Not found the communityID"))
			},
			want:    []domain.Player{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mpr := mock_domain.NewMockIPlayerRepository(mockCtrl)
			mor := mock_domain.NewMockIOwnerRepository(mockCtrl)
			puc := usecase.NewPlayerUseCase(mpr, mor)

			if tt.prepareMockFn != nil {
				tt.prepareMockFn(*mpr)
			}

			got, err := puc.SelectPlayersByCommunityID(
				tt.args.communityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("playerUseCase.SelectPlayersByCommunityID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playerUseCase.SelectPlayersByCommunityID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_playerUseCase_SelectPlayersByCommunityIDAndAttendance(t *testing.T) {
	type fields struct {
		playerRepository domain.IPlayerRepository
	}
	type args struct {
		ctx         *gin.Context
		communityID string
		attendance  bool
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(mpr mock_domain.MockIPlayerRepository)
		want          []domain.Player
		wantErr       bool
	}{
		{
			name: "Not found the communityID",
			args: args{
				ctx:         &gin.Context{},
				communityID: "testCommunityID",
				attendance:  true,
			},
			prepareMockFn: func(mpr mock_domain.MockIPlayerRepository) {
				mpr.EXPECT().SelectPlayersByCommunityID("testCommunityID").
					Return([]domain.Player{}, fmt.Errorf("Not found the communityID"))
			},
			want:    []domain.Player{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mpr := mock_domain.NewMockIPlayerRepository(mockCtrl)
			mor := mock_domain.NewMockIOwnerRepository(mockCtrl)
			upc := usecase.NewPlayerUseCase(mpr, mor)

			if tt.prepareMockFn != nil {
				tt.prepareMockFn(*mpr)
			}

			got, err := upc.SelectAttendPlayers(tt.args.communityID)
			if (err != nil) != tt.wantErr {
				t.Errorf("playerUseCase.SelectPlayersByCommunityIDAndAttendance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("playerUseCase.SelectPlayersByCommunityIDAndAttendance() = %v, want %v", got, tt.want)
			}
		})
	}
}
