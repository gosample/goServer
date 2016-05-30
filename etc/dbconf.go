package etc

import (
	"github.com/astaxie/beego/orm"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:chenweimin@/parkinglots?charset=utf8&loc=Asia%2FShanghai")
}
