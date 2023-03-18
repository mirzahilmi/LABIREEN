package services

import (
	"labireen/entities"
	"labireen/repositories"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(menu entities.MenuRegisterParams) error
	GetAllMenu() ([]entities.Menu, error)
	GetMenu(name string) (entities.Menu, error)
	EditMenu(menu entities.MenuRequestParams) error
	DeleteMenu(id uuid.UUID) error
}

type menuServiceImpl struct {
	rp repositories.MenuRepository
}

func NewMenuService(rp repositories.MenuRepository) MenuService {
	return &menuServiceImpl{rp}
}

func (svc *menuServiceImpl) CreateMenu(menu entities.MenuRegisterParams) error {

	newMenu := entities.Menu{
		ID:         uuid.New(),
		MerchantID: menu.MenuRegister.MerchantID,
		Name:       menu.MenuRegister.Name,
		MenuGroups: make([]entities.MenuGroup, len(menu.MenuRegister.MenuGroups)),
	}

	for i, groupRequest := range menu.MenuRegister.MenuGroups {
		newMenu.MenuGroups[i] = entities.MenuGroup{
			ID:          uuid.New(),
			Name:        groupRequest.Name,
			Description: groupRequest.Description,
			MenuID:      newMenu.ID,
			MenuItems:   make([]entities.MenuItem, len(groupRequest.MenuItems)),
		}

		for j, itemRequest := range groupRequest.MenuItems {
			newMenu.MenuGroups[i].MenuItems[j] = entities.MenuItem{
				ID:          uuid.New(),
				Name:        itemRequest.Name,
				Price:       itemRequest.Price,
				Description: itemRequest.Description,
				Stock:       itemRequest.Stock,
				Photo:       itemRequest.Photo,
				MenuGroupID: newMenu.MenuGroups[i].ID,
			}
		}
	}

	err := svc.rp.Create(&newMenu)
	if err != nil {
		return err
	}

	return nil
}

func (svc *menuServiceImpl) GetAllMenu() ([]entities.Menu, error) {
	menus, err := svc.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return *menus, nil
}

func (svc *menuServiceImpl) GetMenu(name string) (entities.Menu, error) {
	menu, err := svc.rp.GetWhere("name", name)
	if err != nil {
		return entities.Menu{}, err
	}

	return *menu, nil
}

func (svc *menuServiceImpl) EditMenu(menu entities.MenuRequestParams) error {
	menuResp, err := svc.rp.GetWhere("name", menu.MenuRequests.Name)
	if err != nil {
		return err
	}

	menuResp = &entities.Menu{
		MerchantID: menu.MenuRequests.MerchantID,
		Name:       menu.MenuRequests.Name,
		MenuGroups: make([]entities.MenuGroup, len(menuResp.MenuGroups)),
	}

	for i, groupRequest := range menuResp.MenuGroups {
		menuResp.MenuGroups[i] = entities.MenuGroup{
			Name:        groupRequest.Name,
			Description: groupRequest.Description,
			MenuID:      menuResp.ID,
			MenuItems:   make([]entities.MenuItem, len(groupRequest.MenuItems)),
		}

		for j, itemRequest := range groupRequest.MenuItems {
			menuResp.MenuGroups[i].MenuItems[j] = entities.MenuItem{
				Name:        itemRequest.Name,
				Price:       itemRequest.Price,
				Description: itemRequest.Description,
				Stock:       itemRequest.Stock,
				Photo:       itemRequest.Photo,
				MenuGroupID: menuResp.MenuGroups[i].ID,
			}
		}
	}

	if err := svc.rp.Update(menuResp); err != nil {
		return err
	}

	return nil
}

func (svc *menuServiceImpl) DeleteMenu(id uuid.UUID) error {
	if err := svc.rp.Delete(id); err != nil {
		return err
	}

	return nil
}
