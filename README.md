# go-mysql-curd
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
          
            println("##################不用camera View###############")
          
            fmt.Println("当前 item ID值",newRow0.ID)
            fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
            fm.ViewToSource(66)
            fmt.Println("不用camera View items ID:",newRow0.ID)      
        
        </pre></code>