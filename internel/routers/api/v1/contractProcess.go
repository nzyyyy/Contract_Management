package v1

import (
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type ContractProcess struct{}

func NewContractProcess() ContractProcess {
	return ContractProcess{}
}

// @Summary 删除合同流程
// @Produce json
// @Param id path int true "流程id"
// @Success 200 {object} model.Contract "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/contractProcess/delete/{id} [get]
func (con ContractProcess) Delete(c *gin.Context) {
	id, err := convert.StrTo(c.Param("id")).Int()
	response := app.NewResponse(c)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	params := service.DeleteContractProcessRequest{ID: id}
	svc := service.New()
	err = svc.DeleteContractProcess(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteContractProcessFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 修改合同流程
// @Produce json
// @Param id body int true "流程d"
// @Param contractId body int true "合同id"
// @Param type body int true "操作类型"
// @Param State body int true "操作结果"
// @Param content body string true "操作内容"
// @Success 200 {object} model.ContractProcess "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /management/v1/contractProcess/update [post]
func (con ContractProcess) Update(c *gin.Context) {
	params := service.UpdateContractProcessRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.UpdateContractProcess(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorUpdateContractProcessFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
