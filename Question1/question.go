package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//数据库名为“question”，表名为“question”，字段为  （id int ， name varchar ，price int)
//选取方式为随机挑选
func main(){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/question")
	db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败！")
	}else {
		fmt.Println("数据库连接成功！")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	add(1,"小面",6,db)
	add(2,"饭团",7,db)
	add(3,"香菇滑鸡",12,db)
	add(4,"小炒肉",15,db)
	add(5,"黄焖鸡",16,db)
	add(6,"冒菜",18,db)
	selectDB(db)
}

func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from question order by rand() limit 1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var id int
		var name string
		var price int
		err := stmt.Scan(&id,&name,&price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name,price)
	}
}

func add(a int,str string ,b int,db *sql.DB) {
	stmt, err := db.Prepare("insert into question values(?,?,?)")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("预处理失败")
		return
	}
	r, err := stmt.Exec(a,str, b)
	if err != nil {
		fmt.Println("sql失败")
		return
	}
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("结果获取失败")
		return
	}
	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}
}