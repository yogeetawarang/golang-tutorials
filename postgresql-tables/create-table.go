package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// UserModel - comments
type UserModel struct {
	ID      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)”`
}

// Model definition given by gorm
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User - comments
type User struct {
	gorm.Model // fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`will be added
	Name       string
}

func main() {

	server := "127.0.0.1"
	port := 5432
	username := "postgres"
	database := "customerdb"
	password := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		server, port, username, database,
		password)
	dbConnection, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("Connection Established")
	// dbConnection.Debug().AutoMigrate(&User{}) //Model or Struct
	dbConnection.Debug().DropTableIfExists(&UserModel{})
	//Drops table if already exists
	dbConnection.Debug().AutoMigrate(&UserModel{})
	//Auto create table based on Model
}
