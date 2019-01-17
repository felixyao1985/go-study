# go-mysql-curd
目录结构
    
    -camera 数据操作
    -lib    第三方库（可有可无），可以使用GOPATH设置，为了学习方便都放这里了
    -unit   一些自行封装的工具方法
    -app    业务代码
    
学习GO语言 尝试封装mysql 操作
- 操作
<pre><code>
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
    id ,_:= newRow0.Insert()
    fmt.Println("add",id)
    
    newRow1 := camera.UserInfo{
        ID: 524,
        Name: "felix_old",
        Username: "felix_go_test",
    }
    id ,_:= newRow1.Remove()
    fmt.Println("Remove",id)
    println("#################################")
  
    println("##################不用camera View 且会改变对象的值###############")

    fmt.Println("当前 item ID值",newRow0.ID)
    fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
    fm.ViewToSource(66)
    fmt.Println("不用camera View items ID:",newRow0.ID)

    println("##################不用camera Browse###############")
    ListUser := []*camera.UserInfo{}
    camera.DB.BrowseToSource("user","",&ListUser)
    fmt.Println("当前 BrowseToSource item[0] ID值",ListUser[0].Name,ListUser[0].Username) 

</pre></code>
        
httprouter + negroni 来实现 restful 


<pre><code>

    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)


    /*
        negroni.Recovery - 异常（恐慌）恢复中间件
        negroni.Logging - 请求 / 响应 log 日志中间件
        negroni.Static - 静态文件处理中间件，默认目录在 "public" 下.
    */

    n := negroni.Classic()
    n.UseHandler(router)
    n.Run(":3000")  

</pre></code>

grpc
    server.go
    client.go