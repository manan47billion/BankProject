package service

import (
	"github.com/manan47billion/gin-gonic/gin/entity"
)

type AccountService interface {
	Save(entity.Account) entity.Account
	FindAll() []entity.Account
}

type accountService struct {
	accounts []entity.Account
}

func New() AccountService {
	return &accountService{
		accounts: []entity.Account{},
	}
}
func (service *accountService) Save(account entity.Account) entity.Account {
	service.accounts = append(service.accounts, account)
	return account
}

func (service *accountService) FindAll() []entity.Account {
	return service.accounts
}
