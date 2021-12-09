package service

import (
	"contract_management/global"
	"contract_management/internel/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() Service {
	svc := Service{}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
