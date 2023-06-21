package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

var InsertUser = domain.User{
	Name:         "test",
	Email:        "hoge@hoge.com",
	PasswordHash: "password",
}

func GenInsertUserTestFunc(ur domain.IUserRepository) func(t *testing.T) {
	return func(t *testing.T) {
		id, err := ur.InsertUser(&InsertUser)
		if err != nil {
			t.Fatal(err)
		}
		want := id
		user, err := ur.SelectUser(domain.UserCriteria{
			ID:          id,
			IDIsNotNull: true,
		})
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != want {
			t.Errorf("InsertUser == %s, want %s", user.ID, want)
		}
		// 更新
		InsertUser.ID = user.ID
	}
}

func GenDeleteUserTestFunc(ur domain.IUserRepository) func(t *testing.T) {
	return func(t *testing.T) {
		err := ur.DeleteUser(InsertUser.ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = ur.SelectUser(
			domain.UserCriteria{
				ID:          InsertUser.ID,
				IDIsNotNull: true,
			},
		)
		if err == nil {
			t.Fatal(err)
		}
	}
}
