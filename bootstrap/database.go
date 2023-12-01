package bootstrap

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewMySQLDatabase(env *Env) *Client {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/clean?charset=utf8mb4&parseTime=True&loc=Local", env.User, env.Password, env.Host, env.Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error opening MySQL connection: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error getting database instance: ", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	client := &Client{
		DB: db,
	}

	return client
}

func CloseMySQLConnection(client *Client) error {
	sqlDB, err := client.DB.DB()
	if err != nil {
		return fmt.Errorf("error getting database instance: %v", err)
	}

	return sqlDB.Close()
}