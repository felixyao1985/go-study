package main

import (
	"../unit/mysql"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func initDB() *sql.DB{
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{mysql.UserName, ":", mysql.Password, "@tcp(",mysql.IP, ":", mysql.PORT, ")/", "information_schema", "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	con, err := sql.Open("mysql", path)
	if err != nil{
		log.Fatal(err)
	}
	//设置数据库最大连接数
	con.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	con.SetMaxIdleConns(10)
	//验证连接
	if err := con.Ping(); err != nil{
		log.Fatal(err)
	}
	fmt.Println("connnect success")

	return con

}
var DB = initDB()
//var DB = mysql.New()

func getTablelist()([]string){

	sql := "select table_name from `tables` where `table_schema`='vsimanage' and `table_type`='base table'"

	var list = []string{}
	//DB.BrowseToSource("user",sql,&list)
	rows, _ := DB.Query(sql)

	for rows.Next() {
	   var table_name string

		err:= rows.Scan(&table_name)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, table_name)
	}
	return list
}

func getColumnCameraList(table string)([]string)  {

	sql := "select column_name,COLUMN_TYPE,COLUMN_KEY from information_schema.columns where table_schema='vsimanage' and table_name='"+ table +"'"

	var list = []string{}
	//DB.BrowseToSource("user",sql,&list)
	rows, _ := DB.Query(sql)

	var camera_string string
	for rows.Next() {
		var column_name string
		var COLUMN_TYPE string
		var COLUMN_KEY string

		err:= rows.Scan(&column_name,&COLUMN_TYPE,&COLUMN_KEY)
		if err != nil {
			log.Fatal(err)
		}

		//ID int64  `sql:"id" key:"PRIMARY"`
		//Name string  `sql:"name"`
		//Username string    `sql:"username"`

		switch true {
			case strings.Contains(COLUMN_TYPE, "int"):
				camera_string = strings.ToUpper(column_name) + " int64 `sql:\"" + column_name + "\""

			case strings.Contains(COLUMN_TYPE, "varchar"),strings.Contains(COLUMN_TYPE, "text"):
				camera_string = strings.ToUpper(column_name) + " string `sql:\"" + column_name + "\""

			case strings.Contains(COLUMN_TYPE, "float"):
				camera_string = strings.ToUpper(column_name) + " float64 `sql:\"" + column_name + "\""
			case strings.Contains(COLUMN_TYPE, "bool"):
				camera_string = strings.ToUpper(column_name) + " bool `sql:\"" + column_name + "\""
			default:
				camera_string = strings.ToUpper(column_name) + " string `sql:\"" + column_name + "\""
		}
		if COLUMN_KEY == "PRI" {
			camera_string += " key:\"PRIMARY\"`"
		}else{
			camera_string += "`"
		}


		list = append(list, camera_string)
	}
	return list
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name,content string) {
	ext,_ := PathExists(name)

	if ext {
		fmt.Println(name,"文件存在:")
	}else{
		file, error := os.Create(name)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(file)
		file.Close()

		data :=  []byte(content)
		if ioutil.WriteFile(name,data,0644) == nil {
			fmt.Println(name,"camera 写入文件成功:")
		}
	}

}

func writerCamera(table string,ColumnList []string) {

	tableName := strings.ToUpper(table)
	name := "./src/vsiapi/camera/"+tableName+".go"
	tableColumn := "table"+tableName
	content := `
package camera

import (
	"strconv"
	"strings"
)

type ` + tableName + ` struct {`
	for i, length := 0, len(ColumnList); i < length; i++ {
		content += ColumnList[i]+`
	`
	}
	content +=`}

var `+ tableColumn +`   = "`+ table +`"

func (c *`+ tableName +`) Browse(sql string,row int,start int) ([]`+ tableName +`,error){

	_sql := strings.Join([]string{sql," limt ",strconv.Itoa(start)," , ",strconv.Itoa(row)}, "")
	objs , _:= c.BrowseAll(_sql)
	return objs, nil
}

func (c *`+ tableName +`) BrowseAll(sql string) ([]`+ tableName +`,error){

	fm, _ := DB.NewFieldsMap(`+ tableColumn +`, c)
	items := fm.Browse(sql)

	var objs []`+ tableName +`
	for i, olen := 0, len(items); i < olen; i++ {
		objs = append(objs, *items[i].(*`+ tableName +`))
	}
	return objs, nil
}

func (c *`+ tableName +`) View(id int) (`+ tableName +`,error){

	fm, _ := DB.NewFieldsMap(`+ tableColumn +`, c)
	items := fm.View(id)
	return *items.(*`+ tableName +`), nil
}

func (c *`+ tableName +`) Insert() (int64,error){
	fm, _ := DB.NewFieldsMap(`+ tableColumn +`, c)
	return fm.Insert()
}

func (c *`+ tableName +`) Update() (int64,error){
	fm, _ := DB.NewFieldsMap(`+ tableColumn +`, c)
	return fm.Update()
}

func (c *`+ tableName +`) Remove() (int64,error){
	fm, _ := DB.NewFieldsMap(`+ tableColumn +`, c)
	return fm.Remove()
}`

	WriteWithIoutil(name,content)
}


func main() {
	tablelist := getTablelist()

	for i, length := 0, len(tablelist); i < length; i++ {
		table_name := tablelist[i]
		fmt.Println(table_name)

		columnCameraList := getColumnCameraList(table_name)

		writerCamera(table_name,columnCameraList)
	}
}
