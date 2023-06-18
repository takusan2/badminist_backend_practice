package repository

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
)

func TestMatchRepository(t *testing.T) {
	db := database.Connect()
	matchRepository := NewMatchRepository(db)
	insertMatch := domain.Match{
		IsSingles: true,
		PlayerID1: "test1",
		PlayerID2: "test2",
	}
	id, err := matchRepository.InsertMatch("test", &insertMatch)
	if err != nil {
		t.Fatal(err)
	}

	match, err := matchRepository.SelectMatch(id)
	if err != nil {
		t.Fatal(err)
	}
	if match.PlayerID1 != "test1" {
		t.Errorf("InsertMatch == %s, want %s", match.PlayerID1, "test1")
	}
	if match.PlayerID2 != "test2" {
		t.Errorf("InsertMatch == %s, want %s", match.PlayerID2, "test2")
	}
	if match.IsSingles != true {
		t.Errorf("InsertMatch == %t, want %t", match.IsSingles, true)
	}

	match.PlayerID1 = "test3"
	matchRepository.UpdateMatch(&match)
	updatedMatch, err := matchRepository.SelectMatch(id)
	if err != nil {
		t.Fatal(err)
	}
	if updatedMatch.PlayerID1 != "test3" {
		t.Errorf("UpdateMatch == %s, want %s", updatedMatch.PlayerID1, "test3")
	}

	matches, err := matchRepository.SelectMatchesByCommunityID("test_community")
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Errorf("SelectMatchesByCommunityID == %d, want %d", len(matches), 1)
	}

	matchRepository.DeleteMatch(id)
}
