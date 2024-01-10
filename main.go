package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task_tracker/internal/handler"
	"task_tracker/internal/repository"
	"task_tracker/internal/service"
)

var (
	dbHost              = flag.String("dbHost", "localhost", "database host")
	dbUser              = flag.String("dbUser", "task_tracker", "database user")
	dbPass              = flag.String("dbPass", "12332145", "database password")
	dbName              = flag.String("dbName", "task_tracker", "database name")
	dbPort              = flag.String("dbPort", "5432", "database port")
	defaultAllowOrigins = flag.String("cors", "*", "defaultAllowOrigins value string")
	appPort             = flag.String("appPort", "3030", "application start port")

	Conn              *sqlx.DB
	logFormat         = "[${time}] ${status} - ${latency} ${method} ${path}\n"
	defaultCorsMaxAge = 3600
)

func connect() {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", *dbUser, *dbPass, *dbHost, *dbPort, *dbName))
	if err != nil {
		panic(err)
	}
	if db == nil {
		log.Fatalln("не удалось подключиться к базе данных!")
	}
	Conn = db
}

func closeConnection(connect *sqlx.DB) error {
	return connect.Close()
}

func main() {
	flag.Parse()
	connect()
	// Create a new engine
	engine := html.New("./views", ".html")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".html"))

	app := fiber.New(fiber.Config{
		Prefork:   false,
		BodyLimit: 16 * 1024 * 1024,
		Views:     engine,
	})
	//middlewares
	app.Use(logger.New(logger.Config{
		Format: logFormat,
		Output: os.Stdout,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     *defaultAllowOrigins,
		AllowCredentials: true,
		MaxAge:           defaultCorsMaxAge,
	}))
	app.Static("/assets", "./assets")
	rep := repository.NewRepository(Conn)
	services := service.NewService(rep)
	service.Services = services
	h := handler.NewHandler(services)
	h.Init(app)
	//start server
	go func() {
		if err := app.Listen(":" + *appPort); err != nil {
			log.Panicf("не удалось запустить инстанс веб-сервера: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c
	fmt.Println("Остановка веб-сервера...")
	_ = app.Shutdown()
	fmt.Println("--> соединение с базой данных закрыто")
	err := Conn.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("--> сервис остановлен")
}
