package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"Hackathon/model/DTO"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("mysql_username"),
		os.Getenv("mysql_pwd"),
		os.Getenv("mysql_address"),
		os.Getenv("mysql_port"),
		os.Getenv("mysql_db_name"))

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DefaultDatetimePrecision := 0
	database, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB, DefaultDatetimePrecision: &DefaultDatetimePrecision}),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal(err)
	}

	if err = database.AutoMigrate(
		&DTO.User{},
		&DTO.Post{},
		&DTO.PostLike{},
		&DTO.PostSticker{},
		&DTO.Comment{},
		&DTO.CommentLike{},
		&DTO.CommentSticker{},
		&DTO.CommentEmoji{},
	); err != nil {
		log.Fatal(err)
	}
	DB = database

	log.Print("[DATABASE] 연결 완료")
}
