package usecase

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
)

func Test_communityUseCase_UpdateCommunity(t *testing.T) {
	type fields struct {
		communityRepository domain.ICommunityRepository
	}
	type args struct {
		ctx       *gin.Context
		userID    string
		community *domain.Community
	}
	tests := []struct {
		name          string
		prepareMockFn func(mcr *mock_domain.MockICommunityRepository, mor *mock_domain.MockIOwnerRepository)
		args          args
		wantErr       bool
	}{
		{
			"not exists community",
			func(mcr *mock_domain.MockICommunityRepository, mor *mock_domain.MockIOwnerRepository) {
				mor.EXPECT().SelectOwnerByUserIDAndCommunityID("testUserID", "testCommunityID").Return(domain.Owner{
					UserID:      "testUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Admin.String(),
				}, nil)
				mcr.EXPECT().SelectCommunity("testCommunityID").Return(domain.Community{}, fmt.Errorf("not exists community"))
			},
			args{
				ctx:    &gin.Context{},
				userID: "testUserID",
				community: &domain.Community{
					ID:          "testCommunityID",
					Name:        "testName",
					Description: "testDescription",
				},
			},
			true,
		},
		{
			"member is no authorized to update",
			func(mcr *mock_domain.MockICommunityRepository, mor *mock_domain.MockIOwnerRepository) {
				mor.EXPECT().SelectOwnerByUserIDAndCommunityID("testUserID", "testCommunityID").Return(domain.Owner{
					UserID:      "testUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Member.String(),
				}, nil)
			},
			args{
				ctx:    &gin.Context{},
				userID: "testUserID",
				community: &domain.Community{
					ID:          "testCommunityID",
					Name:        "testName",
					Description: "testDescription",
				},
			},
			true,
		},
		{
			"success",
			func(mcr *mock_domain.MockICommunityRepository, mor *mock_domain.MockIOwnerRepository) {
				mor.EXPECT().SelectOwnerByUserIDAndCommunityID("testUserID", "testCommunityID").Return(domain.Owner{
					UserID:      "testUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Admin.String(),
				}, nil)
				mcr.EXPECT().SelectCommunity("testCommunityID").DoAndReturn(
					func(id string) (domain.Community, error) {
						return domain.Community{
							ID:          "testCommunityID",
							Name:        "testName",
							Description: "testDescription",
						}, nil
					},
				)
				mcr.EXPECT().UpdateCommunity(&domain.Community{
					ID:          "testCommunityID",
					Name:        "updateTestName",
					Description: "updateTestDescription",
				}).DoAndReturn(
					func(community *domain.Community) error {
						return nil
					},
				)
			},
			args{
				ctx:    &gin.Context{},
				userID: "testUserID",
				community: &domain.Community{
					ID:          "testCommunityID",
					Name:        "updateTestName",
					Description: "updateTestDescription",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mcr := mock_domain.NewMockICommunityRepository(mockCtrl)
			mor := mock_domain.NewMockIOwnerRepository(mockCtrl)
			tt.prepareMockFn(mcr, mor)
			cu := NewCommunityUseCase(mcr, mor)
			if err := cu.UpdateCommunity(tt.args.ctx, tt.args.userID, tt.args.community); (err != nil) != tt.wantErr {
				t.Errorf("communityUseCase.UpdateCommunity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
