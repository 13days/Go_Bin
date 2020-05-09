package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

// 自定义表名
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_users"
	} else {
		return "users"
	}
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/gogin?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil{
		fmt.Println("数据库连接失败",err)
	}
	fmt.Println("数据库连接成功...")

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	// 创建表
	//db.CreateTable(&User{})

	p := time.Now()
	str := "1"
	user := User{Name: "11", Age: sql.NullInt64{18,true},Birthday:&p, MemberNumber : &str, Email:"213"}

	f1 := db.NewRecord(user) // => 主键为空返回`true`
	fmt.Println("NewRecord:",f1)

	db.Create(&user)
	f1 = db.NewRecord(user) // => 创建`user`后返回`false`
	fmt.Println("NewRecord:",f1)

}
