package main

import (
	"fmt"
	"restfulApi/app"
	"restfulApi/camera"
	"restfulApi/lib/httprouter"
	"restfulApi/lib/negroni"
)

func main() {

	newRow0 := camera.UserInfo{
		ID: 1,
		Name: "felix_old",
		Username: "felix_go_test",
	}

	List ,_:= newRow0.BrowseAll("")

	fmt.Println("列表:",List[0].Name)
	println("#################不会改变对象的值################")
	data ,_:= newRow0.View(55)
	fmt.Println("View data 55 Name",data.Name)
	fmt.Println("View data newRow0 55 Name",newRow0.Name)
	println("#################################")
	data2 ,_:= newRow0.View(44)
	fmt.Println("View data 44 Name",data2.Name)
	fmt.Println("View data newRow0 44 Name",newRow0.Name)
	println("#################################")
	newRow0.Name = "44改名啦"
	newRow0.Update()
	data3 ,_:= newRow0.View(44)
	fmt.Println("View data 44  改名后 Name",data3.Name)
	fmt.Println("View data newRow0 44 改名后 Name",newRow0.Name)
	println("#################################")
	//id ,_:= newRow0.Insert()
	//fmt.Println("add",id)
	newRow1 := camera.UserInfo{
		ID: 524,
		Name: "felix_old",
		Username: "felix_go_test",
	}
	id ,_:= newRow1.Remove()
	fmt.Println("Remove",id)

	println("##################不用camera View 且会改变对象的值###############")

	fmt.Println("当前 item ID值",newRow0.ID)
	fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
	fm.ViewToSource(66)
	fmt.Println("不用camera View items ID:",newRow0.ID)

	println("##################不用camera Browse###############")
	ListUser := []*camera.UserInfo{}
	camera.DB.BrowseToSource("user","",&ListUser)
	fmt.Println("当前 BrowseToSource item[0] ID值",ListUser[0].Name,ListUser[0].Username)

	//router := httprouter.New()
	//router.GET("/", Index)
	//router.GET("/hello/:name", Hello)



	//log.Fatal(http.ListenAndServe(":8080", router))

	/*
		negroni.Recovery - 异常（恐慌）恢复中间件
		negroni.Logging - 请求 / 响应 log 日志中间件
		negroni.Static - 静态文件处理中间件，默认目录在 "public" 下.
	*/

	n := negroni.Classic()
	n.UseHandler(NewRouter())
	n.Run(":3000")
}

func NewRouter() *httprouter.Router {

	router := httprouter.New()
	routes := app.GetRoutes()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}