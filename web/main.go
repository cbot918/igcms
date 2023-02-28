package main

import "github.com/gofiber/fiber"



func main() {
  app := fiber.New()
  app.Static("/", "client/build")
  app.Get("/*", func(ctx *fiber.Ctx) {
    ctx.SendFile("client/build/index.html")
 })
  app.Listen(3001)
}