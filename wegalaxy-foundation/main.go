package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/degalaxy/gp-common/filter"
	"github.com/degalaxy/gp-common/gp_error"
	"github.com/degalaxy/gp-common/log"
	gp_repo "github.com/degalaxy/gp-common/repository"
	gp_service "github.com/degalaxy/gp-common/service"
	"github.com/degalaxy/gp-common/util"
	"github.com/degalaxy/helper/dahelper"
	common_ctrl "github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/common/controller"
	"github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/config"
	health_ctrl "github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/health-check/controller"
	us_ctrl "github.com/degalaxy/wegalaxy-service/wegalaxy-foundation/user-space/controller"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
)

// @title           WeGalaxy Foundation Service
// @version         1.0
// @description     WeGalaxy Foundation Service

// @host      localhost:8080
// @BasePath  /v1/wegalaxy-foundation

const (
	ENV_APP_CONFIG_FILE = "APP_CONFIG_FILE"
	DEFAULT_CONFIG_FILE = "resource/config.yaml"
)

func main() {
	configFilePath := os.Getenv(ENV_APP_CONFIG_FILE)
	if configFilePath == "" {
		configFilePath = DEFAULT_CONFIG_FILE
	}
	initConfig(configFilePath)
	runServerAndWait()
	log.Logger.Info("Server exited")
}

func initConfig(configFilePath string) {
	config.InitConfigAndWatchChanges(configFilePath, configChangeHandler)
	initLogger()
	// Init key
	isDSSourceAuth := config.GlobalConfig.App.UserServiceSourceAuth
	privateKeyPath := strings.TrimSpace(config.GlobalConfig.Server.PrivateKeyPath)
	if isDSSourceAuth {
		if privateKeyPath == "" {
			log.Logger.Fatal("Private key is not set")
		}
		b, err := ioutil.ReadFile(privateKeyPath)
		if err != nil {
			log.Logger.Fatal(fmt.Sprintf("Cannot read file %s: %v", privateKeyPath, err))
		}
		config.GlobalConfig.Server.PrivateKey, err = util.ParsePKCS8PrivateKey(b)
		if err != nil {
			log.Logger.Fatal(fmt.Sprintf("Cannot parse key file %s: %v", privateKeyPath, err))
		}
	}
}

func configChangeHandler(e fsnotify.Event) {
	// Need to reload config and do things necessary
	log.Logger.Info("[Config] File got updated - " + e.String())
}

func initLogger() {
	log.InitLogger(config.GlobalConfig.Log)
}

func runServerAndWait() {
	gin.DisableConsoleColor()
	r := gin.New()

	r.Use(util.CORSMiddleware()) //allows all origin
	r.Use(processFilter()...)
	gin.Recovery()
	initAPI(r)

	// cron job
	log.Logger.Info("Server port: " + config.GlobalConfig.Server.Port)

	srv := &http.Server{
		Addr:    ":" + config.GlobalConfig.Server.Port,
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal("Error during listen: " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Logger.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server forced to shutdown: " + err.Error())
	}

	log.Logger.Info("Server exiting")
}

func initAPI(r *gin.Engine) {

	// ------------------------
	// Controller context
	// ------------------------
	var (
		// redisConnection *gp_repo.RedisConnection
		dbConnection *gp_repo.DBConnection
		err          error
	)
	// if dbConnection, err = gp_repo.NewDBConnection(
	// 	config.GlobalConfig.DB.Login,
	// 	config.GlobalConfig.DB.Password,
	// 	config.GlobalConfig.DB.Host,
	// 	config.GlobalConfig.DB.Port,
	// 	config.GlobalConfig.DB.DatabaseName,
	// 	config.GlobalConfig.DB.ConnectionMaxIdle,
	// 	config.GlobalConfig.DB.ConnectionMaxOpen,
	// ); err != nil {
	// 	log.Logger.Fatal("Error initializing DB Connection" + err.Error())
	// }
	// Init Repository

	verifyTokenService, err := gp_service.NewVerifyTokenService(config.GlobalConfig.AppAuth)
	if err != nil {
		log.Fatalf("Error initializing verify token service, err:%v", err)
	}

	daConfig := &dahelper.DAConfig{
		UserServiceUrl:       config.GlobalConfig.App.UserServiceUrl,
		NeedSourceAuthHeader: config.GlobalConfig.App.UserServiceSourceAuth,
		TradeServiceUrl:      config.GlobalConfig.App.TradeServiceUrl,
		Proxy:                "",
		AppId:                config.GlobalConfig.App.AppId,
		PrivateKey:           config.GlobalConfig.Server.PrivateKey,
	}
	daHelper := dahelper.NewDAHelper(daConfig)

	// Init Controller
	// lckMgr := lockmanager.NewDBLockManager(dbConnection, lockmanager.DEFAULT_KEY_EXPIRY_CHECK_TIME_INTERVAL)
	ctrlCtx, err := common_ctrl.NewControllerContext(
		verifyTokenService,
		daHelper,
		dbConnection, // redisConnection,
	)
	if err != nil {
		log.Logger.Fatal("Error initializing ControllerContext" + err.Error())
	}

	// ------------------------
	// Controller
	// ------------------------
	var (
		healthCheckCtrl *health_ctrl.HealthCheckController
		userCtrl        *us_ctrl.UserController
	)
	if healthCheckCtrl, err = health_ctrl.NewHealthCheckController(); err != nil {
		log.Logger.Fatal("Error initializing HealthCheckController" + err.Error())
	}
	if userCtrl, err = us_ctrl.NewUserController(ctrlCtx); err != nil {
		log.Logger.Fatal("Error initializing UserController" + err.Error())
	}

	// ------------------------
	// Routes
	// ------------------------
	v1 := r.Group("/v1/wegalaxy-foundation")
	{
		user := v1.Group("/users")
		{
			user.GET("", userCtrl.GetUser)
		}
		v1.GET("/health-check", healthCheckCtrl.HealthCheck)
	}
}

func processFilter() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		filter.LogServerFilter,
		filter.RequestIdFilter,
		gin.CustomRecovery(recoverHandler)}
}

// Globally handle panic here, print panic message and return 500
func recoverHandler(c *gin.Context, err interface{}) {
	log.Logger.Error(fmt.Sprintf("Panic occurred: %s", err))
	gp_error.RaiseHttpError(c, gp_error.ErrInternalError)
}
