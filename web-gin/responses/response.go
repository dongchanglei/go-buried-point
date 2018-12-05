package responses

import (
	"github.com/gin-gonic/gin"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/exceptions"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode , errCode int, data interface{}) {
	g.C.JSON(httpCode,gin.H{
		"code": httpCode,
		"msg":  exceptions.GetMsg(errCode),
		"data": data,
	})

}
