package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/sharding"
)

var (
	dsn = "root:Lhj000922!@tcp(106.54.208.133:3306)/pinktok?charset=utf8&parseTime=True&loc=Local"
	DB  *gorm.DB
	err error
)

func InitDB() {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err = DB.Use(sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      3,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "favorite"))
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&Favorite{})
	if err != nil {
		panic(err)
	}
}
