package controllers

import (
	bc "beego1.4.3/controllers/baseController"
	"beego1.4.3/services"
)
import "fmt"
// import "strconv"
// import "encoding/json"


type PostController struct {
	bc.BaseController
	post_service services.PostService
}

func (this *PostController) Test() {
	fmt.Println("controller Test")
	// s := services.PostService{}
	// s.Test()
}

func (this *PostController) Get() {
	data, _ := this.post_service.GetAllList()

	this.Data["Data"] = data
	this.Layout = "common/base.tpl"
	this.TplNames = "post/index.tpl"
}

func (this *PostController) Show() {
	id := this.Ctx.Input.Param(":id")
	data := this.post_service.GetById(id)

	this.Data["Data"] = data
	this.Layout = "common/base.tpl"
	this.TplNames = "post/show.tpl"
}

func (this *PostController) Edit() {
	id := this.Ctx.Input.Param(":id")
	data := this.post_service.GetById(id)


	this.Data["Data"] = data
	this.Layout = "common/base.tpl"
	this.TplNames = "post/edit.tpl"
}

func (this *PostController) Create() {
	this.TplNames = "post/create.tpl"
	this.Layout = "common/base.tpl"
}

func (this *PostController) Post() {
	// var ob models.Posts_Json
	// json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	this.post_service.Add(this.Ctx.Input.RequestBody)
	result := map[string]interface{} {}
	this.AjaxResponse("ok", "", "", result)

	// a := this.Ctx.Input.RequestBody.(byte)
	// fmt.Println(a)
	// val := a.(type)

	// fmt.Println(this.Ctx.Input.RequestBody["title"])
    // var ob model.Post_Json
    // json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
    // fmt.Printf("%+v\n", ob)


	// m := model.Post{Title: "test_title", Body: "test_body"}
	// var s int
	// m := model.Post{}
	// m.User_id = 1
	// status := this.Input().Get("status")
	// s, _ = strconv.Atoi(status)
	// m.Status  = s
	// m.Title   = this.Input().Get("title")
	// m.Content = this.Input().Get("content")
	// m.Add()
 //    this.Data["json"] = map[string]string{"data": "ok"}
 //    this.ServeJson()
	
	// this.Data["Website"] = "sssssss.me"
	// this.Data["Email"] = "ssssss@gmail.com"
	// this.TplNames = "index.tpl"



    // mystruct := { "data":"ok" }
	// bc.AjaxResponse(200, "ok")
    // this.Data["json"] = &mystruct
    // this.ServeJson()
}

func (this *PostController) Update() {
	// result := map[string]interface{} {
	// 	"status" : "ok",
	// 	"erron" : "",
	// 	"errmsg" : "",
	// 	"data" : "",
	// }

	// defer func() {
	// 	this.Data["json"] = &result
	// 	this.ServeJson()
	// }()


	id := this.Ctx.Input.Param(":id")
	this.post_service.Update(this.Ctx.Input.RequestBody, id)

	result := map[string]interface{} {}
	this.AjaxResponse("ok1", "", "", result)
}

func (this *PostController) Delete() {

	this.post_service.Delete(this.Ctx.Input.RequestBody)

	result := map[string]interface{} {}
	this.AjaxResponse("ok2", "", "", result)
}

