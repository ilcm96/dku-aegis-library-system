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
	"github.com/gofiber/template/html/v2"
	"github.com/ilcm96/dku-aegis-library/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init database
	db.InitDB()
	db.InitRedisClient()
	db.InitMinioClient()

	// Template engine setting
	engine := html.New("./template", ".html")
	engine.Reload(true)
	engine.Debug(false)
	engine.AddFunc("sub", func(x, y int) int {
		return x - y
	})

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
	//app.Use(recover.New())
	app.Use(pprof.New())

	// Metric
	app.Get("/monitor", monitor.New())

	// Dependency
	logRepository := repository.NewLogRepository(db.Client)

	userRepository := repository.NewUserRepository(db.Client)
	userService := service.NewUserService(userRepository, db.RedisClient())
	userController := controller.NewUserController(userService)

	bookRepository := repository.NewBookRepository(db.Client, db.MinioClient())
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService, logRepository)

	bookReqRepository := repository.NewBookReqRepository(db.Client)
	bookReqController := controller.NewBookReqController(bookReqRepository)

	viewController := controller.NewViewController(userRepository, bookRepository, logRepository, bookReqRepository)

	// --------------------
	// --- Public Route ---
	// --------------------
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Get("/signup", middleware.NewNoSigninUser(), viewController.SignUp)
	app.Get("/signin", middleware.NewNoSigninUser(), viewController.Signin)
	app.Post("/api/signup", middleware.NewNoSigninUser(), userController.SignUp)
	app.Post("/api/signin", middleware.NewNoSigninUser(), userController.SignIn)

	// ------------------------
	// --- Restricted route ---
	// ------------------------

	// Session middleware
	app.Use(middleware.NewSessionAuth(db.RedisClient()))
	app.Use(middleware.NewRenewSession(db.RedisClient()))

	// Static asset
	app.Static("/asset", "./asset")

	// View route
	app.Get("/", viewController.Index)
	app.Get("/mypage", viewController.MyPage)
	app.Get("/search", viewController.Search)
	app.Get("/book/:id", viewController.BookDetail)
	app.Get("/request", viewController.BookRequest)
	app.Get("/request/history", viewController.BookRequestHistory)

	// Api route
	app.Post("/api/signout", userController.SignOut)
	app.Post("/api/user/withdraw", userController.Withdraw)

	app.Post("/api/book/borrow", bookController.BorrowBook)
	app.Post("/api/book/return", bookController.ReturnBook)
	app.Post("/api/book/search", viewController.SearchResult)

	app.Post("/api/request", bookReqController.CreateBookReq)
	app.Delete("/api/request/:id", bookReqController.DeleteBookReq)

	// -------------------
	// --- Admin route ---
	// -------------------

	app.Use(middleware.NewIsAdmin())
	app.Get("/admin", viewController.Admin)

	app.Get("/admin/user", viewController.AdminUser)
	app.Get("/admin/user/:id", viewController.AdminUserDetail)
	app.Get("/admin/book", viewController.AdminBook)
	app.Get("/admin/book/create", viewController.AdminBookCreate)
	app.Get("/admin/book/:id", viewController.AdminBookDetail)
	app.Get("/admin/request", viewController.AdminRequest)
	app.Get("/admin/request/:id", viewController.AdminRequestDetail)

	app.Post("/api/admin/book", bookController.AdminCreateBook)
	app.Put("/api/admin/book/:id", bookController.AdminUpdateBook)
	app.Post("/api/admin/book/cover/:id", bookController.AdminUpdateBookCover)
	app.Delete("/api/admin/book/:id", bookController.AdminDeleteBook)
	app.Post("/api/admin/user/:id", userController.ChangeStatus)
	app.Post("/api/admin/request/:id", bookReqController.UpdateBookReqApproved)

	// --------------------
	// --- END OF ROUTE ---
	// --------------------

	// Run app
	log.Fatal(app.Listen(":3000"))
	app.Shutdown()
}
