package usecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
	"github.com/takuya-okada-01/badminist-backend/usecase"
)

func Test_communityUseCase_UpdateCommunity(t *testing.T) {
	type fields struct {
		communityRepository domain.ICommunityRepository
	}
	type args struct {
		userID    string
		community domain.Community
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
				mor.EXPECT().SelectOwnerByUserIDAndCommunityID(
					"testUserID",
					"testCommunityID",
				).Return(domain.Owner{
					UserID:      "testUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Admin.String(),
				}, nil)
				mcr.EXPECT().
					SelectCommunityByID("testCommunityID").
					Return(domain.Community{}, fmt.Errorf("not exists community"))
			},
			args{
				userID: "testUserID",
				community: domain.Community{
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
				mor.EXPECT().
					SelectOwnerByUserIDAndCommunityID(
						"testUserID",
						"testCommunityID",
					).
					Return(domain.Owner{
						UserID:      "testUserID",
						CommunityID: "testCommunityID",
						Role:        domain.Member.String(),
					}, nil)
			},
			args{
				userID: "testUserID",
				community: domain.Community{
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
				mor.EXPECT().SelectOwnerByUserIDAndCommunityID(
					"testUserID",
					"testCommunityID",
				).Return(domain.Owner{
					UserID:      "testUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Admin.String(),
				}, nil)
				mcr.EXPECT().
					SelectCommunityByID(
						"testCommunityID",
					).
					Return(
						domain.Community{
							ID:          "testCommunityID",
							Name:        "testName",
							Description: "testDescription",
						},
						nil)
				mcr.EXPECT().
					UpdateCommunity(&domain.Community{
						ID:          "testCommunityID",
						Name:        "updateTestName",
						Description: "updateTestDescription",
					}).
					Return(
						nil,
					)
			},
			args{
				userID: "testUserID",
				community: domain.Community{
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
			cu := usecase.NewCommunityUseCase(mcr, mor)
			resCommunity, err := cu.UpdateCommunity(tt.args.userID, tt.args.community)
			if (err != nil) != tt.wantErr {
				t.Errorf("communityUseCase.UpdateCommunity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if resCommunity.ID != tt.args.community.ID {
				t.Errorf("communityUseCase.UpdateCommunity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if resCommunity.Name != tt.args.community.Name {
				t.Errorf("communityUseCase.UpdateCommunity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if resCommunity.Description != tt.args.community.Description {
				t.Errorf("communityUseCase.UpdateCommunity() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
