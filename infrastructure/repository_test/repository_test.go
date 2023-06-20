package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/infrastructure/database"
	community_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/community"
	owner_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/owner"
	player_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/player"
	user_repository "github.com/takuya-okada-01/badminist-backend/infrastructure/repository/user"
)

func TestRepository(t *testing.T) {
	db := database.Connect()
	ur := user_repository.NewUserRepository(db)
	cr := community_repository.NewCommunityRepository(db)
	or := owner_repository.NewOwnerRepository(db)
	pr := player_repository.NewPlayerRepository(db)
	t.Run("UserInsertAndSelect", GenInsertUserTestFunc(ur))
	t.Run("CommunityInsertAndSelect", GenInsertCommunityTestFunc(cr))
	t.Run("OwnerInsertAndSelect", GenInsertOwnerTestFunc(or))
	t.Run("PlayerInsertAndSelect", GenInsertPlayerTestFunc(pr))

}
