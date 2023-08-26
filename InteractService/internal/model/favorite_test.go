package model

import (
	"fmt"
	"testing"
)

func TestCreateShardingTables(t *testing.T) {
	InitDB()
	for i := 0; i < 3; i++ {
		tableName := fmt.Sprintf("favorite_%d", i) //表名
		err := DB.Table(tableName).AutoMigrate(&Favorite{})
		if err != nil {
			t.Error(err)
		}
	}
}

func TestUpdateFavoriteStatusAction(t *testing.T) {
	InitDB()
	err := UpdateVideoLikedStatus(1, 1, false)
	if err != nil {
		t.Error(err)
	}
}
