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

	//Auto create table based on Model
	user := &UserModel{Name: "John", Address: "New York"}
	dbConnection.Create(user)

	// Internally it will create the query like
	// INSERT INTO 'user_models' ('name','address') VALUES ('John','New York')

	//You can insert multiple records too
	var users = []UserModel{
		UserModel{Name: "Ricky", Address: "Sydney"},
		UserModel{Name: "Adam", Address: "Brisbane"},
		UserModel{Name: "Justin", Address: "California"},
	}

	for _, user := range users {
		dbConnection.Create(&user)
	}
	usernew := &UserModel{Name: "John", Address: "New York"}
	// Select, edit, and save
	dbConnection.Find(&usernew)
	usernew.Address = "Brisbane"
	dbConnection.Save(&usernew)

	// Update with column names, not attribute names
	dbConnection.Model(&usernew).Update("Name", "Jack")

	dbConnection.Model(&usernew).Updates(
		map[string]interface{}{
			"Name":    "Amy",
			"Address": "Boston",
		})

	// UpdateColumn()
	dbConnection.Model(&usernew).UpdateColumn("Address", "Phoenix")
	dbConnection.Model(&usernew).UpdateColumns(
		map[string]interface{}{
			"Name":    "Taylor",
			"Address": "Houston",
		})
	// Using Find()
	dbConnection.Find(&usernew).Update("Address", "San Diego")

	// Batch Update
	dbConnection.Table("user_models").Where("address = ?", "california").Update("name", "Walker")
}
