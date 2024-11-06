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

type CustomerHandler struct {
	custUseCase usecase.CustomerUseCase
}

func (custHandler CustomerHandler) GetAllCustomer(ctx *gin.Context) {
	cust, err := custHandler.custUseCase.GetAllCustomer()
	if err != nil {
		fmt.Printf("CustomerHandler.GetAllCustomer(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching customer data",
		})
		return
	}
	if cust == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cust,
	})
}

func (custHandler CustomerHandler) GetCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID cannot be empty",
		})
		return
	}

	cust, err := custHandler.custUseCase.GetCustomerById(id)
	if err != nil {
		fmt.Printf("CustomerHandler.GetCustomerById(): %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching customer data",
		})
		return
	}
	if cust == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cust,
	})
}

func (custHandler CustomerHandler) GetCustomerByName(ctx *gin.Context) {
	customer := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}
	if customer.FullName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}

	cust, err := custHandler.custUseCase.GetCustomerByName(customer.FullName)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.GetCustomerByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.GetCustomerByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching customer data",
			})
			return
		}
		return
	}
	if cust == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cust,
	})
}

func (custHandler CustomerHandler) InsertCustomer(ctx *gin.Context) {
	cust := &model.CustomerRequestModel{}
	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = custHandler.custUseCase.InsertCustomer(cust, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.InsertCustomer() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.InsertCustomer() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving customer data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (custHandler CustomerHandler) EditCustomer(ctx *gin.Context) {
	cust := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = custHandler.custUseCase.EditCustomerById(cust, ctx)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.EditCustomer() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.EditCustomer() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving customer data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewCustomerHandler(srv *gin.Engine, custUseCase usecase.CustomerUseCase) *CustomerHandler {
	custHandler := &CustomerHandler{
		custUseCase: custUseCase,
	}

	// route
	srv.POST("/customer", middleware.RequireToken(), middleware.LevelUserAdmin(), custHandler.InsertCustomer)
	srv.PUT("/customer", middleware.RequireToken(), middleware.LevelUserAdmin(), custHandler.EditCustomer)
	srv.GET("/customer/:id", middleware.RequireToken(), custHandler.GetCustomerById)
	srv.GET("/customer", middleware.RequireToken(), middleware.LevelUserAdmin(), custHandler.GetCustomerByName)
	srv.GET("/customers", middleware.RequireToken(), middleware.LevelUserAdmin(), custHandler.GetAllCustomer)

	return custHandler
}
