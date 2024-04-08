package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultDbPort = 3306
	// List of MySQL error codes https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	MySQLErrorCodeDuplicateEntry = 1062 // ER_DUP_ENTRY - Means that a duplicate value has been inserted into a column that has a UNIQUE constraint.
)

var db *sql.DB

type MysqlConf struct {
	DbName     string // The database that you want to access
	DbUser     string // The database account that you want to access
	DbPassword string // The password for the database account you want to access
	DbHost     string // The endpoint of the DB instance that you want to access
	DbPort     int    // The port number used for connecting to your DB instance
}

func NewMySqlConf(user, password string) MysqlConf {
	fmt.Println("esto ", os.Getenv("DBPORT"))
	//valueStr := os.Getenv("DBPORT")
	/*value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("error al convertir a entero")
	}*/

	return MysqlConf{
		DbName:     "exampledb",
		DbUser:     "root",
		DbPassword: "developer",
		DbHost:     "mysqldb-example",
		DbPort:     3308,
	}
}

func (mysql MysqlConf) InitMySqlDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("fail open")
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println("fail ping no gusta")
		return nil, err
	}

	//db = conn
	return conn, nil
}

// GetDB devuelve la conexión a la base de datos
func (mysql MysqlConf) GetDB() *sql.DB {
	return db
}
