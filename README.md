# Only demo

## main.go
* (14, 5) // TODO 判断是否需要重置数据库
* (31, 4) // TODO 功能完善
* (32, 4) // TODO 压测
* (33, 4) // TODO 并发优化
* (34, 4) // TODO Mysql优化
* (35, 4) // TODO Redis优化

## database
### mysql.go
* (50, 5) // TODO 优化写法
### redis.go
* (39, 5) // TODO 从数据库中提取剩余 Cap

## function
### auth
#### Login.go
* (37, 5) // TODO 优化 UserType 写法
### book
#### Book.go
* (15, 5) // TODO 压测
* (43, 5) // TODO 选课与否， 添加失败返还redis数据
### course
#### Create.go
* (13, 5) // TODO 判断课程存在与否
#### Get.go
* (12, 5) // TODO 判断课程存在与否
### member
#### Check.go
* (17, 5) // TODO 修改
* (32, 5) // TODO 参数校验
#### Delete.go
* (12, 5) // TODO 避免删除自身
* (13, 5) // TODO 相关删除

## types
### book.go
* (14, 4) // TODO 添加异步检测，直到完成写入再读取
