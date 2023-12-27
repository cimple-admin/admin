package model

type BaseModel struct {
	ID       uint
	CreateAt uint `gorm:"autoCreateTime;not null"`
	UpdateAt uint `gorm:"autoUpdateTime;not null"`
}
