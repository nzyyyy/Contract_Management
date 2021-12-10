package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Right struct{}

func NewRight() Right {
	return Right{}
}

func (r Right) Create(c *gin.Context) {
	params := service.CreateOrUpdateRightRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.CreateRight(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (r Right) Delete(c *gin.Context) {
	id, err := convert.StrTo(c.Param("id")).Int()
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.DeleteRight(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (r Right) Update(c *gin.Context) {
	params := service.CreateOrUpdateRightRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.UpdateRight(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
