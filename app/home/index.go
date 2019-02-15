package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go-study/camera"
	"go-study/lib/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	newRow0 := camera.UserInfo{}

	fm, _ := camera.DB.NewFieldsMap("user", &newRow0)
	fm.ViewToSource(66)

	/*
		编码：
		func Marshal(v interface{}) ([]byte, error)
		func NewEncoder(w io.Writer) *Encoder
		[func (enc *Encoder) Encode(v interface{}) error
		解码:
		func Unmarshal(data []byte, v interface{}) error
		func NewDecoder(r io.Reader) *Decoder
		func (dec *Decoder) Decode(v interface{}) error

		json类型仅支持string作为关键字，因而转义map时，map[int]T类型会报错(T为任意类型）
		Channel, complex, and function types不能被转义
		不支持循环类型的数据，因为这会导致Marshal死循环
		指针会被转义为其所指向的值
	*/
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newRow0); err != nil {
		panic(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
