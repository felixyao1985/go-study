package main

import (
	"fmt"
	"net/http"
	"restfulApi/camera"
	"restfulApi/lib/httprouter"
	"restfulApi/lib/negroni"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}



func main() {

	newRow0 := camera.UserInfo{
		ID: 1,
		Name: "felix_old",
		Username: "felix_go_test",
	}

	List ,_:= newRow0.BrowseAll("")

	fmt.Println("列表:",List[0].Name)
	println("#################################")
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
	println("#################################")

	println("##################不用camera View###############")

	fmt.Println("当前 item ID值",newRow0.ID)
	fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
	fm.ViewToSource(66)
	fmt.Println("不用camera View items ID:",newRow0.ID)




	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)



	//log.Fatal(http.ListenAndServe(":8080", router))

	/*
		negroni.Recovery - 异常（恐慌）恢复中间件
		negroni.Logging - 请求 / 响应 log 日志中间件
		negroni.Static - 静态文件处理中间件，默认目录在 "public" 下.
	*/

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}