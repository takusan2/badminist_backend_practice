package repositorytest

import (
	"testing"

	"github.com/takuya-okada-01/badminist-backend/domain"
)

var InsertCommunity = domain.Community{
	Name:        "testCommunity",
	Description: "testDescription",
}

func GenInsertCommunityTestFunc(cr domain.ICommunityRepository) func(t *testing.T) {
	return func(t *testing.T) {
		id, err := cr.InsertCommunity(&InsertCommunity)
		if err != nil {
			t.Fatal(err)
		}
		want := id
		community, err := cr.SelectCommunity(id)
		if err != nil {
			t.Fatal(err)
		}
		if community.ID != want {
			t.Errorf("InsertCommunity == %s, want %s", community.ID, want)
		}
		// 更新
		InsertCommunity.ID = community.ID
	}
}

func GenDeleteCommunityTestFunc(cr domain.ICommunityRepository) func(t *testing.T) {
	return func(t *testing.T) {
		err := cr.DeleteCommunity(InsertCommunity.ID)
		if err != nil {
			t.Fatal(err)
		}
		_, err = cr.SelectCommunity(InsertCommunity.ID)
		if err == nil {
			t.Fatal(err)
		}
	}
}
