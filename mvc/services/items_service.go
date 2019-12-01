package services

import (
	"net/http"

	"github.com/janhaans/learngo/mvc/domain"
	"github.com/janhaans/learngo/mvc/utils"
)

type itemsService struct{}

var (
	//ItemsService provides access to methods of itemsService
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemID int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
