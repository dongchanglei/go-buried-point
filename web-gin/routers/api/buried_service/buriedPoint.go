package buried_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/service"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/responses"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/exceptions"
	"gitlab.kanche.com/kkche/go-buried-point/web-gin/request"
)

func GetBuriedPoints(c *gin.Context) {
	appG := responses.Gin{C: c}

	buriedService := service.BuriedPoint{
		PageNum:  0,
		PageSize: 10,
	}

	total, err := buriedService.Count()
	if err != nil {
		appG.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}

	burieds, err := buriedService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = burieds
	data["total"] = total

	appG.Response(http.StatusOK, exceptions.SUCCESS, data)
}


func AddBuriedPoint(c *gin.Context) {
	var (
		appG = responses.Gin{C: c}
		form buried_request.AddBuriedJson
	)
	appG.C.Bind(&form)
	buriedService := service.BuriedPoint{
		SourceId:   form.SourceId,
		UserId:     form.UserId,
		CrmUserId:  form.CrmUserId,
		GroupId:    form.GroupId,
		CrmGroupId: form.CrmGroupId,
		GroupLink:  form.GroupLink,
		SourceType: form.SourceType,
	}
	if err := buriedService.Add(); err != nil {
		appG.Response(http.StatusOK, exceptions.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, exceptions.SUCCESS, nil)
}
