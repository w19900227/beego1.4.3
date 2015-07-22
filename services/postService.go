package services

import (
	bs "beego1.4.3/services/baseService"
	"beego1.4.3/models"
	"strconv"
	"encoding/json"
)
import "fmt"

type PostService struct {
	bs.BaseService
	post_model models.Posts
}

func (this *PostService) GetAllList() ([]models.Posts, int64) {
	data, count := this.post_model.GetAllList()
	return data, count
}

func (this *PostService) GetById(id string) (data models.Posts) {
	fmt.Println(id)
	this.post_model.Id, _ = strconv.Atoi(id)
	data, _ = this.post_model.GetById()
	return 
}

func (this *PostService) Add(data []byte) {
	var json_data models.Posts_Json
	json.Unmarshal(data, &json_data)
	
	this.post_model.Title = json_data.Title
	this.post_model.Content = json_data.Content

	this.post_model.Add()
}

func (this *PostService) Update(data []byte, id string) {
	var json_data models.Posts_Json
	json.Unmarshal(data, &json_data)

	this.post_model.Id, _ = strconv.Atoi(id)
	this.post_model.Title = json_data.Title
	this.post_model.Content = json_data.Content

	this.post_model.Update()
}

func (this *PostService) Delete(data []byte) {
	var ids struct {
		Id []string
	}

	json.Unmarshal(data, &ids)

	for _, id := range ids.Id {
		this.post_model.Id, _ = strconv.Atoi(id)
        this.post_model.Delete()
    }
}
/*
func (this *PostService) Add(data []byte) {
	var json_data models.Posts_Json
	var posts models.Posts
	json.Unmarshal(data, &json_data)
	posts.Add(json_data)
}
*/

// type PostController struct {
// 	bc.BaseController
// }

// func (this *PostController) Get() {
// 	this.Data["Website"] = "aaaaaaaaaaaaa.me"
// 	this.Data["Email"] = "aaaaaaaaaaa@gmail.com"

// 	this.TplNames = "user/index.tpl"
// }

// func (this *PostController) Create() {
// 	this.Data["Website"] = "2222222222.me"
// 	this.Data["Email"] = "22222222222@gmail.com"

// 	this.TplNames = "user/create.tpl"
// }

// func (this *PostController) Post() {

// 	// a := (map[string]interface{})(args[0].(map[string]interface{}))
// 	// job := ArgsToJob(&a)
// 	// var ob model.Post
// 	// fmt.Println(this.Ctx.Input.RequestBody)
// 	// json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
// 	// fmt.Println(ob)


// 	 // var i int

//   //        i, _ = strconv.Atoi("123")

//   //        fmt.Println(i)

	
// 	// m := model.Post{Title: "test_title", Body: "test_body"}
// 	// var s int
// 	// m := model.Post{}
// 	// m.User_id = 1
// 	// status := this.Input().Get("status")
// 	// s, _ = strconv.Atoi(status)
// 	// m.Status  = s
// 	// m.Title   = this.Input().Get("title")
// 	// m.Content = this.Input().Get("content")
// 	// m.Add()
//  //    this.Data["json"] = map[string]string{"data": "ok"}
//  //    this.ServeJson()
	
// 	// this.Data["Website"] = "sssssss.me"
// 	// this.Data["Email"] = "ssssss@gmail.com"
// 	// this.TplNames = "index.tpl"



//     // mystruct := { "data":"ok" }
// 	// bc.AjaxResponse(200, "ok")
//     // this.Data["json"] = &mystruct
//     // this.ServeJsonp()
// }

// func (this *PostController) Update() {
// 	// m := model.Post{Title: "test_title", Body: "test_body"}
// 	m := model.Users{}
// 	// m.Id          = this.GetString("id")
// 	m.Account     = this.Input().Get("account")
// 	m.Pw          = this.Input().Get("pw")
// 	m.Email       = this.Input().Get("email")
// 	// m.Category_id = this.Input().Get("category_id")
// 	m.Update()
	
// 	// this.Data["Website"] = "beego.me"
// 	// this.Data["Email"] = "astaxie@gmail.com"
// 	// this.TplNames = "index.tpl"
// }
