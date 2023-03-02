package controller

import (
	"database/sql"
	"fmt"
	"web/internal/model"

	"github.com/cbot918/liby/util"
	"github.com/gofiber/fiber"
)

const (
	NO_TABLE_ERROR = "NO_TABLE_ERROR"
)

type Common struct  {
	DB *sql.DB
}
func NewCommon(db *sql.DB) *Common{
	c := new(Common)
	c.DB = db
	return c
}


func (c *Common) Ping(ctx *fiber.Ctx){
	ctx.Send("pong")
}

func (c *Common) GetAll(ctx *fiber.Ctx) {
	errStr := c.selectAll( ctx.Params("table") )
	if errStr == NO_TABLE_ERROR { 
		ctx.Send("table 不存在")
		return 
	} 
	ctx.Send("查詢成功")
}

func (c *Common) selectAll(tableName string) string {
	row, err := c.DB.Query(fmt.Sprintf("select * from %s", tableName)); 
	if err != nil {
		return NO_TABLE_ERROR
	}
	defer row.Close()

	var users []model.Users
	for row.Next(){ 
		var user model.Users
		err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); util.Checke(err,"row scan failed")
		users = append(users, user)
	} 
	fmt.Println("users: ",users)
	return ""
}
