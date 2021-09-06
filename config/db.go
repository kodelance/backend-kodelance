package config

import (
	"kodelance/user"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_pass := os.Getenv("DB_PASS")
	db_user := os.Getenv("DB_USER")
	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	return DB, nil
}
