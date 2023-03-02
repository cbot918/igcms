package main

import (

	// "web/internal/api"
	"web/internal/api"
	"web/internal/config"
	"web/internal/db"

	_ "github.com/lib/pq"
)


func main() {

  // config instance
  config := config.NewConfig()

  // db conn instance
  conn := db.NewDbConn(config); defer conn.Close()

  // query := "insert into users(name) values('node')"
  // query := "create table test1();"
  // query := "INSERT INTO users (name, email, password) VALUES ('yale', 'yale918@gmail.com', '12345')"
  // result, err := conn.Exec(query); u.Checke(err, "conn exec failed")
  // fmt.Println("result: ",result)

  
  
  api := api.NewServer(conn)
  api.Listen(3000)

}