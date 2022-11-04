package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/prithuadhikary/user-service/model"
	"github.com/prithuadhikary/user-service/service"
	"github.com/prithuadhikary/user-service/util"
	"net/http"
)

type UserController interface {
	Signup(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func (controller userController) Signup(ctx *gin.Context) {
	request := &model.SignupRequest{}
	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		util.RenderBindingErrors(ctx, err.(validator.ValidationErrors))
		return
	}
	err := controller.service.Signup(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
	}
}

func NewUserController(engine *gin.Engine, userService service.UserService) {
	controller := &userController{
		service: userService,
	}
	api := engine.Group("api")
	{
		api.POST("users", controller.Signup)
	}
}
