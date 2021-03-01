package routers

import (
	"github.com/gin-gonic/gin"
	"puhai/server/controllers"
)

func LoadRouter(e *gin.Engine) {
	e.GET("/getData", controllers.ViewData)
}