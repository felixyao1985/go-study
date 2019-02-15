# go-mysql-curd
目录结构
    
    -camera 数据操作
    -lib    第三方库（可有可无），可以使用GOPATH设置，为了学习方便都放这里了
    -unit   一些自行封装的工具方法
    -app    业务代码目录
        - dir 各模块目录
            - index.go 业务代码 
            - route.go 路由控制，会自动编译生成routePath文件内的内容
    
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
#GO 部署

设置环境变量来切换编译环境
"="前后不能有空格

 - 1 设置环境变量 并打印显示，最后运行main.go

   set GOBUILD=test&&set GOBUILD&&go run main.go （貌似编译后无法使用，因为通过set设置的环境变量只在当前窗口有效）

 - 2 os.Args[1] 运行时带参数设置

- log.sh # 实时查看日志
- build.sh # 构建
- run.sh # 启动某一次编译版本
- start.sh # 启动最新版本，并且备份之前前一次运行的版本
- shutdown.sh # 停止
- rollback.sh # 回滚到上一版本

#GO grpc
  
  #### https://github.com/felixyao1985/go-study/blob/master/grpc.md 
 