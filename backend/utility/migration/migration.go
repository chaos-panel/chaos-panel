package migration

import (
	// https://github.com/gogf/gf/tree/master/contrib/drivers
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"

	// https://github.com/golang-migrate/migrate
	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func init() {

	groups := GetDatabaseGroups()

	if len(groups) <= 0 {
		return
	}

	for key, node := range groups {
		from, err := NewEmbededRes("manifest/sql/" + key)

		if err != nil {
			panic(err)
		}

		db, err := g.DB(key).Open(node)

		if err != nil {
			panic(err)
		}

		var to database.Driver

		switch node.Type {
		case "tidb":
			to, err = mysql.WithInstance(db, &mysql.Config{})
		case "mysql":
			to, err = mysql.WithInstance(db, &mysql.Config{})
		case "mariadb":
			to, err = mysql.WithInstance(db, &mysql.Config{})
		case "sqlite":
			to, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		case "pgsql":
			to, err = postgres.WithInstance(db, &postgres.Config{})
		}
		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithInstance("gres", from, node.Type, to)
		if err != nil {
			panic(err)
		}

		g.Log().Info(context.Background(), "db migration start")
		err = m.Up()
		if err != nil {
			g.Log().Info(context.Background(), "db migration error")
			panic(err)
		}
		g.Log().Info(context.Background(), "db migration done")
	}
}
