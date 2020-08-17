package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/manan47billion/gin-gonic/gin/entity"
	"github.com/manan47billion/gin-gonic/gin/service"
)

type AccountController interface {
	FindAll() []entity.Account
	Save(ctx *gin.Context) entity.Account
}

type controller struct {
	service service.AccountService
}

func New(service service.AccountService) AccountController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Account {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Account {
	var account entity.Account
	ctx.BindJSON(&account)
	c.service.Save(account)
	return account
}
