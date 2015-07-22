package models

import (
	// bm "web/models/baseModel"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
    "time"
)

type Posts struct {
	// bm.BaseModel
	Id    int  `orm:"auto"`
	Title  string 
	Content  string 
	Created_at time.Time `orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time `orm:"auto_now_add;type(datetime)"`

}

type Posts_Json struct {
	Id int `json:"id"`
	Title  string `json:"title"`
	Content  string `json:"content"`
}
// //--------------------------------------------------
// type Post3 struct {
// 	User_id int32 `from:"user_id"`
// 	Status  int32 `from:"status"`
// 	Title  string `from:"title"`
// 	Content  string `from:"content"`
// }
// u := model.Post3{}
//     if err := this.ParseForm(&u); err != nil {
//         //handle error
//     }
//     fmt.Printf("%+v\n", u)
// //--------------------------------------------------

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterModel(new(Posts))
    orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8", 30)
}

func (this *Posts) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return 
}

/*
func (this *Posts) Add(data Posts_Json) (id int64, err error) {
	this.User_id, _ = strconv.Atoi(data.User_id)
	this.Category_id, _ = strconv.Atoi(data.Category_id)
	this.Status, _ = strconv.Atoi(data.Status)
	this.Title = data.Title
	this.Content = data.Content
	o := orm.NewOrm()
	id, err = o.Insert(this)
	fmt.Println(err)
	return 
}
*/

func (this *Posts) GetAllList() ([]Posts, int64) {
	var data []Posts
	count, _ := orm.NewOrm().QueryTable("posts").OrderBy("-id").All(&data)
	return data, count
}

func (this *Posts) GetById() (Posts, error) {
	o := orm.NewOrm()
	data := Posts{}
	data.Id = this.Id

	err := o.Read(&data)
	// fmt.Println(err)
	// if err == orm.ErrNoRows {
	//     fmt.Println("No result found.")
	// } else if err == orm.ErrMissPK {
	//     fmt.Println("No primary key found.")
	// } else {
	//     fmt.Println(user.Title)
	// }
	return data, err
}

func (this *Posts) Update() (int64, error) {
	o := orm.NewOrm()
	ret, err := o.Update(this)
	return ret, err
}

func (this *Posts) Delete() (int64, error) {
	o := orm.NewOrm()
	ret, err := o.Delete(&Posts{Id:this.Id}) 
	return ret, err
}
