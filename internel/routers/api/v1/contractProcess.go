package v1

import (
	"bytes"
	"contract_management/internel/service"
	"contract_management/pkg/app"
	"contract_management/pkg/convert"
	"contract_management/pkg/errcode"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContractProcess struct{}

func NewContractProcess() ContractProcess {
	return ContractProcess{}
}

// @Summary  删除合同流程
// @Produce  json
// @Param    id   path      int             true  "流程id"
// @Success  200  {object}  model.Contract  "成功"
// @Failure  400  {object}  errcode.Error   "请求错误"
// @Failure  500  {object}  errcode.Error   "内部错误"
// @Router   /management/v1/contractProcess/delete/{id} [get]
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
}

// @Summary  修改合同流程
// @Produce  json
// @Param    id          body      int                    true  "流程id"
// @Param    contractId  body      int                    true  "合同id"
// @Param    type        body      int                    true  "操作类型"
// @Param    state       body      int                    true  "操作结果"
// @Param    content     body      string                 true  "操作内容"
// @Param    userId      body      int                    true  "操作人id"
// @Success  200         {object}  model.ContractProcess  "成功"
// @Failure  400         {object}  errcode.Error          "请求错误"
// @Failure  500         {object}  errcode.Error          "内部错误"
// @Router   /management/v1/contractProcess/update [post]
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
	var buffer bytes.Buffer
	switch params.Type {
	case 1:
		buffer.WriteString(fmt.Sprintf("通过了id为%d的合同的会签", params.ContractId))
	case 2:
		if params.ContractId == 1 {
			buffer.WriteString(fmt.Sprintf("通过了id为%d的合同的审批", params.ContractId))
		} else {
			buffer.WriteString(fmt.Sprintf("否决了id为%d的合同的审批", params.ContractId))
		}
	case 3:
		buffer.WriteString(fmt.Sprintf("完成了id为%d的合同的签订", params.ContractId))
	default:

	}
	log := service.CreateLogRequest{UserId: params.UserId, Content: buffer.String()}
	go svc.CreateLog(&log)
	response.ToResponse(gin.H{})
}

// @Summary  合同分配
// @Produce  json
// @Param    contractId  body      int                    true  "合同id"
// @Param    proceses        body      []service.Process                    true  "操作类型和操作人集合"
// @Param	operatorId	body 	 int             true  "分配者id"
// @Success  200         {object}  model.ContractProcess  "成功"
// @Failure  400         {object}  errcode.Error          "请求错误"
// @Failure  500         {object}  errcode.Error          "内部错误"
// @Router   /management/v1/contractProcess/create [post]
func (con ContractProcess) Create(c *gin.Context) {
	params := service.CreateContractProcessRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&params)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	err = svc.CreateContractProcess(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorCreateContractProcessFail)
		return
	}
	log := service.CreateLogRequest{UserId: params.OperatorId, Content: "分配合同:" + strconv.Itoa(params.ContractId)}
	go svc.CreateLog(&log)
	response.ToResponse(gin.H{})
}

func (con ContractProcess) GetContractComment(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := convert.StrTo(c.Param("id")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	Type, err := convert.StrTo(c.Param("type")).Int()
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New()
	comment, err := svc.SelectContractComment(id, Type)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteCustomerFail)
		return
	}
	response.ToResponse(comment)
}
