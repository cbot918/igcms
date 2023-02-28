package db

import (
	"database/sql"
	"fmt"
	"log"
	"web/internal/config"

	u "github.com/cbot918/liby/util"
)

func NewDbConn(c *config.Config) *sql.DB{
  // init database
  connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",c.Host,c.Port,c.User,c.Password,c.Dbname)
  conn, err := sql.Open("postgres",connStr); u.Checke(err, "sql open failed"); 
  if err := conn.Ping(); err != nil { log.Fatal(err) }
  return conn
}
