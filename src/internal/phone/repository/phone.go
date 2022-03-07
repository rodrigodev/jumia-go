package repository

import (
	"context"
	"github.com/rodrigodev/jumia-go/src/internal/infrastructure"
	"github.com/rodrigodev/jumia-go/src/internal/phone/model"
	"gorm.io/gorm"
)

type CustomerManager interface {
	GetAllCustomer(ctx context.Context) ([]model.Customer, error)
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewPhoneRepository(db *gorm.DB) (*CustomerRepository, error) {
	r := &CustomerRepository{
		db: db,
	}
	return r, nil
}

func (u *CustomerRepository) GetAllCustomer(ctx context.Context) ([]model.Customer, error) {
	var customerPhones []model.Customer

	tx := u.db.Find(&customerPhones)
	if tx.Error != nil {
		infrastructure.Logger(ctx).Error("error fetching customerPhones")
		return nil, tx.Error
	}
	return customerPhones, nil
}
