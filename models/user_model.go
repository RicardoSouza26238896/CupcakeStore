package models

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"time"

	"github.com/bitebait/cupcakestore/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserFilter struct {
	User       *User
	Pagination *Pagination
}

func NewUserFilter(query string, page, limit int) *UserFilter {
	user := &User{
		Email: query,
	}
	pagination := NewPagination(page, limit)

	return &UserFilter{
		User:       user,
		Pagination: pagination,
	}
}

type User struct {
	gorm.Model
	Email      string    `gorm:"type:varchar(100);unique" validate:"required,email"`
	Password   string    `gorm:"type:varchar(100);" validate:"required,min=8"`
	IsActive   bool      `gorm:"default:true"`
	IsStaff    bool      `gorm:"default:false"`
	FirstLogin time.Time `gorm:"type:timestamp"`
	LastLogin  time.Time `gorm:"type:timestamp"`
}

func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.HashPassword(); err != nil {
		return err
	}

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	var existingProfile Profile
	result := tx.Where("user_id = ?", u.ID).First(&existingProfile)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			profile := &Profile{
				UserID: u.ID,
			}
			return tx.Create(profile).Error
		}
		return result.Error
	}
	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	if err = tx.Where("user_id = ?", u.ID).Delete(&Profile{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) HashPassword() error {
	hash, err := utils.PasswordHasher(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash

	return nil
}

func (u *User) CheckPassword(inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword))
}

func (u *User) UpdatePassword(oldPassword, newPassword string) error {
	if err := u.CheckPassword(oldPassword); err != nil {
		return err
	}

	if newPassword == "" {
		return errors.New("nova senha não pode estar vazia")
	}

	hash, err := utils.PasswordHasher(newPassword)
	if err != nil {
		return err
	}

	u.Password = hash

	return nil
}
