package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysqlUsersUserName = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHostname = "mysql_users_hostname"
	mysqlUsersSchema   = "mysql_users_schemaname"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv(mysqlUsersUserName)
	password := os.Getenv(mysqlUsersPassword)
	host := os.Getenv(mysqlUsersHostname)
	schema := os.Getenv(mysqlUsersSchema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully")

}
