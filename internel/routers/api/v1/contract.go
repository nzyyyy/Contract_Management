package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Contract struct{}

func NewContract() Contract {
	return Contract{}
}

// @Summary 新建合同
// @Produce json
// @Param name body string true "合同名称"
// @Param customerId body int true "客户id"
// @Param beginTime body time.Time true "开始时间"
// @Param endTime body time.Time true "结束时间"
// @Param content body string true "合同内容"
// @Param userId body int true "起草人id"
// @Param fileName string int true "文件名"
// @Param path body string true "文件路径"
// @Param type body string true "文件后缀"
// @Success 200 {object} model.Contract "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/contract/create [post]
func (con Contract) Create(c *gin.Context) {
	params := service.CreateContractRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.CreateContract(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 删除合同
// @Produce json
// @Param id path int true "合同id"
// @Success 200 {object} model.Contract "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/contract/delete/{id} [get]
func (con Contract) Delete(c *gin.Context) {
	id, err := convert.StrTo(c.Param("id")).Int()
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	params := service.DeleteContractRequest{ID: id}

	svc := service.New()
	err = svc.DeleteContract(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 修改合同
// @Produce json
// @Param id body int true "合同id"
// @Param content body string true "合同内容"
// @Success 200 {object} model.Contract "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/contract/update [post]
func (con Contract) Update(c *gin.Context) {
	params := service.UpdateContractRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.UpdateContract(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateContractFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
