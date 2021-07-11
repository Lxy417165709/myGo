package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("baseLearn/18-file/src/index.html") //注意路径
	t.Execute(w, nil)

}

// 上传文件
func upload(w http.ResponseWriter, r *http.Request) {
	// 设置返回格式
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	// 获取表单中的文件
	file, header, _ := r.FormFile("file")

	// 解析文件为二进制数据
	b, _ := ioutil.ReadAll(file)

	// 获取文件名
	fileName := header.Filename

	// 写到本地文件
	err := ioutil.WriteFile(fileName, b, 0777)
	if err != nil {
		fmt.Fprint(w, 0)
	} else {
		fmt.Fprint(w, 1)
	}
}

// 下载文件
func download(w http.ResponseWriter, r *http.Request) {

	filename := r.FormValue("filename")
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprint(w, "文件下载失败 ")
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	_, err = w.Write(f)
	if err != nil {
		fmt.Fprint(w, "文件下载失败")
		return
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", download)
	http.ListenAndServe(":9999", nil)

}
