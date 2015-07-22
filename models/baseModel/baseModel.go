package baseModel

// import (
// 	"github.com/astaxie/beego/orm"
// )
import "github.com/astaxie/beego"

type BaseModel struct {
	DB_Url string
    // orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8", 30)
}

func (this *BaseModel) init() {
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	// db_url  := beego.AppConfig.String("db_host")
	// db_port := beego.AppConfig.String("db_port")
	db_name := beego.AppConfig.String("db_name")
	this.DB_Url = db_user+":"+db_pass+"@/"+db_name+"?charset=utf8"
}

// type Users struct {
// 	Id    int  `orm:"auto"`
// 	Account string 
// 	Pw  string 
// 	Email  string 
// }

// type Categories struct {
// 	Id    int  `orm:"auto"`
// }
// func init() {

// 	user := beego.AppConfig.String("mysqluser")
// 	pass := beego.AppConfig.String("mysqlpass")
// 	urls := beego.AppConfig.String("mysqlurls")
// 	db := beego.AppConfig.String("mysqldb")
// }

