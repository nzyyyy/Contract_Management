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
	//right := v1.NewRight()
	role := v1.NewRole()
	log := v1.NewLog()
	line := v1.NewTimeLine()
	join := v1.NewJoin()
	user := v1.NewUser()
	function := v1.NewFunction()
	apiv1 := r.Group("management/v1")
	{

		apiv1.POST("/contract/create", contract.Create)                                     //添加合同
		apiv1.POST("/contract/delete", contract.Delete)                                     //删除合同
		apiv1.POST("/contract/update", contract.Update)                                     //定稿合同
		apiv1.GET("/contract/allList", contract.GetAllList)                                 //获取所有合同列表
		apiv1.GET("/contract/listByUser/:id", contract.GetListByUser)                       //获取指定id的合同列表
		apiv1.GET("/contract/contractById/:id", contract.GetContractById)                   //获取指定id的合同信息
		apiv1.GET("/contract/contractWithoutOperator", contract.GEtContractWithoutOperator) //获取指定id的合同信息

		apiv1.POST("/contractProcess/update", contractProcess.Update)                       //合同会签、审批、签订
		apiv1.GET("/contractProcess/delete/:id", contractProcess.Delete)                    //删除合同流程信息
		apiv1.POST("/contractProcess/create", contractProcess.Create)                       //分配合同
		apiv1.GET("/contractProcess/comment/:type/:id", contractProcess.GetContractComment) //获取指定合同的会签、审批、签订信息
		apiv1.GET("/contractProcess/operators", contractProcess.GetContractOperators)       //获取合同操作员

		apiv1.POST("/customer/create", customer.Create)          //创建客户
		apiv1.POST("/customer/delete", customer.Delete)          //删除客户
		apiv1.POST("/customer/update", customer.UpdateCustomer)  //修改客户
		apiv1.GET("/customer/listOfPart", customer.ListOfPart)   //获取所有客户列表
		apiv1.GET("/customer/one/:id", customer.GetCustomerById) //获取指定客户信息

		apiv1.POST("/role/create", role.Create)                       //创建角色
		apiv1.POST("/role/delete", role.Delete)                       //删除角色
		apiv1.GET("/role/getAll", role.GetAllRole)                    //获取所有角色
		apiv1.GET("/role/getOne/:id", role.GetRoleById)               //获取指定角色信息
		apiv1.GET("/role/getWithFunc/:id", role.GetRoleWithFunctions) //获取拥有某种权限的角色
		apiv1.GET("/role/getFunc/:id", role.GetFuncByRoleId)          //获取角色的权限

		apiv1.GET("/log/delete/:id", log.Delete) //删除日志
		apiv1.GET("/log/get", log.GetAll)        //获取所有日志

		apiv1.GET("/line/:id", line.GetTimeLineById) //获取合同时间线

		apiv1.GET("/join/getAllProcess", join.GetAllProcessList)           //获取所有合同的当前状态
		apiv1.GET("/join/getProcessByUser/:id", join.GetProcessListByUser) //获取指定用户所起草的合同当前状态
		apiv1.GET("/join/getUserWithRole", join.GetUserWithRole)           //获取所有用户以及角色名称

		apiv1.GET("/user/updateRole/:userId/:roleId", user.UpdateRole) //修改用户角色

		apiv1.GET("/function/list", function.GetList)
	}
	return r
}
