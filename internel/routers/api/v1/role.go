package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Role struct{}

func NewRole() Role {
	return Role{}
}

// @Summary  创建角色
// @Produce  json
// @Param    name         body      int            true  "角色名称"
// @Param    description  body      string         true  "描述信息"
// @Param    functions    body      int            true  "功能操作"
// @Success  200          {object}  model.Role     "成功"
// @Failure  400          {object}  errcode.Error  "请求错误"
// @Failure  500          {object}  errcode.Error  "内部错误"
// @Router   /management/v1/right/create [POST]
func (r Role) Create(c *gin.Context) {
	params := service.CreateRoleRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.CreateRole(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateRoleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary  删除角色
// @Produce  json
// @Param    id   path      int            true  "角色id"
// @Success  200  {object}  model.Role     "成功"
// @Failure  400  {object}  errcode.Error  "请求错误"
// @Failure  500  {object}  errcode.Error  "内部错误"
// @Router   /management/v1/role/delete/{id} [GET]
func (r Role) Delete(c *gin.Context) {
	id, err := convert.StrTo(c.Param("id")).Int()
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.DeleteRole(id)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteRoleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
