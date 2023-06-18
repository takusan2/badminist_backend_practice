package repository

import (
	"fmt"
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
	community_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/community"
	user_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/user"
)

func TestOwnerRepository(t *testing.T) {
	db := database.Connect()
	ownerRepository := NewOwnerRepository(db)
	userRepository := user_repository.NewUserRepository(db)
	communityRepository := community_repository.NewCommunityRepository(db)

	userID, err := userRepository.InsertUser(&domain.User{
		Email:        "hoge@hoge.com",
		PasswordHash: "password",
	})

	communityID, err := communityRepository.InsertCommunity(&domain.Community{
		Name:        "test",
		Description: "test",
	})

	insertOwner := domain.Owner{
		UserID:      userID,
		CommunityID: communityID,
		Role:        domain.Admin.String(),
	}
	fmt.Print(insertOwner)

	err = ownerRepository.InsertOwner(&insertOwner)
	if err != nil {
		t.Fatal(err)
	}

	owner, err := ownerRepository.SelectOwner(insertOwner)
	if err != nil {
		t.Fatal(err)
	}
	if owner.UserID != userID {
		t.Errorf("InsertOwner == %s, want %s", owner.UserID, userID)
	}
	if owner.CommunityID != communityID {
		t.Errorf("InsertOwner == %s, want %s", owner.CommunityID, communityID)
	}

	owner.Role = domain.Member.String()
	ownerRepository.UpdateOwner(&owner)
	updatedOwner, err := ownerRepository.SelectOwner(owner)
	if err != nil {
		t.Fatal(err)
	}
	if updatedOwner.Role != domain.Member.String() {
		t.Errorf("UpdateOwner == %s, want %s", updatedOwner.Role, domain.Member.String())
	}

	ownerRepository.DeleteOwner(owner)

}
