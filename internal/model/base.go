package model

type BaseModel struct {
	ID       uint
	CreateAt uint `gorm:"autoCreateTime,default:0,not null"`
	UpdateAt uint `gorm:"autoUpdateTime,default:0,not null"`
}
