## 项目说明
```shell
使用goload创建项目
# 创建文件夹
mkdir ginchat
# 进入文件夹
cd ginchat
# 初始化一个mod类型的项目，名字叫做ginchat
go mod init ginchat

# 将mod清空，只保留自己使用的（初始化时，肯定都没有）
go mod tiny

# 下载orm（调用数据库框架）依赖包
# 不使用这个了~，下面那个更好 go get github.com/jinzhu/gorm
# 下载
go get gorm.io/gorm

# 下载sqlite驱动
go get gorm.io/driver/sqlite
# 下载mysql驱动
```