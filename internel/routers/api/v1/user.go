package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (u User) UpdateRole(c *gin.Context) {
	response := app.NewResponse(c)
	userId, err := convert.StrTo(c.Param("userId")).Int()
	roleId, err := convert.StrTo(c.Param("roleId")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.RemoveUserRole(userId, roleId)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(gin.H{})
}
