package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restfulApi/camera"
	"restfulApi/lib/httprouter"
	"restfulApi/lib/negroni"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newRow0 := camera.UserInfo{}

	fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
	fm.ViewToSource(66)

	/*
		编码：
		func Marshal(v interface{}) ([]byte, error)
		func NewEncoder(w io.Writer) *Encoder
		[func (enc *Encoder) Encode(v interface{}) error
		解码:
		func Unmarshal(data []byte, v interface{}) error
		func NewDecoder(r io.Reader) *Decoder
		func (dec *Decoder) Decode(v interface{}) error

		json类型仅支持string作为关键字，因而转义map时，map[int]T类型会报错(T为任意类型）
		Channel, complex, and function types不能被转义
		不支持循环类型的数据，因为这会导致Marshal死循环
		指针会被转义为其所指向的值
	*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newRow0); err != nil {
		panic(err)
	}
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