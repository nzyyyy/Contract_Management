package routers

import (
	_ "contract_management/docs"
	v1 "contract_management/internel/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	contract := v1.NewContract()
	contractProcess := v1.NewContractProcess()
	customer := v1.NewCustomer()
	right := v1.NewRight()
	role := v1.NewRole()
	apiv1 := r.Group("management/v1")
	{
		apiv1.POST("/contract/create", contract.Create)
		apiv1.POST("/contract/delete", contract.Delete)
		apiv1.POST("/contract/update", contract.Update)

		apiv1.POST("/contractProcess/update", contractProcess.Update)
		apiv1.GET("/contractProcess/delete/:id", contractProcess.Delete)
		apiv1.POST("/contractProcess/create", contractProcess.Create)

		apiv1.POST("/customer/create", customer.Create)
		apiv1.GET("/customer/delete/:id", customer.Delete)

		apiv1.POST("/right/update", right.Update)
		apiv1.GET("/right/delete/:id", right.Delete)
		apiv1.POST("/right/create", right.Create)

		apiv1.POST("/role/create", role.Create)
		apiv1.GET("/role/delete/:id", role.Delete)
	}
	return r
}
