package repository

import (
	"database/sql"
	"fmt"
	"mncbank/model"
	"mncbank/utils"
)

type LoginRepo interface {
	GetUserByName(name string) (*model.UserModel, error)
}

type loginRepoImpl struct {
	db *sql.DB
}

func (usrRepo *loginRepoImpl) GetUserByName(name string) (*model.UserModel, error) {
	qry := utils.GET_USER_BY_NAME
	usr := &model.UserModel{}
	err := usrRepo.db.QueryRow(qry, name).Scan(&usr.Id, &usr.UserName, &usr.Password, &usr.Role, &usr.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on loginRepoImpl.GetUserByName() : %w", err)
	}
	return usr, nil
}

func NewLoginRepo(db *sql.DB) LoginRepo {
	return &loginRepoImpl{
		db: db,
	}
}
