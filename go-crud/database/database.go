/*
For you to fix this error you have to install the mySQL driver
	using cmd with the command: go get -u gorm.io/driver/mysql
*/

package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/* func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database.")
	}
}


func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	LoadEnvVariables()
	ConnectToDB()
	// Migrate()
}
*/

//New lines of code down here my gee

func LoadEnvVariables() {
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "#D-Wayne_Atlanta")
	os.Setenv("DB_NAME", "student_loan")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
}

func ConnectToDB(){

	LoadEnvVariables()

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error 
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the Database my guy: %v", err)
	}
	log.Println("Database connection established")

}
