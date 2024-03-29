package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Log struct{}

func NewLog() Log {
	return Log{}
}

func (l Log) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := convert.StrTo(c.Param("id")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.DeleteLog(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(gin.H{})
}
func (l Log) GetAll(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	list, err := svc.GetLogList()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	response.ToResponse(list)
}
