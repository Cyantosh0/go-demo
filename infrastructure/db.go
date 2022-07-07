package infrastructure

import (
	"fmt"
	"github/Cyantosh0/demo/lib"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	dsn string
}

func NewDatabase(env *lib.Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println("Status:", err)
	}

	return Database{
		DB:  db,
		dsn: url,
	}
}
