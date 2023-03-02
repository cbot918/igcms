package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"web/internal/model"
	"web/internal/util"

	"github.com/cbot918/liby/jwty"
	"github.com/gofiber/fiber"
)

type Post struct{
	db *sql.DB
}

func NewPost(db *sql.DB) *Post{
	p := new(Post)
	p.db = db
	return p
}

type postResponse struct{
	Result	string `json:"result"`
}
func (p *Post) CreatePost(ctx *fiber.Ctx) {

	header := ctx.Fasthttp.Request.Header.String()
	token := util.GetJwtToken(header)
	token = strings.Trim(token, " ") // work-around 不知道為什麼會怪怪的
	
	j := jwty.New()
	id,email := j.GetIdAndEmail(j.DecodeJwt(token))
	// fmt.Println("decoded email: ",email); fmt.Println("decoded id: ",id)

	post := &model.Post{}
	if err := ctx.BodyParser(post); err != nil {panic(err)}
	// fmt.Println("title: ",post.Title); fmt.Println("content: ",post.Content)

	query := fmt.Sprintf("INSERT INTO post (user_id,title,post_content) VALUES (%d,'%s','%s')",
							id, post.Title, post.Content)
	_,  err := p.db.Exec(query); if err != nil { panic(err) }
	
	res := &postResponse{ Result: "PO文成功" }
	rJson, err := json.Marshal(res); if err != nil {panic(err)}

	fmt.Printf("\n系統訊息: post")
	fmt.Printf("\ntoken: %s", token)
	fmt.Println("user: ", email)
	fmt.Printf("title: %s\ncontent: %s\n\n",post.Title,post.Content)
	fmt.Println("送回客戶端")
	fmt.Println(string(rJson))
	ctx.Set("Content-Type","application/json")
	ctx.SendString(string(rJson))
}

func (p *Post) ReadAllPost(){
	
}

func (p *Post) ReadPosT(){

}

func (p *Post) DeletePosT(){

}