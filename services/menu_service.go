package services

import (
	"labireen/entities"
	"labireen/repositories"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(menu entities.MenuRequestParams, id uuid.UUID) error
	GetMenu(id uuid.UUID) (entities.Menu, error)
	EditMenu(menu entities.MenuRequestParams) error
}

type menuServiceImpl struct {
	rp repositories.MenuRepository
}

func NewMenuService(rp repositories.MenuRepository) MenuService {
	return &menuServiceImpl{rp}
}

func (svc *menuServiceImpl) CreateMenu(menu entities.MenuRequestParams, id uuid.UUID) error {

	newMenu := entities.Menu{
		ID:         uuid.New(),
		MerchantID: id,
		Name:       menu.MenuRequests.Name,
		MenuGroups: make([]entities.MenuGroup, len(menu.MenuRequests.MenuGroups)),
	}

	for i, groupRequest := range menu.MenuRequests.MenuGroups {
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
func (svc *menuServiceImpl) GetMenu(id uuid.UUID) (entities.Menu, error) {
	menu, err := svc.rp.GetWhere("merchant_id", id.String())
	if err != nil {
		return entities.Menu{}, err
	}

	return *menu, nil
}

func (svc *menuServiceImpl) EditMenu(menu entities.MenuRequestParams) error {
	menuSearch, err := svc.rp.GetByID(menu.MenuRequests.MerchantID)
	if err != nil {
		return err
	}

	menuSearch.Name = menu.MenuRequests.Name
	menuSearch.MenuGroups = make([]entities.MenuGroup, len(menuSearch.MenuGroups))

	for i, groupRequest := range menuSearch.MenuGroups {
		menuSearch.MenuGroups[i] = entities.MenuGroup{
			Name:        groupRequest.Name,
			Description: groupRequest.Description,
			MenuItems:   make([]entities.MenuItem, len(groupRequest.MenuItems)),
			MenuID:      menuSearch.ID,
		}

		for j, itemRequest := range groupRequest.MenuItems {
			menuSearch.MenuGroups[i].MenuItems[j] = entities.MenuItem{
				Name:        itemRequest.Name,
				Price:       itemRequest.Price,
				Description: itemRequest.Description,
				Stock:       itemRequest.Stock,
				Photo:       itemRequest.Photo,
				MenuGroupID: menuSearch.MenuGroups[i].ID,
			}
		}
	}

	if err := svc.rp.Update(menuSearch); err != nil {
		return err
	}

	return nil
}
