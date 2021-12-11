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

// @Summary  分配权限
// @Produce  json
// @Param    userId       body      int            true  "用户id"
// @Param    roleId       body      int            true  "角色id"
// @Param    description  body      string         true  "描述信息"
// @Success  200          {object}  model.Right    "成功"
// @Failure  400          {object}  errcode.Error  "请求错误"
// @Failure  500          {object}  errcode.Error  "内部错误"
// @Router   /management/v1/right/create [POST]
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
		response.ToErrorResponse(errcode.ErrorCreateRightFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary  删除权限
// @Produce  json
// @Param    id   path      int            true  "用户id"
// @Success  200  {object}  model.Right    "成功"
// @Failure  400  {object}  errcode.Error  "请求错误"
// @Failure  500  {object}  errcode.Error  "内部错误"
// @Router   /management/v1/right/delete/{id} [GET]
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
		response.ToErrorResponse(errcode.ErrorDeleteRightFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary  更新权限
// @Produce  json
// @Param    userId       body      int            true  "用户id"
// @Param    roleId       body      int            true  "角色id"
// @Param    description  body      string         true  "描述信息"
// @Success  200          {object}  model.Right    "成功"
// @Failure  400          {object}  errcode.Error  "请求错误"
// @Failure  500          {object}  errcode.Error  "内部错误"
// @Router   /management/v1/right/update [POST]
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
		response.ToErrorResponse(errcode.ErrorUpdateRightFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
