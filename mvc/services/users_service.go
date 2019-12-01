package services

import (
	"github.com/janhaans/learngo/mvc/domain"
	"github.com/janhaans/learngo/mvc/utils"
)

type usersService struct{}

var (
	// UsersService provides access to methods of usersService
	UsersService usersService
)

func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UsersDao.GetUser(userID)
}
