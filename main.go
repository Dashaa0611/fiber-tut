// package main

// import (
// 	"github.com/gofiber/fiber/v2"
// )

// type SomeStruct struct {
// 	Name string
// 	Age  uint8
// }

// type User struct {
// 	data SomeStruct
// }

// func main() {
// 	app := fiber.New()
// 	app.Get("/", func(c *fiber.Ctx) error {

// 		var user = User{
// 			data: SomeStruct{
// 				Name: "das",
// 				Age:  23,
// 			},
// 		}
// 		return c.JSON(user.data)

// 	})

// 	app.Listen(":3002")
// }
