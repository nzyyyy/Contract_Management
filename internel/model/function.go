package model

type Function struct {
	ID          int    `gorm:"autoIncrement" json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func (Function) TableName() string {
	return "function"
}
