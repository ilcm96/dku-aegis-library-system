package main

import (
	"log"
	"net/url"

	"github.com/ilcm96/dku-aegis-library/controller"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilcm96/dku-aegis-library/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init database
	db.InitDB()

	// Template engine setting
	engine := html.New("./template", ".html")
	engine.Reload(true)
	engine.Debug(false)

	// Fiber config
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base",
		BodyLimit:   50 * 1024 * 1024,
	})

	// Global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(pprof.New())

	// Metric
	app.Get("/monitor", monitor.New())

	// Dependency
	userRepository := repository.NewUserRepository(db.Client)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository(db.Client)

	viewController := controller.NewViewController(bookRepository)

	// Route
	app.Get("/signup", viewController.SignUp)
	app.Get("/login", viewController.Login)
	app.Post("/api/user/create", userController.SignUp)
	app.Post("/api/user/login", userController.SignIn)

	// JWT middleware
	app.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:token",
		Validator: validateJWT,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			redirectURL := "/login?next=" + url.QueryEscape(c.OriginalURL())
			return c.Redirect(redirectURL)
		},
	}))

	app.Static("/", "./public")
	// View route
	app.Get("/", viewController.Index)
	app.Get("/book", func(c *fiber.Ctx) error {
		return c.SendString("book")
	})

	// Run app
	log.Fatal(app.Listen(":3000"))
}

func validateJWT(c *fiber.Ctx, token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwt-secret"), nil
	})
	if err != nil || !parsedToken.Valid {
		return false, c.Redirect("/login")
	}
	return true, nil
}
