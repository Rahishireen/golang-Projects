package users_db

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*func init() {
	//usersDB, err := sql.Open()
	//if err
}

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host  ="mysql_users_host"
	mysql_users_schema ="mysql_users_schema"
)

var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)

)*/

var (
	Client *sql.DB
)

func init(){
	//dataSourceName:=fmt.Sprintf(":%s@tcp(%s)/%s?charset=utf8","root","rahima","127.0.0.1:8080","users_db")

	var err error
	Client,err =sql.Open("mysql","root:rahima@tcp(127.0.0.1:3306)/users_db")
	if err!=nil{
		panic(err)
	}

	if err = Client.Ping();err!=nil{
		panic(err)
	}
	log.Println(("database Sucessfully Configured"))

}