package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}

func (d Dao) GetTX() *Dao {
	return &Dao{engine: d.engine.Begin()}
}

func (d Dao) Commit() {
	d.engine.Commit()
}

func (d Dao) RollBack() {
	d.engine.Rollback()
}
