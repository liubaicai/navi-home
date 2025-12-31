package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint      `gorm:"primarykey" json:"id"`
	Email             string    `gorm:"uniqueIndex;not null" json:"email"`
	EncryptedPassword string    `gorm:"not null" json:"-"`
	Username          string    `gorm:"not null" json:"username"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Catalogs          []Catalog `gorm:"foreignKey:UserID" json:"catalogs,omitempty"`
	Links             []Link    `gorm:"foreignKey:UserID" json:"links,omitempty"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	result := DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByID(id uint) (*User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func CreateUser(email, username, password string) (*User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	result := DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) Update(email, username, password, currentPassword string) error {
	// Verify current password
	if !u.CheckPassword(currentPassword) {
		return bcrypt.ErrMismatchedHashAndPassword
	}

	u.Email = email
	u.Username = username

	if password != "" {
		if err := u.SetPassword(password); err != nil {
			return err
		}
	}

	return DB.Save(u).Error
}
