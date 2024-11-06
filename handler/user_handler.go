package handler

import (
	"errors"
	"fmt"
	"mncbank/middleware"
	"mncbank/model"
	"mncbank/usecase"
	"mncbank/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type UserHandler struct {
	usrUseCase usecase.UserUseCase
}

func (usrHandler UserHandler) GetAllUser(ctx *gin.Context) {
	usr, err := usrHandler.usrUseCase.GetAllUser()
	if err != nil {
		fmt.Printf("UserHandler.GetAllUser(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching user data",
		})
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID cannot be empty",
		})
		return
	}

	usr, err := usrHandler.usrUseCase.GetUserById(id)
	if err != nil {
		fmt.Printf("UserHandler.GetUserById(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching user data",
		})
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler UserHandler) GetUserByName(ctx *gin.Context) {
	user := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if user.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}

	usr, err := usrHandler.usrUseCase.GetUserByName(user)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.GetUserByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.GetUserByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching user data",
			})
			return
		}
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler UserHandler) InsertUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.InsertUser(usr)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (usrHandler UserHandler) EditUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.EditUserById(usr)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.EditUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.EditUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewUserHandler(srv *gin.Engine, usrUseCase usecase.UserUseCase) *UserHandler {
	usrHandler := &UserHandler{
		usrUseCase: usrUseCase,
	}

	// route
	srv.POST("/user", middleware.RequireToken(), middleware.LevelUserAdmin(), usrHandler.InsertUser)
	srv.PUT("/user", middleware.RequireToken(), middleware.LevelUserAdmin(), usrHandler.EditUser)
	srv.GET("/user/:id", middleware.RequireToken(), middleware.LevelUserAdmin(), usrHandler.GetUserById)
	srv.GET("/user", middleware.RequireToken(), middleware.LevelUserAdmin(), usrHandler.GetUserByName)
	srv.GET("/users", middleware.RequireToken(), middleware.LevelUserAdmin(), usrHandler.GetAllUser)

	return usrHandler
}
