package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// if supabase database not connected please use local database
	dbuser := "postgres"
	dbpass := "12345"
	dbhost := "localhost"
	dbname := "rental-car"
	dbport := "5432"
	/*
		dbuser := "postgres.rwxyxvcywxpoqiovconl"
		dbpass := "livecode3-phase2-v2"
		dbhost := "aws-0-ap-southeast-1.pooler.supabase.com"
		dbname := "postgres"
		dbport := "6543"
	*/

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbhost, dbuser, dbpass, dbname, dbport)

	// dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=require&PreferSimpleProtocol=true", dbuser, dbpass, dbhost, dbport, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	// set max idle connection
	sqlDB, err := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	log.Println("Connection to database closed")
}
