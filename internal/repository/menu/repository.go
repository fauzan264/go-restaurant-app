package menu

import "github.com/fauzan264/go-restaurant-app/internal/model"

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
}