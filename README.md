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

var DB  *gorm.DB
```

```go
dsn := strings.Join([]string{userName, ":", passWord, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=True"}, "")

var err error
DB , err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    PrepareStmt: true,
})
```

#### 初始化设置

```go
sqlDb, _ := DB .DB()
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
DB .Exec("DROP TABLE courses")
DB .Exec("DROP TABLE s_courses")
DB .Exec("DROP TABLE members")

// 创建新表
if err := DB .AutoMigrate(&types.Members{}, &types.Courses{}, &types.SCourses{}); err != nil {
    return
}

// 初始化新表(添加管理员账户)
DB .Exec(
    "INSERT INTO members (nickname,username,user_type,password)" +
    "VALUES ('Admin','JudgeAdmin',1,'JudgePassword')",
)
```



### redis

#### 连接redis

```go
var RDB *redis.Client

RDB = redis.NewClient(&redis.Options{
    Addr:     "127.0.0.1:6379",
    Password: "",  // no password set
    DB:       0,   // use default DB
    PoolSize: 100, // 连接池大小
})

// 利用根Context创建一个父Context
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

_, err := RDB.Ping(ctx).Result()

if err != nil {
    fmt.Println("open redis fail")
    return
}
```

#### 删除redis缓存

```go
// 删除 redis 缓存
res, err := RDB.FlushDB(ctx).Result()
if err != nil {
    panic(err)
}
fmt.Println("delete redis:", res)
```

## auth

### Login

#### 获取请求

```go
var request types.LoginRequest
if err := c.ShouldBind(&request); err != nil {
    log.Println(err)
    return
}
```



#### 提取数据行，提取失败返回密码错误

```go
var line types.Members
database.DB.Model(&types.Member{}).Where(&request).First(&line)
if line == (types.Members{}) {
    c.JSON(200, types.LoginResponse{
        Code: types.WrongPassword,
    })
    return
}
```

##### gorm

* Model() : 指定数据行
* Where() : 条件
* Find() : 查找第一个

#### cookie，redis设置

##### 获取sessionKey避免明文写入

```go
// 获取唯一标识符 uuid 作为该数据行的键
sessionKey := uuid.NewV4().String()
```

##### 写入cookie

```go
c.SetCookie("camp-session", sessionKey, 3600, "/", "", false, false)
```

###### 参数说明

1. 第一个参数为 cookie 名
2. 第二个参数为 cookie 值
3. 第三个参数为 cookie 有效时长，当 cookie 存在的时间超过设定时间时，cookie 就会失效，它就不再是我们有效的 cookie；设为-1是移除cookie
4. 第四个参数为 cookie 所在的目录
5. 第五个为所在域，表示我们的 cookie 作用范围
6. 第六个表示是否只能通过 https 访问
7. 第七个表示 cookie 是否可以通过 js代码进行操作。



##### redis设置

* 将 sessionKey 作为键 UserID 和 UserType 作为值写入redis

```go
ctx := context.Background()
data := map[string]interface{}{
    "UserID":   line.UserID,
    "UserType": fmt.Sprint(line.UserType),
}
//log.Println(data)
if err := database.RDB.HMSet(ctx, sessionKey, data).Err(); err != nil {
    log.Fatal(err)
    return
}
```



##### session 与 cookie

1. Cookie以文本文件格式存储在浏览器中，而session存储在服务端它存储了限制数据量
2. cookie的存储限制了数据量，只允许4KB，而session是无限量的
3. 可以轻松访问cookie值但是我们无法轻松访问会话值，因此它更安全
4. 设置cookie时间可以使cookie过期。但是使用session-destory（），我们将会销毁会话。



### Logout

* 从 cookie 中提取 sessionKey， 删除对应的 cookie 和 redis

### WhoAmI

#### 获取user_id

1. 从 cookie 中提取 sessionKey  `sessionKey, err := c.Cookie("camp-session")`
2. 从 redis 中提取 user_id 

```go
rets, err := database.RDB.HMGet(context.Background(), sessionKey, "UserID").Result()
if err != nil {
    log.Println(err)
}
userID := rets[0].(string)
```

#### 从数据库提取 member 并返回

```go
var member types.Member
database.DB.Table("members").Where("user_id = ?", userID).Find(&member)
c.JSON(200, types.WhoAmIResponse{
   Code: types.OK,
   Data: member,
})
```

## member

### Create

#### 权限校验

1. 从 cookie 中提取  sessionKey
2. 从 redis 中提取 UserType
3. 检验 UserType

#### 数据行提取(包括软删除)

* 提取 

```go
database.DB.Model(types.Members{}).Unscoped().Where("username = ?", request.Username).Find(&line)
```

* 状态判断
	* 不存在数据行 : `line == (types.Members{})`
	* 软删除 :` line.Deleted.Valid`

#### 操作

* 创建 : `database.DB.Model(types.Members{}).Create(&request)`
* 查询:  `database.DB.Model(types.Members{}).Where("username=?", request.Username).Find(&member)`
* 更新 : `database.DB.Model(types.Members{}).Where("user_id=?", request.UserID).Update("Nickname", request.Nickname)`
* 删除 : `database.DB.Model(types.Members{}).Where("user_id=?", request.UserID).Delete(&types.Members{})`
