package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

func GenInsertPlayerTestFunc(pr domain.IPlayerRepository) func(t *testing.T) {
	return func(t *testing.T) {

		insertPlayer := &domain.Player{
			Name:        "test",
			Sex:         true,
			Age:         20,
			Level:       10,
			CommunityID: InsertCommunity.ID,
		}

		id, err := pr.InsertPlayer(insertPlayer)
		if err != nil {
			t.Fatal(err)
		}
		want := id
		player, err := pr.SelectPlayer(id)
		if err != nil {
			t.Fatal(err)
		}
		if player.ID != want {
			t.Errorf("InsertPlayer == %s, want %s", player.ID, want)
		}
	}
}

func GenDeletePlayerTestFunc(pr domain.IPlayerRepository) func(t *testing.T) {
	return func(t *testing.T) {
		insertPlayer := &domain.Player{
			Name:        "test",
			Sex:         true,
			Age:         20,
			Level:       10,
			CommunityID: InsertCommunity.ID,
		}

		err := pr.DeletePlayer(insertPlayer.ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = pr.SelectPlayer(insertPlayer.ID)
		if err == nil {
			t.Fatal(err)
		}
	}
}
