package main

import (
	"fmt"
	"io/ioutil"
)

func GetDirList(dirpath string) ([]string, error) {
	var ret []string
	dir_list, e := ioutil.ReadDir(dirpath)
	if e != nil {
		fmt.Println("read dir error")
	}
	for _, v := range dir_list {
		if v.IsDir() {
			ret = append(ret, v.Name())
		}
	}
	return ret, e
}

func writerRoute() {
	dirList, _ := GetDirList("./src/go-study/app/")
	fmt.Println(dirList)
	name := "./src/go-study/app/routes.go"
	content := `
package app

import (`
	for _, name := range dirList {
		content += `"` + "go-study/app/" + name + `"
`
	}

	content += `
	"../router"
)



func mergeRoutes()  []router.Route{
	routes := []router.Route{}
`
	content += "routes = merge(routes"
	for _, name := range dirList {
		content += ", " + name + ".Routes"
	}
	content += ")"
	content += `
		return routes
	}`

	WriteWithIoutil(name, content)
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		//fmt.Println("写入文件成功:",content)
		fmt.Println("routes 写入文件成功:")
	}
}

func main() {
	writerRoute()
}
