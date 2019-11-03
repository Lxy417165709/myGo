package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type user struct {

	Username string	`json:"Username"`
	Password string	`json:"Password"`
	Email string	`json:"Email"`
	Sex string		`json:"Sex"`
	Phone string	`json:"Phone"`
}

var db *sql.DB
var err error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func openDataBase() {
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=User sslmode=disable")
	fmt.Println(db)
	checkErr(err)
}
func usnIsExist(usn string) int64{
	stmt, err := db.Prepare("SELECT * FROM userdetail WHERE  username=$1" )
	checkErr(err)
	res, err := stmt.Exec(usn)
	affect, err := res.RowsAffected()
	checkErr(err)
	return affect;
}
func getUser(usn string) user{
	rows,err:=db.Query("SELECT * FROM PUBLIC.userdetail where username = $1",usn)
	checkErr(err)
	u:= user{}
	for rows.Next(){
		var id int
		err = rows.Scan(&id,&u.Username,&u.Password,&u.Sex,&u.Email,&u.Phone)
		checkErr(err)
	}
	return u
}
func insertData(usn string,psw string,sex string,email string,phone string) int64{

	if usnIsExist(usn) > 0{
		return -1
	}
	stmt, err := db.Prepare("INSERT INTO userdetail(username,psword,sex,email,phone) VALUES($1,$2,$3,$4,$5) RETURNING uid")
	checkErr(err)
	res, err := stmt.Exec(usn, psw, sex, email, phone)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	return affect;
}
func register(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入内容
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	err := r.ParseForm()
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	for k := range r.Form {
		fmt.Println(k)
		var User user

		err := json.Unmarshal([]byte(k),&User)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		fmt.Fprintln(w,insertData(User.Username,User.Password,User.Sex,User.Email,User.Phone))

	}

}
func search(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入内容
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json


	err := r.ParseForm()
	checkErr(err)
	for k := range r.Form {
		var User user

		err := json.Unmarshal([]byte(k),&User)
		checkErr(err)
		var wtUser user = getUser(User.Username)
		ans,err := json.Marshal(wtUser)
		checkErr(err);
		fmt.Fprintln(w,string(ans))
	}

}
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	err := r.ParseForm()
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	for k := range r.Form {
		var User user

		err := json.Unmarshal([]byte(k),&User)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		if(User.Username == "root"){
			if(User.Password == "root"){
				fmt.Fprintln(w,1)
				return
			} else{
				fmt.Fprintln(w,0)
				return
			}
		}
		var realyPsw string = getUser(User.Username).Password

		if realyPsw == "" {
			fmt.Fprintln(w,-1)
			return
		}

		if User.Password == realyPsw{
			fmt.Fprintln(w,1)
			return
		} else{
			fmt.Fprintln(w,0)
			return
		}
	}

}
func getAllUser() []user {
	rows,err:=db.Query("SELECT * FROM PUBLIC.userdetail")
	checkErr(err)
	allUser := []user{}

	for rows.Next(){
		u:= user{}
		var id int
		err = rows.Scan(&id,&u.Username,&u.Password,&u.Sex,&u.Email,&u.Phone)
		checkErr(err)
		allUser = append(allUser, u)
	}
	return allUser

}
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	err := r.ParseForm()
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	for k := range r.Form {
		var User user
		err := json.Unmarshal([]byte(k),&User)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		stmt,err := db.Prepare("DELETE FROM PUBLIC.userdetail where username=$1")
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		res,err:= stmt.Exec(User.Username)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		affect,err :=res.RowsAffected()
		fmt.Fprintln(w,affect)

	}
}
func changePsw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	err := r.ParseForm()
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	for k := range r.Form {
		var User user
		err := json.Unmarshal([]byte(k), &User)
		if (err != nil) {
			fmt.Fprint(w, -2)
			return
		}
		var wtUser user = getUser(User.Username)
		if (wtUser.Username == "") {
			fmt.Fprint(w, -1)
			return
		}
		if (User.Email == wtUser.Email || User.Phone == wtUser.Phone) {
			stmt, err := db.Prepare("UPDATE  PUBLIC.userdetail SET psword=$2 where username=$1")
			if (err != nil) {
				fmt.Fprint(w, -2)
				return
			}
			res, err := stmt.Exec(User.Username, User.Password)
			if (err != nil) {
				fmt.Fprint(w, -2)
				return
			}
			affect, err := res.RowsAffected()
			if (err != nil) {
				fmt.Fprint(w, -2)
				return
			}
			fmt.Fprintln(w, affect)
			return
		} else {
			fmt.Fprint(w, 0)
			return
		}
	}
}
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	err := r.ParseForm()
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	for k := range r.Form {
		var User user
		err := json.Unmarshal([]byte(k),&User)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		stmt,err := db.Prepare("UPDATE  PUBLIC.userdetail SET psword=$2,sex=$3,phone=$4,email=$5 where username=$1")
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		res,err:= stmt.Exec(User.Username,User.Password,User.Sex,User.Phone,User.Email)
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		affect,err :=res.RowsAffected()
		if(err != nil){
			fmt.Fprint(w,-2)
			return
		}
		fmt.Fprintln(w,affect)

	}
}
func userInformation(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	ans,err := json.Marshal(getAllUser())
	if(err != nil){
		fmt.Fprint(w,-2)
		return
	}
	fmt.Println(getAllUser())
	fmt.Fprintln(w,string(ans))
}
func test(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	fmt.Fprintln(w,15)
}
func sayHello(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	fmt.Fprintf(w,"师兄师姐好!!!")
}

func main() {
	openDataBase()
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/test", test)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/search", search)
	http.HandleFunc("/update", update)
	http.HandleFunc("/changePsw", changePsw)
	http.HandleFunc("/userInformation", userInformation)
	err := http.ListenAndServe(":9091", nil)  //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
