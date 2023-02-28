package api

import "github.com/gofiber/fiber"


func NewServer() *fiber.App{
  app := fiber.New()

  app.Static("/", "client/build")

  app.Get("/*", func(ctx *fiber.Ctx) {
    ctx.SendFile("client/build/index.html")
 })

 return app
}