package db

import (
	"fmt"
	"time"

	"github.com/bshadmehr98/cheshm/models"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB orm.Ormer
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase(
		"default",
		"sqlite3",
		"./db.sqlite3",
	)

	orm.DefaultTimeLoc = time.UTC

	// register model
	orm.RegisterModel(new(models.User))

	// Error.
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}

	DB = orm.NewOrm()
}
