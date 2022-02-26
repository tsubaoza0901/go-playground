package repository

import "gorm.io/gorm"

// User user info
type User struct {
	gorm.Model
	Name string `gorm:""`
	Age  uint   `gorm:""`
}

func (u *User) CreateUser(db *gorm.DB, user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
