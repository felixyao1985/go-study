package camera

import (
	"reflect"
	"strconv"
	"strings"
)

type UserInfo struct {
	ID int64  `sql:"id" key:"PRIMARY"`
	Name string  `sql:"name"`
	Username string    `sql:"username"`
}

const table  string = "user"

func (c *UserInfo) List(items []interface{}) ([]UserInfo,error){
	var objs []UserInfo
	for i, olen := 0, len(items); i < olen; i++ {
		objs = append(objs, *items[i].(*UserInfo))
	}
	return objs, nil
}

func (c *UserInfo) Browse(sql string,row int,start int) ([]UserInfo,error){

	_sql := strings.Join([]string{sql," limt ",strconv.Itoa(start)," , ",strconv.Itoa(row)}, "")
	objs , _:= c.BrowseAll(_sql)
	return objs, nil
}

func (c *UserInfo) BrowseAll(sql string) ([]UserInfo,error){

	fm, _ := DB.NewFieldsMap(table, c)
	items := fm.Browse(sql)

	var objs []UserInfo
	for i, olen := 0, len(items); i < olen; i++ {
		objs = append(objs, *items[i].(*UserInfo))
	}
	return objs, nil
}

func (c *UserInfo) View(id int) (UserInfo,error){

	fm, _ := DB.NewFieldsMap(table, c)
	items := fm.View(id)

	elem := reflect.ValueOf(c).Elem()
	reftype := elem.Type()

	elem2 := reflect.ValueOf(items).Elem()

	for i, flen := 0, reftype.NumField(); i < flen; i++ {
		Addr :=elem.Field(i).Addr().Interface()
		Fieldtype := elem.Field(i).Type().String()

		switch Fieldtype{
		case "int64":
			*Addr.(*int64) = elem2.Field(i).Int()
			break
		case "string":
			*Addr.(*string) = elem2.Field(i).String()
			break
		case "float64":
			*Addr.(*float64) = elem2.Field(i).Float()
			break
		case "bool":
			*Addr.(*bool) = elem2.Field(i).Bool()
			break
		default:
		}

	}

	return *items.(*UserInfo), nil
}

func (c *UserInfo) Insert() (int64,error){
	fm, _ := DB.NewFieldsMap(table, c)
	return fm.Insert()
}

func (c *UserInfo) Update() (int64,error){
	fm, _ := DB.NewFieldsMap(table, c)
	return fm.Update()
}

func (c *UserInfo) Remove() (int64,error){
	fm, _ := DB.NewFieldsMap(table, c)
	return fm.Remove()
}