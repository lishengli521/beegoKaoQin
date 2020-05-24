package utils

type JSONStr struct {
	//必须的大写开头
	Code string
	Msg  string
	Data interface{} `json:"user"`//key重命名,最外面是反引号
}

func Succee(value interface{}) interface{} {
	data:=JSONStr{"200", "获取成功",  value}
	return data
}
func Error(error string) interface{} {
	var data=JSONStr{"500",error,nil}
	return data
}
