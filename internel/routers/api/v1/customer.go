package v1

import (
	"contract_management/internel/model"
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Customer struct{}

func NewCustomer() Customer {
	return Customer{}
}

// @Summary  创建客户
// @Produce  json
// @Param    name     body      string          true  "客户姓名"
// @Param    address  body      string          true  "地址"
// @Param    tel      body      string          true  "电话"
// @Param    fax      body      string          true  "传真"
// @Param    code     body      string          true  "邮编"
// @Param    bank     body      string          true  "银行名称"
// @Param    account  body      string          true  "银行账户"
// @Param    userId  body      int          true  "操作人id"
// @Success  200      {object}  model.Customer  "成功"
// @Failure  400      {object}  errcode.Error   "请求错误"
// @Failure  500      {object}  errcode.Error   "内部错误"
// @Router   /management/v1/customer/create [POST]
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
		response.ToErrorResponse(errcode.ErrorCreateCustomerFail)
		return
	}
	log := service.CreateLogRequest{UserId: params.UserId, Content: "创建用户:" + params.Name}
	go svc.CreateLog(&log)
	response.ToResponse(gin.H{})
}

// @Summary  删除客户
// @Produce  json
// @Param    customerId   body      int             true  "客户id"
// @Param    UserId   body      int             true  "操作人id"
// @Success  200  {object}  model.Customer  "成功"
// @Failure  400  {object}  errcode.Error   "请求错误"
// @Failure  500  {object}  errcode.Error   "内部错误"
// @Router   /management/v1/customer/delete/{id} [GET]
func (cus Customer) Delete(c *gin.Context) {
	params := service.DeleteCustomerRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.DeleteCustomer(params.CustomerId)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	log := service.CreateLogRequest{UserId: params.UserId, Content: "删除用户,id为:" + strconv.Itoa(params.CustomerId)}
	go svc.CreateLog(&log)
	response.ToResponse(gin.H{})
}
func (cus Customer) ListOfPart(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New()
	list, err := svc.GetCustomerPartList()
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(list)
}
func (cus Customer) GetCustomerById(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := convert.StrTo(c.Param("id")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	customer, err := svc.GetCustomerById(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(customer)
}
func (cus Customer) UpdateCustomer(c *gin.Context) {
	params := model.Customer{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.UpdateCustomer(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateCustomerFail)
		return
	}
	response.ToResponse(gin.H{})
}
