package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq" //import postgres driver
)

// DB db object
var DB *sql.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int64
	User     string
	DBName   string
	DBName2  string
	Password string
}

// BuildDBConfig builds db config object from environment variables
func BuildDBConfig() DBConfig {

	port, _ := strconv.ParseInt(os.Getenv("DBPORT"), 10, 0)

	dbConfig := DBConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     port,
		User:     os.Getenv("DBUSER"),
		DBName:   os.Getenv("DBNAME"),
		Password: os.Getenv("DBPASS"),
	}
	return dbConfig
}

// DbURL get db connection string
func (dbConfig DBConfig) DbURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)
}

// GetDB2 same as getDB but returns an error when it cannot connect.
func GetDB2() (*sql.DB, error) {
	DB, err := sql.Open("postgres", BuildDBConfig().DbURL())
	if err != nil {
		log.Println("GetDB2: Failed to connect to DB: " + err.Error())
		return DB, err
	}

	return DB, nil
}
