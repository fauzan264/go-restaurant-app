package resto

import "github.com/fauzan264/go-restaurant-app/internal/model"

type Usecase interface{
	GetMenu(menuType string) ([]model.MenuItem, error)
}