package main

import (
	"flag"
	"fmt"
	"log"

	"os"
	"path/filepath"

	"github.com/caoyong2619/elotus/internal/config"
	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/database/migrations"
	"github.com/caoyong2619/elotus/internal/route"
	"github.com/caoyong2619/elotus/internal/services"
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
		runMigrate(args[1:])
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

	authService := services.NewAuthService(database.Engine)
	auth := route.NewAuth(authService)

	e.POST(`/auth/register`, auth.Register())

	err := e.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func runMigrate(args []string) {
	m := migrate.New(database.Engine, migrate.DefaultOptions, migrations.Migrations())

	if len(args) == 0 {
		args = []string{"migrate"}
	}

	switch args[0] {
	case "migrate":
		if err := m.Migrate(); err != nil {
			log.Fatal(err.Error())
		}

		log.Println("migrate success")
	case `rollback`:
		rollbackArgs := args[1:]

		// if rollback id is empty, rollback last migration
		if len(rollbackArgs) == 0 {
			if er := m.RollbackLast(); er != nil {
				log.Fatal(er.Error())
			}

			log.Println("rollback success")
			return
		}

		mig := &migrate.Migration{
			ID: rollbackArgs[1],
		}

		if err := m.RollbackMigration(mig); err != nil {
			log.Fatal(err.Error())
		}

		log.Println("rollback success")
	}
}

func rollbackLatest(m *migrate.Migrate) error {
	return m.RollbackLast()
}
