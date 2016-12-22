package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data/data.db")
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(Project), new(Build))
	verbose, _ := beego.AppConfig.Bool("Verbose")
	orm.RunSyncdb("default", false, verbose)
}