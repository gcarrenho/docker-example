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

// Config app
type Config struct {
	Environment string
	DBConfig    mysql.MySQLConf
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Error().Msg("Could not load config: " + err.Error())
		return
	}

	db, err := config.DBConfig.InitMySqlDB(config.DBConfig)
	if err != nil {
		log.Error().Msg("Could not initialize database: " + err.Error())
		return
	}

	defer db.Close()

	mysqlOrdersRepository := orders.NewMYSQLOrdersRepository(db)
	ordersComponent := orders.NewOrdersComponentImpl(*mysqlOrdersRepository)

	router := setupRouter()

	web.NewOrdersHandlers(router.Group("/"), ordersComponent)

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

	// Starting the server http in a gorutine
	go func() {
		if err := ginSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msg("Could not start http server: " + err.Error())
		}
	}()

	// Wait ending signal of context
	<-ctx.Done()

	// Stop the server and clean the resources
	stop()

	// Creating a child context with a waiting time to close the server
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := ginSrv.Shutdown(ctxShutdown); err != nil {
		log.Error().Msg("Forcing shutdown: " + err.Error())
	}

	log.Debug().Msg("Stopped")
}

func loadConfig() (Config, error) {
	var config Config

	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Overload(fmt.Sprintf("../%s.env", env))
	if err != nil {
		return config, err
	}

	dbConfig := mysql.NewMySqlConf(os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASSWORD"))
	config = Config{
		Environment: env,
		DBConfig:    dbConfig,
	}

	return config, err
}
