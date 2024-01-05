package model

import "github.com/gofiber/fiber/v2"

type User struct {
	BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Cover    string `gorm:"not null"`
	Password string `gorm:"not null;"`
}

func (u User) ToUserInfo() fiber.Map {
	return fiber.Map{
		"ID":    u.ID,
		"Name":  u.Name,
		"Email": u.Email,
		"Cover": u.Cover,
	}
}
