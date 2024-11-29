package handlers

import (
	"demo/interfaces"
	"demo/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	interfaces.IUser
}

func NewUserHandler(iuser interfaces.IUser) *Userhandler {
	return &Userhandler{IUser: iuser}
}

func (u *Userhandler) CreateUser() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := new(models.User)
		err := ctx.Bind(user)
		if err != nil {
			log.Println(err.Error())
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}
		err = user.Validate()
		if err != nil {
			log.Println(err.Error())
			ctx.String(http.StatusBadRequest, err.Error())
			ctx.Abort()
			return
		}

		// similate a database object is created and stored

		user.Status = "active"
		user.LastModified = time.Now().Unix()
		user.Id = 0
		//user, err = database.CreateUser(db, user)
		user, err = u.IUser.CreateUser(user)
		if err != nil {
			log.Println(err.Error())
			ctx.String(http.StatusBadRequest, "could not create user")
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusCreated, user)

	}
}
func ValidateToken() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("token") == "123456" {
			ctx.Next()
		} else {
			log.Println("invalid token")
			ctx.String(http.StatusBadRequest, "invalid token")
			ctx.Abort()
			return
		}
	}
}
