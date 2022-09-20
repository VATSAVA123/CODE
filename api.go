package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name     string `json: "Namee"`
	Id       string `json: "Idd"`
	Password string `json: "Passwordd"`
	City     string `json: "Cityy"`
}

var userinfo = []user{
	{"Vatsav", "514", "vat", "uppal"},
	{"preetham", "544", "pre", "secbad"},
	{"kiran", "520", "kir", "warangal"},
	{"sharath", "525", "sha", "ap"},
}

func getallusers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, userinfo)
}

func getuserbyid(c *gin.Context) {
	id := c.Param("id")
	for _, value := range userinfo {
		if id == value.Id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
}

func main1() {
	router := gin.Default()
	router.GET("/getallusers", getallusers)
	router.GET("/getuserbyid/:id", getuserbyid)

	router.Run("localhost:8080")

}
