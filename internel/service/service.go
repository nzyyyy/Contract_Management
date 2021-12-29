package service

import (
	"contract_management/global"
	"contract_management/internel/dao"
)

type Email struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
type Service struct {
	dao *dao.Dao
}

func New() Service {
	svc := Service{}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
