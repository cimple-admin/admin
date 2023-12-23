package model

type User struct {
	BaseModel
	Name  string `gorm:"default:'', not null"`
	Email string `gorm:"unique"`
	Cover string `gorm:"default:'', not null"`
}
