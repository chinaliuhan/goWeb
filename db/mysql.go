package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}

func addUserInfo() {
	//连接MySQL, 第一个参数是数据库类型,第二个是数据连接方式,里面的()括号是必须的,没有密码就留空,格式为:用户:密码@tcp(链接:端口)/数据库名称?charset=utf8
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)

	//拼装SQL
	stmt, err := db.Prepare("insert userinfo set username=?, department=?,created=?")
	checkErr(err)

	//填充数据,执行操作
	res, err := stmt.Exec("nihao", "开发部", "2017-01-01")
	checkErr(err)

	//返回插入的ID
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("插入数据返回的ID:")
	fmt.Println(id)
}

func modifyUserInfo(uid string) {
	//连接MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//拼装SQL
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	//修改数据
	res, err := stmt.Exec("shaKaLaKa", uid)
	checkErr(err)
	//获取受影响行数
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("修改数据的受影响行数:")
	fmt.Println(affect)
}

func getUserInfo() {
	//连接MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//拼装SQL
	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	//遍历结果集
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("查询数据的结果集:")
		fmt.Println(uid, username, department, created)
	}
}

func delUserInfo(uid string) {
	//连接MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//拼装SQL
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	//执行SQL
	res, err := stmt.Exec(uid)
	//获取受影响行数
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("删除数据时返回的受影响行数")
	fmt.Println(affect)
	db.Close()
}

func transactionDel(uid string) {
	//连接MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//开启事务
	tran, err := db.Begin()
	checkErr(err)
	res, err := tran.Exec("delete from userinfo where uid=?", uid)
	if err != nil {
		tran.Rollback()
		panic(err)
	}
	//获取受影响行数
	affect, err := res.RowsAffected()
	checkErr(err)

	//提交事务
	tran.Commit()
	//tran.Rollback()

	fmt.Println(affect)
}
func main() {

	uid := "3"
	addUserInfo()
	modifyUserInfo(uid)
	getUserInfo()
	//delUserInfo(uid)
	transactionDel(uid)

}
