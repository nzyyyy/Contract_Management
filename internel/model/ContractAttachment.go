package model

import (
	"time"

	"gorm.io/gorm"
)

type ContractAttachment struct {
	ContractId int    `gorm:"column:con_id;primaryKey"`
	FileName   string `gorm:"column:fileName"`
	Path       string
	Type       string
	UploadTime time.Time `gorm:"column:uploadTime;autoCreateTime"`
}

func (ContractAttachment) TableName() string {
	return "contract_attachment"
}

func (c ContractAttachment) Create(db *gorm.DB) error {
	return db.Create(&c).Error
}
func (c ContractAttachment) Delete(db *gorm.DB) error {
	return db.Delete(&c).Error
}
