package repository

import (
	"fmt"
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
)

func TestPlayerRepository(t *testing.T) {
	db := database.Connect()
	playerRepository := NewPlayerRepository(db)
	insertPlayer := domain.Player{
		Name: "test",
		Sex:  true,
		Age:  20,
	}

	id, err := playerRepository.InsertPlayer("test", &insertPlayer)
	if err != nil {
		t.Fatal(err)
	}

	want := id
	player, err := playerRepository.SelectPlayer(id)
	fmt.Printf("%+v", player)
	if err != nil {
		t.Fatal(err)
	}

	if player.ID != want {
		t.Errorf("InsertPlayer == %s, want %s", player.ID, want)
	}
	player.Age = 30
	player.Attendance = true
	playerRepository.UpdatePlayer(&player)
	updatedPlayer, err := playerRepository.SelectPlayer(id)
	if updatedPlayer.Age != 30 {
		t.Errorf("UpdatePlayer == %d, want %d", updatedPlayer.Age, 30)
	}
	if updatedPlayer.Attendance != true {
		t.Errorf("UpdatePlayer == %t, want %t", updatedPlayer.Attendance, true)
	}

	playerRepository.DeletePlayer(id)
}
