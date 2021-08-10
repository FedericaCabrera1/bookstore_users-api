package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "MYSQL_USERS_USERNAME"
	mysql_users_password = "MYSQL_USERS_PASSWORD"
	mysql_users_host     = "MYSQL_USERS_HOST"
	mysql_users_schema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB //we create the Client in the DB

	username = os.Getenv(mysql_users_username) //attempt to get the username from an enviroment variable (already set from the terminal) called mysql_users_username
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() { //we are creating the connection in the init function
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName) //Open opens a DB specified by its name and a specific data source name
	if err != nil {
		panic(err) //if there´s ant error, we throw a panic, we are not going to start the application as we need the db in order to work
	}
	//if we reach this point means that we have a valid DB
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("database successfully configured")

}
