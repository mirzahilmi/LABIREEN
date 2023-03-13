package repositories

import (
	"labireen/order_service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *entities.Menu) error
	GetByID(id uuid.UUID) (*entities.Menu, error)
	GetWhere(param string, args string) (*entities.Menu, error)
	Update(menu *entities.Menu) error
	Delete(menu *entities.Menu) error
}

type menuRepositoryImpl struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepositoryImpl{db}
}

func (r *menuRepositoryImpl) Create(menu *entities.Menu) error {
	return r.db.Create(&menu).Error
}
func (r *menuRepositoryImpl) GetByID(id uuid.UUID) (*entities.Menu, error) {
	var menu entities.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		return nil, err
	}

	return &menu, nil
}
func (r *menuRepositoryImpl) GetWhere(param string, args string) (*entities.Menu, error) {
	var menu entities.Menu
	if err := r.db.Where(param+" = ?", args).Preload("MenuGroups.MenuItems").First(&menu).Error; err != nil {
		return nil, err
	}

	return &menu, nil

}

func (r *menuRepositoryImpl) Update(menu *entities.Menu) error {
	return r.db.Save(&menu).Error
}
func (r *menuRepositoryImpl) Delete(menu *entities.Menu) error {
	return r.db.Delete(&menu).Error
}
