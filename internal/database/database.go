package database

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"
)

var (
	Engine *xorm.Engine
)

func Init() error {
	var err error
	Engine, err = xorm.NewEngine("sqlite3", viper.GetString("database.dsn"))
	if err != nil {
		return err
	}
	
	return nil
}

func Migrate(migrations []*migrate.Migration) error {
	m := migrate.New(Engine, migrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
