package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.POST("/users", CreateUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run()
}

func GetUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, users)
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, user)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
	} else {
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := db.Where("id = ?", id).Delete(&user)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
