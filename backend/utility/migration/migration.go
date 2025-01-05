package migration

import (
	// https://github.com/gogf/gf/tree/master/contrib/drivers
	"context"
	"strconv"
	"strings"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gres"

	// https://github.com/golang-migrate/migrate
	"github.com/golang-migrate/migrate/v4"

	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func Migrate(ctx context.Context) {
	migration := gbuild.Info().Data["migration"]
	if migration == nil {
		g.Log().Info(ctx, "migration ignore: build info not found")
		return
	}
	items := strings.Split(migration.(string), ",")
	if len(items) == 0 {
		g.Log().Info(ctx, "migration ignore: build info not found")
		return
	}
	for _, item := range items {
		if item == "" {
			continue
		}
		if !strings.Contains(item, ":") {
			continue
		}
		config := strings.SplitN(item, ":", 2)
		if len(config) != 2 {
			continue
		}
		path := config[0]
		version := config[1]
		if path == "" || version == "" {
			continue
		}
		handleMigration(ctx, path, version)
	}
}

func handleMigration(ctx context.Context, path string, version string) {

	steps, err := strconv.Atoi(version)
	if err != nil {
		panic(err)
	}

	groups := GetDatabaseGroups()

	if len(groups) <= 0 {
		g.Log().Info(ctx, "migration ignore: no database config found")
		return
	}

	for key, node := range groups {
		entries := gres.ScanDirFile(path+key, "*.sql", true)
		if len(entries) == 0 {
			break
		}

		from, err := NewEmbededRes(path + key)

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
		g.Log().Info(ctx, "db migration start")
		err = m.Migrate(uint(steps))
		_, _ = m.Close()
		if err == nil {
			g.Log().Info(ctx, "db migration done")
			return
		}
		if strings.Contains(err.Error(), "no change") {
			g.Log().Info(ctx, "db migration done: no change")
			return
		} else {
			g.Log().Info(ctx, "db migration error")
			panic(err)
		}
	}
}
