package main

import(
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gilangmahardhika/go-fiber-tasklist/handler"
)

func main() {

	app := fiber.New()
	
	app.Get("/", tasks.GetHelloWorld)

	log.Fatal(app.Listen(":8080"))

}
