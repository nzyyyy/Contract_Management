package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Function struct{}

func NewFunction() Function {
	return Function{}
}

func (f Function) GetList(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	list, err := svc.GetFunctionList()
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(list)
}
