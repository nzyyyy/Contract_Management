package errcode

var (
	ErrorCreateContractFail = NewError(20010001, "创建合同失败")
	ErrorUpdateContractFail = NewError(20010002, "更新合同失败")
	ErrorDeleteContractFail = NewError(20010003, "删除合同失败")

	ErrorUpdateContractProcessFail = NewError(20020001, "更新合同流程失败")
	ErrorDeleteContractProcessFail = NewError(20020002, "删除合同流程失败")
)
