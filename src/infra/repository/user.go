package repository

import "gorm.io/gorm"

// User user info
type User struct {
	gorm.Model
	Name string `gorm:""`
	Age  uint   `gorm:""`
}

func CreateUser(db *gorm.DB, user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, user User) error {
	if err := db.Omit("created_at", "deleted_at").Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func FetchUsers(db *gorm.DB) ([]User, error) {
	users := []User{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUserByID(db *gorm.DB, id uint) error {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func TruncateUsersTable(db *gorm.DB) error {
	if err := db.Exec("TRUNCATE TABLE users").Error; err != nil {
		return err
	}
	return nil
}
