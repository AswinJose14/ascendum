package main

import (
	"ascendum/gameEngine"
	"ascendum/models"
	"ascendum/userEngine"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r := setUpRoutes()
		if err := r.Run(":8080"); err != nil {
			panic(err)
		}
	}()
	time.Sleep(2 * time.Second)
	userEng := userEngine.UserEngine{
		Count: 1000,
	}
	userEng.GenerateUsers()

	go func() {
		time.Sleep(5 * time.Second)
		gameEngine.AnnonceWinners()
	}()

	select {}

}

func setUpRoutes() *gin.Engine {
	fmt.Println("SettingUp Rputes")
	r := gin.Default()
	r.POST("/submit", SubmitHandler)
	return r
}

func SubmitHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request object",
		})
	}
	gameEngine := gameEngine.GetGameEngine()
	gameEngine.EvaluvateAnswer(&user)
}
