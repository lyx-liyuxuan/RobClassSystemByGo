# RobClassSystemByGo

## 结构

* .
	* database
		* mysql.go : Mysql 相关设置
		* redis.go : Redis 相关设置
	* types
		* types.go : 		通用类型
		* database.go : Mysql类型
		* login.go : 登陆模块类型
		* member.go : 成员模块类型
		* course.go : 课程模块类型
		* book.go : 抢课模块类型

## 数据库设计

### 表格

* 成员表 Members
* 课程表 Courses
* 选课表 SCourses

### Members

|   名称   |     类型      |               限制               |    含义    |
| :------: | :-----------: | :------------------------------: | :--------: |
|    id    |    BIGINT     | PRIMARYKEY; <br />AUTO_INCREMENT |   用户ID   |
| nickname |  VARCHAR(32)  |             NOT NULL             |    别名    |
| username |  VARCHAR(32)  |             NOT NULL             |   用户名   |
|   type   | UserType(int) |             NOT NULL             |  用户类型  |
| password |   CHAR(32)    |             NOT NULL             |    密码    |
| deleted  |               |                                  | 软删除标记 |

#### VARCHAR 与 CHAR 的区别



|  类型   |        长度        |                    存储方式                     |     存储容量     |                        优势                         |
| :-----: | :----------------: | :---------------------------------------------: | :--------------: | :-------------------------------------------------: |
| VARCHAR |      可变长度      | 额外使用字节记录长度<br />5.0以后长度按字符展示 | 受其他列数据影响 |                      节约内存                       |
|  CHAR   | 固定长度(空格填充) |                 固定字节数存储                  |     255字节      | 短字符串<br />长度相近字符串(MD5)<br />常变动字符串 |



#### 软删除

* 含义 : 记录不会被数据库。但 GORM 会将 `DeletedAt` 置为当前时间， 并且不能再通过普通的查询方法找到该记录。

* 定义 ： gorm.DeletedAt类型(struct)
* 查找 : `db.Unscoped().Where("age = 20").Find(&users)`
* 永久删除 : `db.Unscoped().Delete(&order)`

### courses

|    名称    |    类型     |              限制               |   含义   |
| :--------: | :---------: | :-----------------------------: | :------: |
| course_id  |   BIGINT    | PRIMARYKE; <br />AUTO_INCREMENT |  课程id  |
|    name    | VARCHAR(32) |            NOT NULL             |  课程名  |
|    cap     |     INT     |            NOT NULL             | 课程总量 |
| teacher_id | VARCHAR(32) |                                 |  教师id  |

* 教师暂不考虑落库检测

### s_courses

|    名称    |  类型  |               限制               |  含义  |
| :--------: | :----: | :------------------------------: | :----: |
|   sc_id    | BIGINT | PRIMARYKE;  <br />AUTO_INCREMENT | 选课号 |
| course_id  | BIGINT |  FOREIGN KEY courses.course_id   | 课程id |
| student_id | BIGINT |   FOREIGN KEY members.user_id    | 学生id |

## 初始化

### mysql

#### 创建连接

```go
const (
	userName = "root"
	passWord = "12345678"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "nowUse"
)

var Db *gorm.DB
```

```go
dsn := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True"}, "")

var err error
Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    PrepareStmt: true,
})
```

#### 初始化设置

```go
sqlDb, _ := Db.DB()
	// 设置空闲连接数
	sqlDb.SetConnMaxIdleTime(10)
	// 最大连接数
	sqlDb.SetMaxOpenConns(100)
	// 连接复用连接时间
	sqlDb.SetConnMaxLifetime(time.Hour)
```



#### 重建表(可选)

```go
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
```



## redis

### 连接redis

