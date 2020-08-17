package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

var db *gorm.DB
var err error

type User struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err = gorm.Open("mysql", "root:646233@tcp(127.0.0.1:3306)/test_gorm")
	if err != nil {
		log.Fatal("db connect error")
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/users", index)
	r.GET("/users/:id", show)
	r.POST("/users", store)
	r.PUT("/users/:id", update)
	r.DELETE("/users/:id", destroy)
	_ = r.Run()
}

func index(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(200, users)
}

func show(c *gin.Context) {
	id , _ := strconv.Atoi(c.Params.ByName("id"))
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}
	c.JSON(200, user)
}

func store(c *gin.Context) {
	var user User
	_ = c.BindJSON(&user)
	db.Create(&user)
	c.JSON(200, user)
}

func update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		user.Name = "had changed"
		db.Save(&user)
		c.JSON(200, user)
	}
}

func destroy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var user User
	db.First(&user, id)
	if user.ID == 0 {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else {
		_ = c.BindJSON(&user)
		db.Delete(&user)
		c.JSON(200, gin.H{"message": "delete success"})
	}
}