package database

import (
	"github.com/google/uuid"
	"github.com/imdario/mergo"
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	Email     string `json:"email"`
	UUID      string `gorm:"primaryKey" json:"uuid"`
	Validated bool   `gorm:"default:false"`
}

func (email *Email) BeforeCreate(tx *gorm.DB) error {
	email.UUID = (uuid.New()).String()
	email.Validated = false
	return nil
}

func (email *Email) Create() *Email {
	db := GetDatabase()
	db.Create(email)

	return email
}

func (email *Email) ReadSingle() *Email {
	db.First(email)
	return email
}

func (email Email) ReadMultiple() (emails []Email) {
	db.Where(email).Find(&emails)
	return
}

func (email *Email) Update() *Email {
	oldEmail := Email{
		UUID: email.UUID,
	}

	oldEmail.ReadSingle()

	db.Model(&oldEmail).Updates(*email)

	mergo.Merge(email, oldEmail)

	return email
}

func (email *Email) Delete() *Email {
	db.Delete(email)
	return email
}
