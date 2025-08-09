package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/caoyong2619/elotus/internal/config"
	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/database/migrations"
	"github.com/spf13/viper"
	"xorm.io/xorm/migrate"
)

func testError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

func TestMain(m *testing.M) {
	var err error

	if err := config.Init(`./config.yaml`); err != nil {
		testError(err)
	}

	// remove db file if exists
	_ = os.Remove(viper.GetString(`database.dsn`))

	if err := database.Init(); err != nil {
		testError(err)
	}

	if err != nil {
		fmt.Printf("create db file failed, err: %v", err)
		os.Exit(1)
	}

	mig := migrate.New(database.Engine, migrate.DefaultOptions, migrations.Migrations())

	if err := mig.Migrate(); err != nil {
		testError(err)
	}

	os.Exit(m.Run())
}
