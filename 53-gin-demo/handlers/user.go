package handlers

import (
	"demo/models"
	"log"
	"net/http"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
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

	user.Id = rand.Intn(100)
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	ctx.JSON(http.StatusCreated, user)

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
