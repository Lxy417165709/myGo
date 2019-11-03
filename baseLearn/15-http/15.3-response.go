package main

func main() {

}

/*

	1. 响应头 设置为 返回文件
	w.Header().Set("Content-Type","application/octet-stream")
	w.Header().Set("Content-Disposition","attachment;filename="+filename)

	2. 响应头 设置为 返回json
	w.Header().Set("Content-Type","application/json;charset=utf-8")

 */