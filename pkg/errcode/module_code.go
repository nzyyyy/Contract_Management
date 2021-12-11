package errcode

var (
	ErrorCreateContractFail = NewError(20010001, "创建合同失败")
	ErrorUpdateContractFail = NewError(20010002, "更新合同失败")
	ErrorDeleteContractFail = NewError(20010003, "删除合同失败")

	ErrorUpdateContractProcessFail = NewError(20020001, "更新合同流程失败")
	ErrorDeleteContractProcessFail = NewError(20020002, "删除合同流程失败")
	ErrorCreateContractProcessFail = NewError(20020003, "分配合同失败")

	ErrorCreateCustomerFail = NewError(20030001, "新增客户失败")
	ErrorDeleteCustomerFail = NewError(20030002, "删除客户失败")

	ErrorUpdateRightFail = NewError(20040001, "更新权限失败")
	ErrorDeleteRightFail = NewError(20040002, "删除权限失败")
	ErrorCreateRightFail = NewError(20040003, "分配权限失败")

	ErrorCreateRoleFail = NewError(20050001, "新增角色失败")
	ErrorDeleteRoleFail = NewError(20050002, "删除角色失败")
)
