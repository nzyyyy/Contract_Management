package model

type Function struct {
	ID          int `gorm:"autoIncrement"`
	Name        string
	Url         string
	Description string
}

func (Function) TableName() string {
	return "function"
}
