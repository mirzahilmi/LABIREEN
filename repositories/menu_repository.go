package repositories

import (
	"labireen/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *entities.Menu) error
	GetAll() (*[]entities.Menu, error)
	GetByID(id uuid.UUID) (*entities.Menu, error)
	GetWhere(param string, args string) (*entities.Menu, error)
	Update(menu *entities.Menu) error
	Delete(id uuid.UUID) error
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

func (r *menuRepositoryImpl) GetAll() (*[]entities.Menu, error) {
	var menus []entities.Menu
	if err := r.db.Preload("MenuGroups.MenuItems").Find(&menus).Error; err != nil {
		return nil, err
	}
	return &menus, nil
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

func (r *menuRepositoryImpl) Delete(id uuid.UUID) error {
	var menu entities.Menu
	if err := r.db.Where("merchant_id = ?", id).Preload("MenuGroups.MenuItems.OrderItems").First(&menu).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&menu).Error; err != nil {
		return err
	}

	return nil
}
