package database

import (
	"RobClassSystemByGo/types"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	userName = "root"
	passWord = "12345678"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "nowUse"
)

var Db *gorm.DB

func ConnectDb() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8mb4"
	dsn := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True"}, "")

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		fmt.Println("open database fail")
		return
	}
	sqlDb, _ := Db.DB()
	// 设置空闲连接数
	// 数量 connections = ((core_count * 2) + effective_spindle_count)
	sqlDb.SetConnMaxIdleTime(10)
	// 最大连接数
	sqlDb.SetMaxOpenConns(100)
	// 连接复用连接时间
	sqlDb.SetConnMaxLifetime(time.Hour)
	fmt.Println("connect success")
}

func InitDb() {
	ConnectDb()

	// 删除原表
	Db.Exec("DROP TABLE courses")
	Db.Exec("DROP TABLE s_courses")
	Db.Exec("DROP TABLE members")

	// 创建新表
	if err := Db.AutoMigrate(&types.Members{}, &types.Courses{}, &types.SCourses{}); err != nil {
		return
	}

	// 初始化新表(添加管理员账户)
	Db.Exec(
		"INSERT INTO members (nickname,username,user_type,password)" +
			"VALUES ('Admin','JudgeAdmin',1,'JudgePassword')",
	)
}
