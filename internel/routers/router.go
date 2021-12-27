package routers

import (
	_ "contract_management/docs"
	"contract_management/internel/middleware"
	v1 "contract_management/internel/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	contract := v1.NewContract()
	contractProcess := v1.NewContractProcess()
	customer := v1.NewCustomer()
	right := v1.NewRight()
	role := v1.NewRole()
	log := v1.NewLog()
	join := v1.NewJoin()
	apiv1 := r.Group("management/v1")
	{
		apiv1.POST("/contract/create", contract.Create)
		apiv1.POST("/contract/delete", contract.Delete)
		apiv1.POST("/contract/update", contract.Update)
		apiv1.GET("/contract/allList", contract.GetAllList)
		apiv1.GET("/contract/listByUser/:id", contract.GetListByUser)
		apiv1.GET("/contract/contractById/:id", contract.GetContractById)

		apiv1.POST("/contractProcess/update", contractProcess.Update)
		apiv1.GET("/contractProcess/delete/:id", contractProcess.Delete)
		apiv1.POST("/contractProcess/create", contractProcess.Create)
		apiv1.GET("/contractProcess/comment/:id/:type", contractProcess.GetContractComment)

		apiv1.POST("/customer/create", customer.Create)
		apiv1.POST("/customer/delete", customer.Delete)
		apiv1.GET("/customer/listOfPart", customer.ListOfPart)
		apiv1.GET("/customer/one/:id", customer.GetCustomerById)

		apiv1.POST("/right/update", right.Update)
		apiv1.POST("/right/delete", right.Delete)
		apiv1.POST("/right/create", right.Create)

		apiv1.POST("/role/create", role.Create)
		apiv1.POST("/role/delete", role.Delete)

		apiv1.POST("/log/delete", log.Delete)

		apiv1.GET("/join/getAllProcess", join.GetAllProcessList)
		apiv1.GET("/join/getProcessByUser/:id", join.GetProcessListByUser)
		apiv1.GET("/join/getUserWithRole", join.GetUserWithRole)
	}
	return r
}
