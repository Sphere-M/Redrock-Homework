package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/first")
	db.Ping()
	defer func() {
		if db!=nil{
			db.Close()
		}
	}()
	if err != nil{
		fmt.Println("数据库连接失败！")
	return
	}
	stmt,err:=db.Prepare("insert into people values(default,?,?)")
	defer func() {
		if stmt!=nil{
			stmt.Close()
		}
	}()
	if err!=nil{
		fmt.Println("预处理失败")
		return
	}
	r,err:=stmt.Exec("张三","重庆")
	if err!=nil{
		fmt.Println("sql失败")
		return
	}
	count,err:=r.RowsAffected()
	if err!=nil{
		fmt.Println("结果获取失败")
		return
	}
	if count>0{
		fmt.Println("新增成功")
	}else{
		fmt.Println("新增失败")
	}
    selectDB(db)
	updateDB(db)
	deleteDB(db)
	selectDB(db)
}


func selectDB(db *sql.DB)  {

	stmt, err := db.Query("select * from people;")

	if err != nil {

		log.Fatal(err)

	}

	defer stmt.Close()

	for stmt.Next(){

		var id int

		var name string

		var address string

		err := stmt.Scan(&id,&name,&address)

		if err != nil {

			log.Fatal(err)

		}

		fmt.Println(id,name)

	}



}

func updateDB(db *sql.DB)  {

	stmt, err := db.Prepare("update people set name = '李四' where id = 1")

	if err != nil{

		log.Fatal(err)

	}

	stmt.Exec();

}



func deleteDB(db *sql.DB)  {

	stmt, err := db.Prepare("delete from people where id =5")

	if err != nil{

		log.Fatal(err);

	}

	stmt.Exec();

}