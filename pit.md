#遇到的一些坑随手记

- 包引用要以gopath来引用，不要使用文件相对目录来引，否则打包会找不到

- 自动生成的文件要先生成 再打包。打包并不会生成这些文件

- 要避免包重复引用的情况

- main 包中的不同的文件的代码不能相互调用，其他包可以。

- 交叉便宜的坑

     - golang.org/x/sys/unix 可能被墙可以去 https://github.com/golang/sys 下载
     - 必须在cmd 状态下编译，bash下不起作用
     
交叉编译

    set GOARCH=amd64
    
    set GOOS=linux
    
    go build xx.go
     