package migrations

import (
	"github.com/caoyong2619/elotus/internal/database"
	"xorm.io/xorm"
	"xorm.io/xorm/migrate"
)

func Migrations() []*migrate.Migration {
	return []*migrate.Migration{
		{
			ID: "20250809204037795",
			Migrate: func(engine *xorm.Engine) error {
				return engine.Sync2(&database.User{}, &database.AuthToken{}, &database.Upload{})
			},
			Rollback: func(engine *xorm.Engine) error {
				return engine.DropTables(&database.User{}, &database.AuthToken{}, &database.Upload{})
			},
		},
	}
}
