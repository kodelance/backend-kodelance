package config

import (
	"fmt"
	"kodelance/user"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	// dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		"5432",
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	return DB, nil
}
