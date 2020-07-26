package database

import (
	"errors"

	"github.com/eikendev/pushbits/model"

	"gorm.io/gorm"
)

// CreateApplication creates an application.
func (d *Database) CreateApplication(application *model.Application) error {
	return d.gormdb.Create(application).Error
}

// DeleteApplication deletes an application.
func (d *Database) DeleteApplication(application *model.Application) error {
	return d.gormdb.Delete(application).Error
}

// UpdateApplication updates an application.
func (d *Database) UpdateApplication(app *model.Application) error {
	return d.gormdb.Save(app).Error
}

// GetApplicationByID returns the application for the given ID or nil.
func (d *Database) GetApplicationByID(ID uint) (*model.Application, error) {
	app := new(model.Application)
	err := d.gormdb.First(&app, ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return app, err
}

// GetApplicationByToken returns the application for the given token or nil.
func (d *Database) GetApplicationByToken(token string) (*model.Application, error) {
	app := new(model.Application)
	err := d.gormdb.Where("token = ?", token).First(app).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return app, err
}
