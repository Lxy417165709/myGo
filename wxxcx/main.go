package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)
var db *sql.DB
var err error
//type User struct {
//
//	Username string	`json:"Username"`
//	Password string	`json:"Password"`
//	Email string	`json:"Email"`
//	Sex string		`json:"Sex"`
//	Phone string	`json:"Phone"`
//}
type Desc struct {
	Desc string
	Finished int32
	Belong string
}

type Item struct {
	Name string
	Lists	[]Desc
	Finished_index	[]int32

}

//type Content struct {
//	Items []Item
//
//}
type Unit struct {
	Username string
	OpFlag int32
	Contents interface{}
}
func checkErr(err error,w http.ResponseWriter){
	if err != nil{
		fmt.Fprint(w,-2)
		return
	}
}
func openDataBase() {
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=Wxxcx sslmode=disable")
	fmt.Println(db)
	if(err != nil){
		fmt.Println("数据库打开发生错误")
		return
	}
}
func updateDesc(desc1 Desc,w http.ResponseWriter) int64{
	stmt,err := db.Prepare("UPDATE userdescs SET finished=$1 where adesc=$2 and belong=$3 ")
	checkErr(err,w);
	res,err := stmt.Exec(desc1.Finished,desc1.Desc,desc1.Belong)
	checkErr(err,w);
	affect,err := res.RowsAffected()
	checkErr(err,w);
	return affect;
}
func insertDesc(desc1 Desc,w http.ResponseWriter) int64{

	stmt, err := db.Prepare("INSERT INTO userdescs(adesc,finished,belong) VALUES($1,$2,$3)")
	//fmt.Println(stmt)
	checkErr(err,w)
	res, err := stmt.Exec(desc1.Desc,desc1.Finished,desc1.Belong)
	checkErr(err,w)
	affect, err := res.RowsAffected()
	checkErr(err,w)
	return affect;
}
func deleteDesc(desc1 Desc,w http.ResponseWriter) int64{

	stmt, err := db.Prepare("DELETE FROM userdescs WHERE adesc=$1 and belong=$2;")
	fmt.Println(stmt)
	checkErr(err,w)
	res, err := stmt.Exec(desc1.Desc,desc1.Belong)
	checkErr(err,w)
	affect, err := res.RowsAffected()
	checkErr(err,w)
	return affect;
}
func deleteItem(desc1 Desc,w http.ResponseWriter) int64{

	stmt, err := db.Prepare("DELETE FROM userdescs WHERE belong=$1;")
	fmt.Println(stmt)
	checkErr(err,w)
	res, err := stmt.Exec(desc1.Belong)
	checkErr(err,w)
	affect, err := res.RowsAffected()
	checkErr(err,w)
	return affect;
}
func selectDesc(w http.ResponseWriter) []Desc{

	rows, err := db.Query("SELECT * FROM userdescs ORDER BY finished")
	checkErr(err,w)
	var res = make([]Desc, 0,1)
	for rows.Next() {
		var desc Desc
		err = rows.Scan(&desc.Desc,&desc.Finished,&desc.Belong)
		checkErr(err,w)
		res = append(res, desc)
	}


	fmt.Println(res)
	return res;
}

func insertItem(item1 Item,w http.ResponseWriter) int64{

	stmt, err := db.Prepare("INSERT INTO useritems(itemname) VALUES($1)")
	fmt.Println(stmt)
	checkErr(err,w)
	res, err := stmt.Exec(item1.Name)
	if(res==nil){
		return 0;
	}
	checkErr(err,w)
	affect, err := res.RowsAffected()
	checkErr(err,w)
	return affect;
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//fmt.Fprintf(w,"hello!!!")
	err:=r.ParseForm()
	checkErr(err,w)
	//selectDesc(w)
	for k := range r.Form {
		var unit Unit
		err := json.Unmarshal([]byte(k),&unit)
		fmt.Println(unit)
		checkErr(err,w)



		if(unit.OpFlag ==1 || unit.OpFlag ==2 || unit.OpFlag ==3){
			var desc Desc
			err := mapstructure.Decode(unit.Contents,&desc)
			checkErr(err,w)
			fmt.Println(desc)
			if(unit.OpFlag ==1){

				fmt.Println(insertDesc(desc,w))
			}
			if(unit.OpFlag ==2){
				fmt.Println(deleteDesc(desc,w))
			}
			if(unit.OpFlag ==3){
				fmt.Println(updateDesc(desc,w))
			}



		}
		if(unit.OpFlag ==5){
			var desc Desc
			err := mapstructure.Decode(unit.Contents,&desc)
			checkErr(err,w)
			fmt.Println(desc)
			fmt.Println(insertDesc(desc,w))
		}
		if(unit.OpFlag ==6){
			var desc Desc
			err := mapstructure.Decode(unit.Contents,&desc)
			checkErr(err,w)
			fmt.Println(desc)
			fmt.Println(deleteItem(desc,w))
		}
	}

}
func get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//fmt.Fprintf(w,"hello!!!")
	checkErr(err,w)
	ans,err := json.Marshal(selectDesc(w))
	checkErr(err,w)
	fmt.Fprintf(w,string(ans))

}
func main() {
	openDataBase()
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/get", get)
	err := http.ListenAndServe(":9091", nil)  //设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
