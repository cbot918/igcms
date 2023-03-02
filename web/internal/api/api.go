package api

import (
	"database/sql"
	"web/internal/api/controller"

	"github.com/gofiber/fiber"
)


func NewServer(db *sql.DB) *fiber.App{
  app := fiber.New()

  // app.Use(Jwt())

  app.Static("/", "client/build")

  app.Get("/web/*", func(ctx *fiber.Ctx) {
    ctx.SendFile("client/build/index.html")
  })

  c := controller.NewCommon(db)
  a := controller.NewAuth(db)
  p := controller.NewPost(db)

  app.Get("/api/ping", c.Ping)
  app.Get("/api/getall/:table",c.GetAll)

  app.Post("/api/login", a.Login)
  app.Post("/api/regist", a.Regist)
  app.Post("/api/validator", a.Validator)

  app.Post("/api/post/create", p.CreatePost)
  app.Get("/api/post/read/all", p.CreatePost)
  app.Post("/api/post/read/:pid", p.CreatePost)
  app.Post("/api/post/delete/:pid", p.CreatePost)

 return app
}