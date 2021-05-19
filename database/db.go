package database

import (
	"database/sql"
	"demo_api/models"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type Connection interface {
	Close()
	DB() *sql.DB
}

type conn struct {
	session *sql.DB
}

func (c *conn) DB() *sql.DB {
	var err error
	c.session, err = sql.Open("mysql", connstr())
	if err != nil {
		panic(err)
	}
	return c.session
}

func (c *conn) Close() {
	c.session.Close()
}

func InitDB() Connection {
	var c conn
	var err error
	c.session, err = sql.Open("mysql", connstr())
	if err != nil {
		panic(err)
	}

	DBConn, err = gorm.Open(mysql.New(mysql.Config{
		Conn: c.session,
	}))

	if err != nil {
		panic("Failed connect to Database")
	}
	MigrateDB(c.session, DBConn, &models.Product{})
	return &c
}

func MigrateDB(conn *sql.DB, db *gorm.DB, model interface{}) {
	result := db.AutoMigrate(model)
	if result != nil {
		fmt.Println("Database Already")
	}
}

func connstr() string {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Println("error on load db port from env:", err.Error())
		port = 27017
	}
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_PROTOCOL"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"),
	)
}
