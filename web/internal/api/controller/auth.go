package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"web/internal/model"

	"github.com/cbot918/liby/jwty"
	"github.com/gofiber/fiber"
)

type Auth struct {
	DB *sql.DB
}
func NewAuth(db *sql.DB) *Auth{
	a := new(Auth)
	a.DB = db
	return a
}

type loginResponse struct{
	AuthStatus	string `json:"authstatus"`
	Token				string `json:"token"`
}
func (a *Auth) Login(ctx *fiber.Ctx){
	user := &model.Users{}; ctx.BodyParser(user); 
	
	if !a.emailExists(user.Email) { 
		ctx.SendString("帳號不存在, 請重新註冊帳號")
		return
	}

	token, err:= jwty.New().FastJwt(int(user.Id), user.Email); if err != nil {panic(err)}
	
	res := &loginResponse{ AuthStatus: "登入成功", Token: token, }
	rJson, err := json.Marshal(res); if err != nil {panic(err)}
	fmt.Println("系統訊息: login")
	fmt.Printf("name: %s\nemail: %s\n\n",a.getName(user.Email),user.Email)
	fmt.Println("送回客戶端")
	fmt.Println(string(rJson))
	ctx.Set("Content-Type","application/json")
	ctx.SendString(string(rJson))
}

func (a *Auth) Regist(ctx *fiber.Ctx) {
	user := &model.Users{}; ctx.BodyParser(user)
	
	if a.emailExists(user.Email){
		ctx.SendString("此email已有人使用")
		return
	}

	if !a.regist(user.Email,user.Password){
		ctx.SendString("註冊失敗, 請調整資料重新嘗試")
		return
	}

	ctx.SendString("註冊成功, 請手動跳轉登入")
}

func (a *Auth) Validator(ctx *fiber.Ctx) {
	header := ctx.Fasthttp.Request.Header.String()
	if !a.jwtExists(header){
		fmt.Println("jwt驗證失敗")
		return
	}
	fmt.Println("jwt驗證成功")
	fmt.Println(a.getJwtString(header),)
}

func (a *Auth) jwtExists(header string)  bool {
	return regexp.MustCompile(`Authorization: Bearer`).MatchString(header)
}

func (a *Auth) getJwtString(header string) string {
	return regexp.MustCompile(`Authorization: .*\r`).FindString(header)
}

func (a *Auth) emailExists(targetEmail string) bool {
	var temp string
	query := fmt.Sprintf("SELECT email FROM users WHERE email='%s'",targetEmail)
	err := a.DB.QueryRow(query).Scan(&temp)
	if err != nil { 
		if err == sql.ErrNoRows{ return false }
	}
	return true
}

func (a *Auth) regist(email string, password string) bool {
	name := a.getName(email)
	
	query := fmt.Sprintf("insert into users (email, password, name) values ('%s','%s','%s')",
							email, password, name )
	_ , err := a.DB.Exec(query)

	return err == nil
}

func (a *Auth) getName(email string) string {
	return strings.Trim(regexp.MustCompile(`.*@`).FindString(email), "@")
}