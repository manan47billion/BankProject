package main

import (
	"log"
	_ "net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/manan47billion/gin-gonic/gin/Account"
	"github.com/manan47billion/gin-gonic/gin/Entity"
	//	"github.com/manan47billion/gin-gonic/gin/controller"
	//	"github.com/manan47billion/gin-gonic/gin/service"
)

/*var (
	accountService    service.AccountService       = service.New()
	accountController controller.AccountController = controller.New(accountService)
)*/
var account = &Entity.Account{}

func main() {
	server := gin.Default()

	/*	server.GET("/accounts", func(ctx *gin.Context) {
			ctx.JSON(200, accountController.FindAll())
		})

		server.POST("/createaccount", func(ctx *gin.Context) {
			ctx.JSON(200, accountController.Save(ctx))
		}) */
	server.POST("/createaccount", func(context *gin.Context) {
		//		ctx.JSON(200, accountController.FindAll())
		req := context.Request
		err := json.NewDecoder(req.Body).Decode(account)
		if err != nil {
			log.Fatal()
		}
		Account.CreateAccount()
	})

	server.Run(":8080")
}
