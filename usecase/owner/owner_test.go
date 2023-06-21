package usecase

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/takuya-okada-01/badminist-backend/domain"
	mock_domain "github.com/takuya-okada-01/badminist-backend/mock"
)

func Test_ownerUseCase_UpdateOwner(t *testing.T) {
	type fields struct {
		ownerRepository domain.IOwnerRepository
	}
	type args struct {
		ctx    *gin.Context
		userID string
		owner  *domain.Owner
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(mor mock_domain.MockIOwnerRepository)
		wantErr       bool
	}{
		{
			name: "Not authorized to update",
			args: args{
				ctx:    &gin.Context{},
				userID: "testUserID",
				owner: &domain.Owner{
					UserID:      "testOwnerUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Staff.String(),
				},
			},
			prepareMockFn: func(mor mock_domain.MockIOwnerRepository) {
				mor.EXPECT().SelectOwner(domain.OwnerCriteria{
					UserID:               "testUserID",
					UserIDIsNotNull:      true,
					CommunityID:          "testCommunityID",
					CommunityIDIsNotNull: true,
				}).
					Return(domain.Owner{
						UserID:      "testOwnerUserID",
						CommunityID: "testCommunityID",
						Role:        domain.Member.String(),
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "Not found the userID in the community",
			args: args{
				ctx:    &gin.Context{},
				userID: "testUserID",
				owner: &domain.Owner{
					UserID:      "testOwnerUserID",
					CommunityID: "testCommunityID",
					Role:        domain.Staff.String(),
				},
			},
			prepareMockFn: func(mor mock_domain.MockIOwnerRepository) {
				mor.EXPECT().SelectOwner(domain.OwnerCriteria{
					UserID:               "testUserID",
					UserIDIsNotNull:      true,
					CommunityID:          "testCommunityID",
					CommunityIDIsNotNull: true,
				}).
					Return(domain.Owner{}, fmt.Errorf("not found the userID in the community"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mor := mock_domain.NewMockIOwnerRepository(mockCtrl)
			ouc := &ownerUseCase{
				ownerRepository: mor,
			}
			if tt.prepareMockFn != nil {
				tt.prepareMockFn(*mor)
			}

			if err := ouc.UpdateOwner(tt.args.ctx, tt.args.userID, tt.args.owner); (err != nil) != tt.wantErr {
				t.Errorf("ownerUseCase.UpdateOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ownerUseCase_DeleteOwner(t *testing.T) {
	type fields struct {
		ownerRepository domain.IOwnerRepository
	}
	type args struct {
		ctx            *gin.Context
		userID         string
		delUserID      string
		delCommunityID string
	}
	tests := []struct {
		name          string
		args          args
		prepareMockFn func(mor mock_domain.MockIOwnerRepository)
		wantErr       bool
	}{
		{
			name: "Not authorized to delete",
			args: args{
				ctx:            &gin.Context{},
				userID:         "testUserID",
				delUserID:      "testDelUserID",
				delCommunityID: "testCommunityID",
			},
			prepareMockFn: func(mor mock_domain.MockIOwnerRepository) {
				mor.EXPECT().
					SelectOwner(domain.OwnerCriteria{
						UserID:               "testUserID",
						UserIDIsNotNull:      true,
						CommunityID:          "testCommunityID",
						CommunityIDIsNotNull: true,
					}).
					Return(domain.Owner{
						UserID:      "testOwnerUserID",
						CommunityID: "testCommunityID",
						Role:        domain.Member.String(),
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "Not found the userID in the community",
			args: args{
				ctx:            &gin.Context{},
				userID:         "testUserID",
				delUserID:      "testDelUserID",
				delCommunityID: "testDelCommunityID",
			},
			prepareMockFn: func(mor mock_domain.MockIOwnerRepository) {
				mor.EXPECT().
					SelectOwner(domain.OwnerCriteria{
						UserID:               "testUserID",
						UserIDIsNotNull:      true,
						CommunityID:          "testDelCommunityID",
						CommunityIDIsNotNull: true,
					}).
					Return(domain.Owner{}, fmt.Errorf("not found the userID in the community"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mor := mock_domain.NewMockIOwnerRepository(mockCtrl)
			ouc := &ownerUseCase{
				ownerRepository: mor,
			}
			if tt.prepareMockFn != nil {
				tt.prepareMockFn(*mor)
			}

			if err := ouc.DeleteOwner(tt.args.ctx, tt.args.userID, tt.args.delUserID, tt.args.delCommunityID); (err != nil) != tt.wantErr {
				t.Errorf("ownerUseCase.DeleteOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
