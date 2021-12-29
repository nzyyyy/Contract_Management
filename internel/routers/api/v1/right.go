package v1

type Right struct{}

func NewRight() Right {
	return Right{}
}

//func (r Right) Create(c *gin.Context) {
//	params := service.CreateOrUpdateRightRequest{}
//	response := app.NewResponse(c)
//	err := c.ShouldBind(&params)
//	if err != nil {
//		response.ToErrorResponse(errcode.InvalidParams)
//		return
//	}
//	svc := service.New()
//	err = svc.CreateRight(&params)
//	if err != nil {
//		response.ToErrorResponse(errcode.ErrorCreateRightFail)
//		return
//	}
//	log := service.CreateLogRequest{UserId: params.OperatorId, Content: fmt.Sprintf("对用户id:%d授权角色id:%d的角色", params.UserId, params.RoleId)}
//	go svc.CreateLog(&log)
//	response.ToResponse(gin.H{})
//}
//
//
//func (r Right) Delete(c *gin.Context) {
//	params := service.DeleteRightRequest{}
//	response := app.NewResponse(c)
//	err := c.ShouldBind(&params)
//	if err != nil {
//		response.ToErrorResponse(errcode.InvalidParams)
//		return
//	}
//	svc := service.New()
//	err = svc.DeleteRight(params.UserId)
//	if err != nil {
//		response.ToErrorResponse(errcode.ErrorDeleteRightFail)
//		return
//	}
//	log := service.CreateLogRequest{UserId: params.OperatorId, Content: fmt.Sprintf("删除用户id为%d的权限", params.UserId)}
//	go svc.CreateLog(&log)
//	response.ToResponse(gin.H{})
//}
//
//
//func (r Right) Update(c *gin.Context) {
//	params := service.CreateOrUpdateRightRequest{}
//	response := app.NewResponse(c)
//	err := c.ShouldBind(&params)
//	if err != nil {
//		response.ToErrorResponse(errcode.InvalidParams)
//		return
//	}
//	svc := service.New()
//	err = svc.UpdateRight(&params)
//	if err != nil {
//		response.ToErrorResponse(errcode.ErrorUpdateRightFail)
//		return
//	}
//	log := service.CreateLogRequest{UserId: params.OperatorId, Content: fmt.Sprintf("对用户id:%d修改为角色id:%d的角色", params.UserId, params.RoleId)}
//	go svc.CreateLog(&log)
//	response.ToResponse(gin.H{})
//}
