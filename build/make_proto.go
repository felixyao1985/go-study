package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

/*
 实践失败=v= 生成的shell是OK的 但是无法用go执行
*/
func GetProtoList(dirpath string) ([]string, error) {
	var ret []string
	dir_list, e := ioutil.ReadDir(dirpath)
	if e != nil {
		fmt.Println("read dir error", e)
	}
	for _, v := range dir_list {
		filenameWithSuffix := path.Base(v.Name())  //获取文件名带后缀
		fileSuffix := path.Ext(filenameWithSuffix) //获取文件后缀

		if fileSuffix == ".proto" {
			ret = append(ret, v.Name())
		}
	}
	return ret, e
}

//func GetCurrentDirectory() string {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
//	if err != nil {
//		log.Fatal(err)
//	}
//	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
//}

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])

	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func writerProto() {

	List, _ := GetProtoList("./src/go-study/proto/")
	fmt.Println(List)

	for _, name := range List {
		command := "d: && cd d:/var/gowork/src/go-study/proto && protoc --go_out=plugins=grpc:. " + name
		params := []string{}
		//执行cmd命令: ls -l
		execCommandOnly(command, params)
	}

}
func execCommandOnly(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	fmt.Println(cmd.Start())
	return true
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println(err2)
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true
}

func main() {
	writerProto()
}
