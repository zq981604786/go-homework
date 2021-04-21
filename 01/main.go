package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

type Tags struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *Tags) Find() (*Tags, error) {
	err = DB.QueryRow("select id,name from tags where id = ?", 2).Scan(&t.Id, &t.Name)
	return t, err
}

func init() {
	DB, err = sql.Open("mysql", "root:123456@(127.0.0.1:3306)/testdb")
	err = DB.Ping()
	if err != nil {
		fmt.Println("Database connection failed")
		panic(err)
	}
}

func main() {
	defer DB.Close()
	fmt.Printf("login err:%v\n", Login())
	_, err2 := FindList()
	fmt.Printf("find err:%v\n", err2)
}

// 登录不存在应该返回错误
func Login() error {
	tag := Tags{Id: 2}
	_, err1 := tag.Find()
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			return errors.New("账号或密码不存在")
		}
	}
	return nil
}

// 查询不存在应该返回空的
func FindList() ([]Tags, error) {
	tag := Tags{Id: 2}
	var tags []Tags
	_, err1 := tag.Find()
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			return tags, nil
		}
	}
	return tags, nil
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.msg
}
