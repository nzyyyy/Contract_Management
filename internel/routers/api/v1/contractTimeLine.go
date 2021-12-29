package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type TimeLine struct{}

func NewTimeLine() TimeLine {
	return TimeLine{}
}

func (t TimeLine) GetTimeLineById(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := convert.StrTo(c.Param("id")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	line, err := svc.GetTimeLineById(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
	}

	response.ToResponse(line)
}
