package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Customer struct{}

func NewCustomer() Customer {
	return Customer{}
}

// @Summary 创建客户
// @Produce json
// @Param name body string true "客户姓名"
// @Param address body string true "地址"
// @Param tel body string true "电话"
// @Param fax body string true "传真"
// @Param code body string true "邮编"
// @Param bank body string true "银行名称"
// @Param account body string true "银行账户"
// @Success 200 {object} model.Customer "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/customer/create [POST]
func (cus Customer) Create(c *gin.Context) {
	params := service.CreateCustomerRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.CreateCustomer(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除客户
// @Produce json
// @Param id path int true "客户id"
// @Success 200 {object} model.Customer "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/customer/delete/{id} [GET]
func (cus Customer) Delete(c *gin.Context) {
	id, err := convert.StrTo(c.Param("id")).Int()
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.DeleteCustomer(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
