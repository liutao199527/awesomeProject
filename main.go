package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:123456@tcp(192.168.99.100:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB ,err = gorm.Open("mysql",dsn)

	if err != nil{
		return
	}

	return DB.DB().Ping()
}

// Todomodel
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func main()  {
	// 创建数据库
	// CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()

	if err != nil{
		panic(err)
	}
	fmt.Println("数据库连接成功..........")

	// 数据库绑定模型
	DB.AutoMigrate(&Todo{})
	defer DB.Close() //程序退出关闭数据连接

	r := gin.Default()

	// 告诉gin静态文件哪里去找
	r.Static("/static","static")

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")

	// 解析模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	
	//v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写代码事项，点击提交，会发送请求到这里
			// 1. 从请求中把数据拿出来

			// 2. 存入数据库

		})
		// 查看
		// 1. 查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			
		})
		
		// 2. 查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {
			
		})
		// 修改
		// 修改某一待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}
	

	r.Run(":8080")
}
