package repository

import (
	"database/sql"
	"fmt"
	"mncbank/model"
	"mncbank/utils"
	"time"
)

type MemberRepo interface {
	InsertMember(*model.MemberModel) error
	GetAllMember() ([]model.MemberModel, error)
	GetMemberById(string) (*model.MemberModel, error)
	EditMember(string, *model.MemberModel) error
	DeleteMember(id string) error
}

type memberRepoImpl struct {
	db *sql.DB
}

func (c *memberRepoImpl) InsertMember(newMember *model.MemberModel) error {
	fmt.Println(utils.UuidGenerate())
	fmt.Println()
	insertStatement := utils.INSERT_MEMBER
	_, err := c.db.Exec(insertStatement, newMember.ID, newMember.CustomerID, newMember.Type, time.Now().AddDate(0, 1, 0), newMember.CreatedAt, newMember.CreatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (c *memberRepoImpl) GetAllMember() ([]model.MemberModel, error) {
	rows, err := c.db.Query(utils.GET_ALL_MEMBER)
	var uAt sql.NullTime
	var uBy sql.NullString
	if err != nil {
		return nil, err
	}
	var members []model.MemberModel
	member := model.MemberModel{}
	for rows.Next() {
		err = rows.Scan(&member.ID, &member.CustomerID, &member.Type, &member.Expire, &member.CreatedAt, &member.CreatedBy, &uAt, &uBy)
		if err != nil {
			return nil, err
		}
		if uAt.Valid {
			member.UpdatedAt = uAt.Time
		}
		if uBy.Valid {
			member.UpdatedBy = uBy.String
		}
		members = append(members, member)
	}
	return members, nil
}

func (c *memberRepoImpl) GetMemberById(id string) (*model.MemberModel, error) {
	var uAt sql.NullTime
	var UBy sql.NullString
	qry := utils.GET_MEMBER_ID
	member := &model.MemberModel{}
	err := c.db.QueryRow(qry, id).Scan(&member.ID, &member.CustomerID, &member.Type, &member.Expire, &member.CreatedAt, &member.CreatedBy, &uAt, &UBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetMemberById() : %w", err)
	}

	if uAt.Valid {
		member.UpdatedAt = uAt.Time
	}
	if UBy.Valid {
		member.UpdatedBy = UBy.String
	}

	return member, nil
}

func (c *memberRepoImpl) EditMember(id string, newMember *model.MemberModel) error {

	qry := utils.EDIT_MEMBER
	result, err := c.db.Exec(qry, newMember.UpdatedAt, newMember.UpdatedBy, newMember.Type, newMember.Expire, id)
	if err != nil {
		return fmt.Errorf("EditMember() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("editMember(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID %s not found", id)
	}
	return nil
}

func (c *memberRepoImpl) DeleteMember(id string) error {

	qry := utils.DELETE_MEMBER
	result, err := c.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("deleteMember() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteMember(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID %s not found", id)
	}
	return nil
}

// Func seperti consturcor
func NewMemberDbRepository(db *sql.DB) MemberRepo {
	return &memberRepoImpl{
		db: db,
	}
}
