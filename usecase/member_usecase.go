package usecase

import (
	"fmt"
	"mncbank/model"
	"mncbank/repository"
	"mncbank/utils"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MemberUsecase interface {
	InsertMember(*model.MemberModel, *gin.Context) error
	GetAllMember() ([]model.MemberModel, error)
	GetMemberById(string) (*model.MemberModel, error)
	EditMember(string, *model.MemberModel, *gin.Context) error
	DeleteMember(id string) error
}
type memberUsecaseImpl struct {
	repoMember repository.MemberRepo
	repoCus    repository.CustomerRepo
}

func (c *memberUsecaseImpl) InsertMember(member *model.MemberModel, ctx *gin.Context) error {
	typeMbr := strings.ToLower(member.Type)
	switch typeMbr {
	case "bronze":
	case "silver":
	case "gold":
	default:
		return fmt.Errorf("there are only 3 types of members (Bronze, Silver, Gold)")
	}
	cust, err := c.repoCus.GetCustomerById(member.CustomerID)
	if err != nil {
		return fmt.Errorf("error Test 1 %v", err)
	}
	if cust == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("ID not found %v", member.CustomerID),
		}
	}
	session := sessions.Default(ctx)
	username := session.Get("Username")
	createdby, _ := username.(string)
	member.CreatedBy = createdby
	member.CreatedAt = time.Now().UTC()
	return c.repoMember.InsertMember(member)
}

func (c *memberUsecaseImpl) GetMemberById(id string) (*model.MemberModel, error) {
	return c.repoMember.GetMemberById(id)
}

func (c *memberUsecaseImpl) GetAllMember() ([]model.MemberModel, error) {
	return c.repoMember.GetAllMember()
}

func (c *memberUsecaseImpl) EditMember(id string, member *model.MemberModel, ctx *gin.Context) error {
	mbr, err := c.repoMember.GetMemberById(id)
	if err != nil {
		return err
	}
	if mbr == nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("ID not found %v", id),
		}
	}
	member.Expire = mbr.Expire.AddDate(0, 1, 0)
	session := sessions.Default(ctx)
	username := session.Get("Username")
	updatedby, _ := username.(string)
	member.UpdatedBy = updatedby
	member.UpdatedAt = time.Now().UTC()
	return c.repoMember.EditMember(id, member)
}

func (c *memberUsecaseImpl) DeleteMember(id string) error {
	return c.repoMember.DeleteMember(id)
}

func NewMemberUseCase(repoMember repository.MemberRepo, repoCus repository.CustomerRepo) MemberUsecase {
	return &memberUsecaseImpl{
		repoMember: repoMember,
		repoCus:    repoCus,
	}
}
