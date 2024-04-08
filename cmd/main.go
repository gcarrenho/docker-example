package main

import (
	"context"
	"docker-example/databases/mysql"
	"docker-example/internal/app/handlers/web"
	"docker-example/internal/app/orders"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	localPort = "8080"
)

func main() {

	err := loadConfig()
	if err != nil {
		fmt.Println("fail loading fiile .env", err)
	}

	// DB connection
	dbConfig := mysql.NewMySqlConf(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASSWORD"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&clientFoundRows=true",
		dbConfig.DbUser, dbConfig.DbPassword, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName)

	db, err := dbConfig.InitMySqlDB(dsn)
	if err != nil {
		//logger.Error().Msg(err.Error())
		fmt.Println("problem ", err)

		panic("problem with db ")
	}

	//db := dbConfig.GetDB()
	defer db.Close()

	mysqlOrdersRepository := orders.NewMYSQLOrdersRepository(db)
	ordersComponent := orders.NewOrdersComponentImpl(*mysqlOrdersRepository)

	router := setupRouter()

	web.NewOrdersHandlers(router, ordersComponent)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	runServer(ctx, stop, router)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		requestid.New(),
		//middleware.RequestIDMiddleware,
		//gin.LoggerWithConfig(logConf()),
		gin.Recovery(),
	)

	return router
}

func runServer(ctx context.Context, stop context.CancelFunc, router *gin.Engine) {
	log.Debug().Msg("Running")
	// HTTP Server
	ginSrv := &http.Server{
		Addr:    ":" + localPort,
		Handler: router,
	}
	go func() {
		runErr := ginSrv.ListenAndServe()
		if runErr != nil && !errors.Is(runErr, http.ErrServerClosed) {
			log.Fatal().Msg("could not start http server: " + runErr.Error())
		}
	}()

	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if shutdownErr := ginSrv.Shutdown(ctx); shutdownErr != nil {
		log.Error().Msg("Forcing shutdown: " + shutdownErr.Error())
	}
	log.Debug().Msg("Stopped")
}

func loadConfig() error {
	var err error
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		err = godotenv.Overload("../dev.env")
	case "staging":
		err = godotenv.Overload("../stg.env")
	case "prod":
		err = godotenv.Overload("../pro.env")
	default:
		err = godotenv.Overload("../dev.env")
	}

	return err
}
