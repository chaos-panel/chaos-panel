package migration

import (
	// https://github.com/gogf/gf/tree/master/contrib/drivers
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/chaos-plus/chaos-plus/utility/utils"
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
		g.Log().Warning(ctx, "migration ignore: build info not found")
		return
	}
	items := strings.Split(migration.(string), ",")
	if len(items) == 0 {
		g.Log().Warning(ctx, "migration ignore: build info not found")
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

	groups := utils.GetDatabaseGroups()

	if len(groups) <= 0 {
		g.Log().Warning(ctx, "migration ignore: no database config found")
		return
	}

	for group, node := range groups {
		entries := gres.ScanDirFile(path+node.Type, "*.sql", true)
		if len(entries) == 0 {
			break
		}

		from, err := NewEmbededRes(path + node.Type)

		if err != nil {
			panic(err)
		}

		db, err := g.DB(group).Open(node)

		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
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
		case "pgsql":
			to, err = postgres.WithInstance(db, &postgres.Config{})
		case "postgres":
			to, err = postgres.WithInstance(db, &postgres.Config{})
		case "postgresql":
			to, err = postgres.WithInstance(db, &postgres.Config{})
		case "sqlite":
			to, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		}
		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithInstance("gres", from, node.Type, to)
		if err != nil {
			panic(err)
		}
		cur, dirty, err := m.Version()
		if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
			panic(err)
		}
		if dirty {
			panic("db migration error: dirty")
		}

		if cur > 0 && cur == uint(steps) {
			g.Log().Info(ctx, "db migration skip: no change")
			return
		} else if steps > 0 {
			g.Log().Infof(ctx, "db migration start: %d -> %d", cur, uint(steps))
			err = m.Migrate(uint(steps))
		} else {
			g.Log().Info(ctx, "db migration start:", group, node.Type)
			err = m.Up()
		}

		_, _ = m.Close()
		if err == nil {
			g.Log().Info(ctx, "db migration done")
			return
		}
		if errors.Is(err, migrate.ErrNoChange) {
			g.Log().Info(ctx, "db migration done: no change")
			return
		} else {
			g.Log().Info(ctx, "db migration error")
			panic(err)
		}
	}
}
