package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/routers/api/buried_service"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv := r.Group("/api/v1")

	//apiv.Use()
	{
		apiv.POST("/hello", buried_service.AddBuriedPoint)
		apiv.GET("/hello", buried_service.GetBuriedPoints)
	}

	return r
}
