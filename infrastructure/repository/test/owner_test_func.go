package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

func GenInsertOwnerTestFunc(or domain.IOwnerRepository) func(t *testing.T) {
	return func(t *testing.T) {
		var insertOwner = domain.Owner{
			UserID:      InsertUser.ID,
			CommunityID: InsertCommunity.ID,
			Role:        domain.Admin.String(),
		}

		err := or.InsertOwner(insertOwner)
		if err != nil {
			t.Fatal(err)
		}
		owner, err := or.SelectOwnerByUserIDAndCommunityID(
			insertOwner.UserID,
			insertOwner.CommunityID,
		)
		if err != nil {
			t.Fatal(err)
		}
		if owner.UserID != insertOwner.UserID {
			t.Errorf("InsertOwner == %s, want %s", owner.UserID, insertOwner.UserID)
		}
		if owner.CommunityID != insertOwner.CommunityID {
			t.Errorf("InsertOwner == %s, want %s", owner.CommunityID, insertOwner.CommunityID)
		}
		if owner.Role != insertOwner.Role {
			t.Errorf("InsertOwner == %s, want %s", owner.Role, insertOwner.Role)
		}
	}
}

func GenDeleteOwnerTestFunc(or domain.IOwnerRepository) func(t *testing.T) {
	return func(t *testing.T) {
		var insertOwner = domain.Owner{
			UserID:      InsertUser.ID,
			CommunityID: InsertCommunity.ID,
			Role:        domain.Admin.String(),
		}

		err := or.DeleteOwner(insertOwner.UserID, insertOwner.CommunityID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = or.SelectOwnerByUserIDAndCommunityID(
			insertOwner.UserID,
			insertOwner.CommunityID,
		)
		if err == nil {
			t.Fatal(err)
		}
	}
}
