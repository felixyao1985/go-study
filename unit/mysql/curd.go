package mysql

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"reflect"
	_ "restfulApi/lib/mysql"
	"strings"
)



//注意方法名大写，就是public
func initDB() *sql.DB{
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	con, err := sql.Open("mysql", path)
	if err != nil{
		checkErr(err)
	}
	//设置数据库最大连接数
	con.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	con.SetMaxIdleConns(10)
	//验证连接
	if err := con.Ping(); err != nil{
		checkErr(err)
	}
	fmt.Println("connnect success")

	return con

}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type DB struct {
	con *sql.DB
}

func New() *DB {
	con := initDB()
	return &DB{
		con,
	}
}
/*
	字段对象

	sql.XXX 是go 自带的SQL数据类型，可以处理null 字段。是个坑。
	在rows.Scan 的时候 同事也要接受他的err数据，否则代码虽然不会报错，但是下条数据中为null的字段会继承上一条数据
	IntSave    sql.NullInt64
	StringSave sql.NullString
	FloatSave  sql.NullFloat64
	BoolSave   sql.NullBool

*/
type field struct {
	Name       string
	Tag        string
	Type       string
	Key        bool
	Addr       interface{}
	IntSave    sql.NullInt64
	StringSave sql.NullString
	FloatSave  sql.NullFloat64
	BoolSave   sql.NullBool
}

type _FieldsMap struct {
	dataobj  interface{}
	reftype reflect.Type
	fields  []field
	table   string
	db *sql.DB
}
func (obj *_FieldsMap) GetFields() []field {
	return obj.fields
}

func newFieldsMap(table string, dataobj interface{})(*_FieldsMap, error){

	//reflect.Value.Elem() 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
	elem := reflect.ValueOf(dataobj).Elem()
	//获取反射源的构造对象
	reftype := elem.Type()
	//fmt.Println("reftype",reflect.ValueOf(x))
	//fmt.Println("dataobj",dataobj)
	//fmt.Println("elem",elem)
	//fmt.Println("reftype",reftype)
	//fmt.Println(elem.FieldByName(reftype.Field(0).Name),elem.Field(0))
	/*
		这里都是对 对象源（类似 JAVA class）进行操作，并非对象（new 的对象）
		reftype.NumField() 获取反射源的条目
		reftype.Field()    获取字段
	*/
	var fields []field
	for i, flen := 0, reftype.NumField(); i < flen; i++ {

		var field field
		//获取 class 属性的类型
		field.Type = reftype.Field(i).Type.String()
		field.Name = reftype.Field(i).Name
		field.Tag = reftype.Field(i).Tag.Get("sql")
		//获取对象name 的指针地址
		field.Addr = elem.Field(i).Addr().Interface()

		if(reftype.Field(i).Tag.Get("key") == "") {
			field.Key = false
		}else{
			field.Key = true
		}
		fields = append(fields, field)
	}


	return &_FieldsMap{
		dataobj:  dataobj,
		reftype: reftype,
		fields:  fields,
		table:   table,
	}, nil
}

// NewFieldsMap 生成一个新的对象
func (c *DB) NewFieldsMap(table string, dataobj interface{})(*_FieldsMap, error){
	//fmt.Println("dataobj",dataobj)
	nfm, _ := newFieldsMap(table,dataobj)
	nfm.db = c.con
	return nfm, nil
}

// GetFieldValues 提取结构体中的值数组
func (fds *_FieldsMap) GetFieldValues() []interface{} {

	var values []interface{}
	for i, flen := 0, len(fds.fields); i < flen; i++ {
		values = append(values, fds.GetFieldValue(i))
	}

	return values
}

// GetFieldValue  提取结构体中的值
func (fds *_FieldsMap) GetFieldValue(idx int) interface{} {

	switch fds.fields[idx].Type {
	case "int64":
		return *fds.fields[idx].Addr.(*int64)
	case "string":
		return *fds.fields[idx].Addr.(*string)
	case "float64":
		return *fds.fields[idx].Addr.(*float64)
	case "bool":
		return *fds.fields[idx].Addr.(*bool)
	default:
	}

	return nil
}

// 把要处理的字段 转化成 SQL string
// example:" `field0`, `field1`, `field2`, `field3` "
func (c *_FieldsMap) SQLFieldsStr() string {

	var tagsStr string
	for i, flen := 0, len(c.fields); i < flen; i++ {
		if len(tagsStr) > 0 {
			tagsStr += ", "
		}
		tagsStr += "`"
		tagsStr += c.fields[i].Tag
		tagsStr += "`"
	}
	if len(tagsStr) > 0 {
		tagsStr += " "
		tagsStr = " " + tagsStr
	}

	return tagsStr
}

// GetFieldSaveAddrs 获取结构体内每个"name"的 指针
func (obj *_FieldsMap) GetFieldSaveAddrs() []interface{} {

	var addrs []interface{}
	for i, flen := 0, len(obj.fields); i < flen; i++ {
		addrs = append(addrs, obj.GetFieldSaveAddr(i))
	}

	return addrs
}

// GetFieldSaveAddr 获取结构体内每个"name"的值
func (fds *_FieldsMap) GetFieldSaveAddr(idx int) interface{} {

	switch fds.fields[idx].Type {
	case "int64":
		return &fds.fields[idx].IntSave
	case "string":
		return &fds.fields[idx].StringSave
	case "float64":
		return &fds.fields[idx].FloatSave
	case "bool":
		return &fds.fields[idx].BoolSave
	default:
	}

	return nil
}

// MapBackToObject MAP化对象
func (fds *_FieldsMap) MapBackToObject() interface{} {

	//item := reflect.ValueOf(fds.dataobj).Elem()

	for i, flen := 0, len(fds.fields); i < flen; i++ {
		//fieldInfo := item.Type().Field(i)
		switch fds.fields[i].Type {
		case "int64":
			if fds.fields[i].IntSave.Valid {
				*fds.fields[i].Addr.(*int64) = fds.fields[i].IntSave.Int64
				//item.FieldByName(fieldInfo.Name).SetInt(fds.fields[i].IntSave.Int64)
			}
			break
		case "string":
			if fds.fields[i].StringSave.Valid {
				*fds.fields[i].Addr.(*string) = fds.fields[i].StringSave.String
				//item.FieldByName(fieldInfo.Name).SetString(fds.fields[i].StringSave.String)
			}
			break
		case "float64":
			if fds.fields[i].FloatSave.Valid {
				*fds.fields[i].Addr.(*float64) = fds.fields[i].FloatSave.Float64
				//item.FieldByName(fieldInfo.Name).SetFloat(fds.fields[i].FloatSave.Float64)
			}
			break
		case "bool":
			if fds.fields[i].BoolSave.Valid {
				*fds.fields[i].Addr.(*bool) = fds.fields[i].BoolSave.Bool
				//item.FieldByName(fieldInfo.Name).SetBool(fds.fields[i].BoolSave.Bool)
			}
			break
		default:
		}

	}

	return fds.dataobj
}


/*尝试处理数组*/
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}


func (c *DB) BrowseToSource(table string,sql string, dataobj interface{}){

	/*
		reflect中,最重要的是Value类,只有先获取到一个对象或者变量的Value对象后,我们才可以对这个对象或者变量进行更进一步的分析和处理。
		reflect.ValueOf()方法获取Value对象。
		获取变量的值使用value.Interface()方法，该方法会返回一个value的值，不过类型是interface。

		对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。
		这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
	*/
	elem := reflect.Indirect(reflect.ValueOf(dataobj))
	reftype := elem.Type()

	//获取元素对象的值
	fmt.Println("NewListFieldsMap elem:",elem)
	//获取元素对象的类型
	fmt.Println("NewListFieldsMap reftype:",reftype)

	elemobj := reflect.Indirect(reflect.New(reftype.Elem().Elem())).Addr()

	//获取元素对象的元素类型
	fmt.Println("NewListFieldsMap reftype.Elem():",reftype.Elem())
	//在挖一层
	fmt.Println("NewListFieldsMap reftype.Elem().Elem():",reftype.Elem().Elem())


	//nobj := reflect.New(reftype).Interface()
	obj,_:=newFieldsMap(table, elemobj.Interface())
	fmt.Println("BrowseToSource fieldsMap:",obj)
	con := c.con
	_sql := strings.Join([]string{"SELECT ",obj.SQLFieldsStr()," FROM ",obj.table,sql}, "")

	rows, err := con.Query(_sql)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		nobj := reflect.Indirect(reflect.New(reftype.Elem().Elem())).Addr()
		//nobj := reflect.New(obj.reftype).Interface()
		fieldsMap,err:=newFieldsMap(obj.table, nobj.Interface())
		if err != nil {
			log.Fatal(err)
		}

		err = rows.Scan(fieldsMap.GetFieldSaveAddrs()...)
		//var name string
		if err != nil {
			log.Fatal(err)
		}

		fieldsMap.MapBackToObject()
		elem = reflect.Append(elem, nobj)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	deepCopy(dataobj, elem.Interface())
	fmt.Println("BrowseToSource fieldsMap elem:",elem)
}