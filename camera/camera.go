package camera

import "restfulApi/unit/mysql"

var DB = mysql.New()

func New(camera interface{})  {
	println(camera)
}

func init() {
	println("Camera")
}