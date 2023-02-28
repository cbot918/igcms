package main

import (
	"fmt"

	// "web/internal/api"
	"web/internal/config"
	"web/internal/db"

	u "github.com/cbot918/liby/util"
	_ "github.com/lib/pq"
)


func main() {

  // config instance
  config := config.NewConfig()

  // db conn instance
  conn := db.NewDbConn(config); defer conn.Close()
  // query := "insert into users(name) values('node')"
  query := "create table test1();"
  result, err := conn.Exec(query); u.Checke(err, "conn exec failed")
  fmt.Println("result: ",result)

  
  // api := api.NewServer()
  // api.Listen(3000)

}