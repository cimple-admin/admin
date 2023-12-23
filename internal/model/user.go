package model

type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Cover    string `gorm:"not null"`
	Password string `gorm:"not null;"`
}
