package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"strconv"
)

const (
	defaultDbPort = 3306
	// List of MySQL error codes https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	MySQLErrorCodeDuplicateEntry = 1062 // ER_DUP_ENTRY - Means that a duplicate value has been inserted into a column that has a UNIQUE constraint.
)

type MySQLConf struct {
	DbName     string // The database that you want to access
	DbUser     string // The database account that you want to access
	DbPassword string // The password for the database account you want to access
	DbHost     string // The endpoint of the DB instance that you want to access
	DbPort     int    // The port number used for connecting to your DB instance
}

func NewMySqlConf(user, password string) MySQLConf {
	valueStr := os.Getenv("DBPORT")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("error al convertir a entero")
	}

	return MySQLConf{
		DbName:     os.Getenv("DBNAME"),
		DbUser:     os.Getenv("MYSQLUSER"),
		DbPassword: os.Getenv("MYSQLPASSWORD"),
		DbHost:     os.Getenv("DBHOST"),
		DbPort:     value,
	}
}

func (mysql MySQLConf) InitMySqlDB(config MySQLConf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&clientFoundRows=true",
		config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
