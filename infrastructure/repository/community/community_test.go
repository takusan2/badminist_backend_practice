package repository

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
)

func TestCommunityRepository(t *testing.T) {
	db := database.Connect()
	communityRepository := NewCommunityRepository(db)
	insertCommunity := domain.Community{
		Name:        "test",
		Description: "test",
	}

	id, err := communityRepository.InsertCommunity(&insertCommunity)
	if err != nil {
		t.Fatal(err)
	}

	want := id
	community, err := communityRepository.SelectCommunity(id)
	if err != nil {
		t.Fatal(err)
	}
	if community.ID != want {
		t.Errorf("InsertCommunity == %s, want %s", community.ID, want)
	}

	community.Name = "test_updated"
	communityRepository.UpdateCommunity(&community)
	community, err = communityRepository.SelectCommunity(id)
	if err != nil {
		t.Fatal(err)
	}
	if community.Name != "test_updated" {
		t.Errorf("UpdateCommunity == %s, want %s", community.Name, "test_updated")
	}

	// communityRepository.DeleteCommunity(id)
}
