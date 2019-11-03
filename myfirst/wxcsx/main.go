package  main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
type user struct {
	Nickname string
	City string
	Country string
	Avatarurl string
	Time string
	Rgstime string
	Gender string
	Oid string

}
func checkErr(err error){
	if err != nil {
		panic(err)
	}
}
func isExist(User user) bool{
	stmt,err := db.Prepare("SELECT * FROM wxxcxuser WHERE oid = $1")
	checkErr(err);
	res,err := stmt.Exec(User.Oid);
	checkErr(err)
	affect,err := res.RowsAffected()
	checkErr(err)
	if affect != 0 {
		return true
	} else {
		return false
	}


}
func insertUser(User user) int64{

	if !isExist(User) {
		return updateUser(User);
	} else {
		stmt,err := db.Prepare("UPDATE wxxcxuser SET nickname=$1,city=$2,country=$3,gender=$4,avatarurl=$5,time=$6 where oid=$7")
		checkErr(err)
		res,err := stmt.Exec(User.Nickname,User.City,User.Country,User.Gender,User.Avatarurl,User.Time,User.Oid)
		checkErr(err);
		affect,err := res.RowsAffected()
		checkErr(err)
		return affect
	}



}
func deleteUser(User user) int64{
	stmt,err := db.Prepare("DELETE FROM wxxcxuser where oid=$1")
	checkErr(err)
	res,err := stmt.Exec(User.Oid)
	checkErr(err)
	affect,err := res.RowsAffected()
	checkErr(err)
	return affect
}
func updateUser(User user) int64{
	stmt,err := db.Prepare("INSERT INTO wxxcxuser(nickname,city,country,gender,avatarurl,time,oid,rgstime) values($1,$2,$3,$4,$5,$6,$7,$8)")
	checkErr(err);
	res,err := stmt.Exec(User.Nickname,User.City,User.Country,User.Gender,User.Avatarurl,User.Time,User.Oid,User.Time)
	checkErr(err);
	affect,err := res.RowsAffected()
	checkErr(err);
	return affect;
}


func login(writer http.ResponseWriter, requester *http.Request){
	writer.Header().Set("Access-Control-Allow-Origin","*")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	writer.Header().Set("content-type", "application/json")             //返回数据格式是json
	err := requester.ParseForm()
	checkErr(err)
	var User user
	User.Oid = requester.Form["oid"][0];
	User.Nickname = requester.Form["nickname"][0];
	User.Gender = requester.Form["gender"][0];
	User.City = requester.Form["city"][0];
	User.Country = requester.Form["country"][0];
	User.Time = requester.Form["time"][0];
	User.Avatarurl = requester.Form["avatarurl"][0];
	fmt.Println(insertUser(User))
}

func openDatabase(){
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=User sslmode=disable")
	checkErr(err)
}
func main(){
	openDatabase()
	http.HandleFunc("/login",login);
	err := http.ListenAndServe(":8888",nil);
	if err != nil {
		log.Fatal("err",err)
	}

}