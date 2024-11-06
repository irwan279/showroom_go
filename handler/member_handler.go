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
)

type MemberHandler struct {
	router  *gin.Engine
	usecase usecase.MemberUsecase
}

func (mh *MemberHandler) GetAllMember(ctx *gin.Context) {
	members, err := mh.usecase.GetAllMember()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, members)
}
func (mh *MemberHandler) GetMemberById(ctx *gin.Context) {
	id := ctx.Param("id")
	members, err := mh.usecase.GetMemberById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, members)
}

func (mh *MemberHandler) InsertMember(ctx *gin.Context) {
	var member *model.MemberModel

	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	member.ID = utils.UuidGenerate()
	if err := mh.usecase.InsertMember(member, ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, member)
}

func (mh *MemberHandler) EditMember(ctx *gin.Context) {
	id := ctx.Param("id")
	mbr := &model.MemberModel{}
	err := ctx.ShouldBindJSON(&mbr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = mh.usecase.EditMember(id, mbr, ctx)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("memberHandler.UpdateMember() 2 : %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("memberHandler.UpdateMember() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while updating member data.",
			})
			return
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func (mh *MemberHandler) DeleteMember(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id cannot be empty",
		})
		return
	}

	err := mh.usecase.DeleteMember(id)
	if err != nil {
		fmt.Printf("MemberHandler.GetmemberById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred while fetching member data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewMemberHandler(r *gin.Engine, usecase usecase.MemberUsecase) *MemberHandler {
	controller := MemberHandler{
		router:  r,
		usecase: usecase,
	}
	r.GET("/member", middleware.RequireToken(), middleware.LevelUserAdmin(), controller.GetAllMember)
	r.GET("/member/:id", middleware.RequireToken(), controller.GetMemberById)
	r.POST("/member", middleware.RequireToken(), middleware.LevelUserAdmin(), controller.InsertMember)
	r.PUT("/member/:id", middleware.RequireToken(), middleware.LevelUserAdmin(), controller.EditMember)
	r.DELETE("/member/:id", middleware.RequireToken(), middleware.LevelUserAdmin(), controller.DeleteMember)
	return &controller
}
