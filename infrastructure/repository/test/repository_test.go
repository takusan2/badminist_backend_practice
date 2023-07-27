package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
	"github.com/takuya-okada-01/badminist-backend/infrastructure/repository"
)

func TestRepository(t *testing.T) {
	db := database.Connect()
	ur := repository.NewUserRepository(db)
	cr := repository.NewCommunityRepository(db)
	or := repository.NewOwnerRepository(db)
	pr := repository.NewPlayerRepository(db)
	t.Run("UserInsertAndSelect", GenInsertUserTestFunc(ur))
	t.Run("CommunityInsertAndSelect", GenInsertCommunityTestFunc(cr))
	t.Run("OwnerInsertAndSelect", GenInsertOwnerTestFunc(or))
	t.Run("PlayerInsertAndSelect", GenInsertPlayerTestFunc(pr))

}
