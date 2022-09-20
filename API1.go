package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type users struct {
	Empid    string `json:"EMPID" gorm:"primary_key;"`
	Name     string `json:"NAME"`
	Password string `json:"PASSWORD"`
	City     string `json:"CITY"`
}

var DB *gorm.DB

func databaseconn() {
	connect := "host=localhost user=postgres password=1234 dbname=DB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Print("success", db)
	db.AutoMigrate(&users{})
	DB = db

}
func show(c *gin.Context) {
	var showallusers []users
	DB.Find(&showallusers)
	c.IndentedJSON(http.StatusOK, gin.H{"message": showallusers})
}

func getbyid(c *gin.Context) {
	var user users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"BindJSON-Error": err.Error()})
		return
	}

	if err := DB.First(&user, user.Empid).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"DB-Error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusFound, gin.H{"User found :)": user})
}
func create(c *gin.Context) {
	var newuser users

	if err := c.ShouldBindJSON(&newuser); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error-BindJSON": err.Error()})
		return
	}

	err := DB.Create(&newuser).Error
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"New user created :)": newuser})
}
func delete(c *gin.Context) {
	var userToDelete users

	if err := c.ShouldBindJSON(&userToDelete); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := DB.Delete(userToDelete).Error
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete successful."})
}

func main() {

	databaseconn()
	router := gin.Default()
	router.GET("/showallusers", show)
	router.GET("/userbyid", getbyid)
	router.POST("/createnewuser", create)
	router.DELETE("/deleteuser", delete)
	router.Run()
}
