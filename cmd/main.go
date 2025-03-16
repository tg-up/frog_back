package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "icecreambash/tgup_backend/docs"
	"icecreambash/tgup_backend/internal/configs"
	"icecreambash/tgup_backend/internal/routes"
	"icecreambash/tgup_backend/pkg/database"
	"icecreambash/tgup_backend/seeds/platform"
	"log"
)

func init() {
	configs.LoadConfig()
	err := database.InitDB()
	if err != nil {
		log.Fatalf("init db err: %v", err)
	}
}

// @title           Swagger TG-UP API DOCS
// @version         0.0.2
// @description     API для взаимодействия с TG-UP
// @termsOfService  http://tg-up.com/terms/legal

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	seedFlag := flag.String("seed", "", "seed file")
	flag.Parse()

	if *seedFlag != "" {
		platform.PlatformSeeds()
		return
	}

	if configs.GlobalConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	ginDemon := gin.Default()

	routes.InjectRoutes(ginDemon)

	ginDemon.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := ginDemon.Run(configs.GlobalConfig.ServerHost + ":" + configs.GlobalConfig.ServerPort); err != nil {
		log.Fatalf("Down on startup ginDemon: %v", err)
	}
}
