package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ryonryon/League/src/controllers"
	"github.com/ryonryon/League/src/models"
	"github.com/ryonryon/League/src/interfaces/handler"
)

var (
	R *gin.Engine
)

func init() {
	R := gin.Default()
	handler.DB.AutoMigrate(models.User{})
	fmt.Print("[RYON-debug] Migrated\n")
	R.GET("/users", controllers.IndexUserController)
	R.GET("/users/:id", controllers.IndexOneUserController)
	R.POST("/users/create", controllers.CreateUserController)
	R.PUT("/users/:id", controllers.UpdateUserController)
	R.DELETE("/users/:id", controllers.DestroyUserController)
}

func main() {
	err := R.Run()
	if err != nil {
		panic(err)
	}
}