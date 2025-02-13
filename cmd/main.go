package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"icecreambash/tgup_backend/internal/config"
	"icecreambash/tgup_backend/internal/routes"
	"icecreambash/tgup_backend/pkg/database"
	"icecreambash/tgup_backend/seeds/platform"
	"log"
)

func init() {
	config.LoadConfig()
	err := database.InitDB()
	if err != nil {
		log.Fatalf("init db err: %v", err)
	}
}

func main() {
	seedFlag := flag.String("seed", "", "seed file")
	flag.Parse()

	if *seedFlag != "" {
		platform.PlatformSeeds()
		return
	}

	if config.GlobalConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	ginDemon := gin.Default()

	routes.InjectRoutes(ginDemon)

	if err := ginDemon.Run(config.GlobalConfig.ServerHost + ":" + config.GlobalConfig.ServerPort); err != nil {
		log.Fatalf("Down on startup ginDemon: %v", err)
	}
}
