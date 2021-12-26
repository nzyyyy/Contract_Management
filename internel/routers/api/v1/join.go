package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Join struct {
}

func NewJoin() Join {
	return Join{}
}

func (j Join) GetAllProcessList(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	list, err := svc.GetAllProcessList()
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(list)
}
func (j Join) GetProcessListByUser(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := convert.StrTo(c.Param("id")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	list, err := svc.GetProcessListByUser(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(list)
}

func (j Join) GetUserWithRole(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	list, err := svc.GetUserWithRole()
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(list)
}
