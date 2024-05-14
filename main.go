package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ilcm96/dku-aegis-library/middleware"

	"github.com/ilcm96/dku-aegis-library/controller"
	"github.com/ilcm96/dku-aegis-library/repository"
	"github.com/ilcm96/dku-aegis-library/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
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

	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Global middleware
	app.Use(middleware.NewSlog(logger))
	app.Use(recover.New())
	app.Use(pprof.New())

	// Metric
	app.Get("/monitor", monitor.New())

	// Dependency
	logRepository := repository.NewLogRepository(db.Client)

	userRepository := repository.NewUserRepository(db.Client)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository(db.Client)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService, logRepository)

	viewController := controller.NewViewController(bookRepository, logRepository)

	// --------------------
	// --- Public Route ---
	// --------------------

	app.Get("/signup", viewController.SignUp)
	app.Get("/login", viewController.Login)
	app.Post("/api/user/create", userController.SignUp)
	app.Post("/api/user/login", userController.SignIn)

	// ------------------------
	// --- Restricted route ---
	// ------------------------

	// JWT middleware
	app.Use(middleware.NewJwt())

	// Static asset
	app.Static("/asset", "./asset")

	// View route
	app.Get("/", viewController.Index)
	app.Get("/mypage", viewController.MyPage)
	app.Get("/search", viewController.Search)

	// Api route
	app.Post("/api/book/borrow", bookController.BorrowBook)
	app.Post("/api/book/return", bookController.ReturnBook)
	app.Post("/api/book/search", viewController.SearchResult)

	// --------------------
	// --- END OF ROUTE ---
	// --------------------

	// Run app
	log.Fatal(app.Listen(":3000"))
}
