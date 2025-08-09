package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"
	"os"
	"path/filepath"

	"github.com/caoyong2619/elotus/internal/config"
	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/database/migrations"
	"github.com/gin-gonic/gin"
	"xorm.io/xorm/migrate"
)

var (
	configFile string
)

func main() {
	bootstrap()
	args := os.Args[1:]

	if len(args) == 0 {
		args = []string{`usage`}
	}

	switch args[0] {
	case "serve":
		serve()
	case "migrate":
		runMigrate()
	default:
		fmt.Println("Usage: elotus [serve|migrate]")
	}

}

func init() {
	wd, _ := os.Getwd()
	flag.StringVar(&configFile, "config", filepath.Join(wd, "config.yaml"), "config file path")

}

func bootstrap() error {
	flag.Parse()

	if err := config.Init(configFile); err != nil {
		log.Fatal(err.Error())
	}

	if err := database.Init(); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func serve() {
	e := gin.Default()

	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func runMigrate() {
	m := migrate.New(database.Engine, migrate.DefaultOptions, migrations.Migrations())
	if err := m.Migrate(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("migrate success")
}
