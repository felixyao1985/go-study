package mysql

import (
	"log"
	"reflect"
	"strings"
)

func (obj *_FieldsMap) Browse(sql string) []interface{} {
	con := obj.db
	_sql := strings.Join([]string{"SELECT ", obj.SQLFieldsStr(), " FROM ", obj.table, sql}, "")

	rows, err := con.Query(_sql)
	if err != nil {
		log.Fatal(err)
	}
	var objs []interface{}
	//fmt.Println(obj.reftype)
	for rows.Next() {
		nobj := reflect.New(obj.reftype).Interface()
		fieldsMap, err := newFieldsMap(obj.table, nobj)
		if err != nil {
			log.Fatal(err)
		}
		//是个坑。
		//在rows.Scan 的时候 同事也要接受他的err数据，否则代码虽然不会报错，但是下条数据中为null的字段会继承上一条数据
		err = rows.Scan(fieldsMap.GetFieldSaveAddrs()...)
		//var name string
		if err != nil {
			log.Fatal(err)
		}

		fieldsMap.MapBackToObject()
		objs = append(objs, nobj)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return objs

}

func (obj *_FieldsMap) View(id int) interface{} {
	con := obj.db
	_sql := strings.Join([]string{"SELECT ", obj.SQLFieldsStr(), " FROM ", obj.table, " where id = ? "}, "")

	row := con.QueryRow(_sql, id)
	/*
		生成了一个新的对象
	*/
	nobj := reflect.New(obj.reftype).Interface()
	fieldsMap, err := newFieldsMap(obj.table, nobj)
	if err != nil {
		log.Fatal(err)
	}

	err = row.Scan(fieldsMap.GetFieldSaveAddrs()...)
	//var name string
	if err != nil {
		log.Fatal(err)
	}

	fieldsMap.MapBackToObject()

	return nobj

}

// 新增
func (obj *_FieldsMap) Insert() (int64, error) {
	con := obj.db
	var vs string
	var tagsStr string
	var values []interface{}
	for i, flen := 0, len(obj.fields); i < flen; i++ {

		//祛除主键
		if !obj.fields[i].Key {
			if len(vs) > 0 {
				vs += ", "
			}
			vs += "?"

			if len(tagsStr) > 0 {
				tagsStr += ", "
			}
			tagsStr += "`"
			tagsStr += obj.fields[i].Tag
			tagsStr += "`"

			values = append(values, obj.GetFieldValue(i))
		}
	}

	if len(tagsStr) > 0 {
		tagsStr += " "
		tagsStr = " " + tagsStr
	}

	sqlstr := "INSERT INTO `" + obj.table + "` (" + tagsStr + ") " +
		"VALUES (" + vs + ")"
	//fmt.Println(sqlstr)
	tx, _ := con.Begin()
	res, err := tx.Exec(sqlstr, values...)
	if err != nil {
		log.Fatal("Exec fail", err)
	}
	//将事务提交
	tx.Commit()

	//获得上一个插入自增的id
	return res.LastInsertId()
}

// 更新
func (obj *_FieldsMap) Update() (int64, error) {
	con := obj.db
	var tagsStr string
	var whereSql string
	var keyVal int64 = 0
	var values []interface{}
	for i, flen := 0, len(obj.fields); i < flen; i++ {

		//祛除主键
		if obj.fields[i].Key {
			keyVal = obj.GetFieldValue(i).(int64)
			whereSql = " where `" + obj.fields[i].Tag + "` = ? "
		} else {
			if len(tagsStr) > 0 {
				tagsStr += ", "
			}
			tagsStr += "`"
			tagsStr += obj.fields[i].Tag
			tagsStr += "`"
			tagsStr += " = ?"

			values = append(values, obj.GetFieldValue(i))
		}
	}

	if keyVal == 0 {
		return 0, nil
	}

	values = append(values, keyVal)

	if len(tagsStr) > 0 {
		tagsStr += " "
		tagsStr = " " + tagsStr
	}

	sqlstr := "UPDATE `" + obj.table + "` SET " + tagsStr + whereSql
	//fmt.Println(sqlstr)
	tx, _ := con.Begin()
	res, err := tx.Exec(sqlstr, values...)
	if err != nil {
		log.Fatal("Exec fail", err)
	}
	//将事务提交
	tx.Commit()

	//获得上一个插入自增的id
	return res.LastInsertId()
}

// 删除
func (obj *_FieldsMap) Remove() (int64, error) {
	con := obj.db
	var whereSql string
	var keyVal int64 = 0
	for i, flen := 0, len(obj.fields); i < flen; i++ {
		//祛除主键
		if obj.fields[i].Key {
			keyVal = obj.GetFieldValue(i).(int64)
			whereSql = " where `" + obj.fields[i].Tag + "` = ? "
		}
	}

	if keyVal == 0 {
		return 0, nil
	}

	sqlstr := "DELETE FROM `" + obj.table + "`  " + whereSql
	//fmt.Println(sqlstr)
	tx, _ := con.Begin()
	res, err := tx.Exec(sqlstr, keyVal)
	if err != nil {
		log.Fatal("Exec fail", err)
	}

	//将事务提交
	tx.Commit()

	//获得上一个插入自增的id
	return res.RowsAffected()
}

/*无返回操作*/
func (obj *_FieldsMap) ViewToSource(id int) {
	con := obj.db
	_sql := strings.Join([]string{"SELECT ", obj.SQLFieldsStr(), " FROM ", obj.table, " where id = ? "}, "")

	row := con.QueryRow(_sql, id)
	err := row.Scan(obj.GetFieldSaveAddrs()...)
	//var name string
	if err != nil {
		log.Fatal(err)
	}
	obj.MapBackToObject()
}

/*
暂且一放。稍后收拾
研究构造方法返回
*/
//func (obj *_FieldsMap) BrowseToSource(sql string)  {
//	println("BrowseToSource : ",obj.fields)
//	con := obj.db
//	_sql := strings.Join([]string{"SELECT ",obj.SQLFieldsStr()," FROM ",obj.table,sql}, "")
//
//	rows, err := con.Query(_sql)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var objs []interface{}
//	fmt.Println(obj.reftype)
//	for rows.Next() {
//		nobj := reflect.New(obj.reftype).Interface()
//		fieldsMap,err:=newFieldsMap(obj.table, nobj)
//		if err != nil {
//			log.Fatal(err)
//		}
//		//是个坑。
//		//在rows.Scan 的时候 同事也要接受他的err数据，否则代码虽然不会报错，但是下条数据中为null的字段会继承上一条数据
//		err = rows.Scan(fieldsMap.GetFieldSaveAddrs()...)
//		//var name string
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fieldsMap.MapBackToObject()
//		objs = append(objs, nobj)
//	}
//
//	if err := rows.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//}
