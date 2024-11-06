package repository

import (
	"database/sql"
	"fmt"
	"mncbank/model"
	"mncbank/utils"
)

type CustomerRepo interface {
	InsertCustomer(cust *model.CustomerRequestModel) error
	GetCustomerById(id string) (*model.CustomerModel, error)
	GetCustomerByName(name string) (*model.CustomerModel, error)
	GetAllCustomer() ([]*model.CustomerModel, error)
	EditCustomerById(cust model.CustomerModel) error
	GetCustomerByUserId(id string) (*model.CustomerModel, error)
	GetCustomerWithMemberByIdCustomer(idCustomer string) (string, error)
}

type customerRepoImpl struct {
	db *sql.DB
}

func (custRepo *customerRepoImpl) InsertCustomer(cust *model.CustomerRequestModel) error {
	tx, err := custRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 1 : %w", err)
	}

	qryUser := "INSERT INTO ms_user(id,username, password, role, is_active ) VALUES($1, $2, $3, $4, $5) RETURNING id"
	err = tx.QueryRow(qryUser, utils.UuidGenerate(), cust.Username, cust.Password, cust.Role, cust.Active).Scan(&cust.User_id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 2  : %w", err)
	}

	qry := "INSERT INTO ms_customer (id, id_user, full_name, NIK, noPhone, email, address, created_at, created_by, balance) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);"
	_, err = tx.Exec(qry, utils.UuidGenerate(), cust.User_id, cust.FullName, cust.NIK, cust.NoPhone, cust.Email, cust.Address, cust.CreatedAt, cust.CreatedBy, cust.Balance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer() 3 : %w", err)
	}

	tx.Commit()
	return nil
}

func (custRepo *customerRepoImpl) GetCustomerById(id string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_ID
	cust := &model.CustomerModel{}
	err := custRepo.db.QueryRow(qry, id).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NIK, &cust.NoPhone, &cust.Email, &cust.Address, &cust.CreatedAt, &cust.UpdatedAt, &cust.CreatedBy, &cust.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerById() : %w", err)
	}
	return cust, nil
}

func (custRepo *customerRepoImpl) GetCustomerWithMemberByIdCustomer(idCustomer string) (string, error) {
	qry := utils.GET_CUST_ID_MEMBER
	var cust string
	err := custRepo.db.QueryRow(qry, idCustomer).Scan(&cust)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", fmt.Errorf("error on customerRepoImpl.GetCustomerById() : %w", err)
	}
	return cust, nil
}

func (custRepo *customerRepoImpl) GetCustomerByUserId(id string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_USRID
	cust := &model.CustomerModel{}
	err := custRepo.db.QueryRow(qry, id).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NIK, &cust.NoPhone, &cust.Email, &cust.Address, &cust.CreatedAt, &cust.UpdatedAt, &cust.CreatedBy, &cust.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerById() : %w", err)
	}
	return cust, nil
}

func (custRepo *customerRepoImpl) GetAllCustomer() ([]*model.CustomerModel, error) {
	qry := utils.GET_ALL_CUSTOMER
	rows, err := custRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on customerRepoImpl.GetAllCustomer() : %w", err)
	}
	defer rows.Close()
	var arrCustomer []*model.CustomerModel
	for rows.Next() {
		cust := &model.CustomerModel{}
		rows.Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NIK, &cust.NoPhone, &cust.Email, &cust.Address, &cust.CreatedAt, &cust.UpdatedAt, &cust.CreatedBy, &cust.UpdatedBy)
		arrCustomer = append(arrCustomer, cust)
	}
	return arrCustomer, nil
}

func (custRepo *customerRepoImpl) GetCustomerByName(name string) (*model.CustomerModel, error) {
	qry := utils.GET_CUST_NAME
	cust := &model.CustomerModel{}
	err := custRepo.db.QueryRow(qry, name).Scan(&cust.ID, &cust.User_id, &cust.FullName, &cust.NIK, &cust.NoPhone, &cust.Email, &cust.Address, &cust.CreatedAt, &cust.UpdatedAt, &cust.CreatedBy, &cust.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetCustomerByName() : %w", err)
	}
	return cust, nil
}

func (custRepo *customerRepoImpl) EditCustomerById(cust model.CustomerModel) error {
	qry := utils.EDIT_CUST_ID
	_, err := custRepo.db.Exec(qry, cust.FullName, cust.NIK, cust.NoPhone, cust.Email, cust.Address, cust.UpdatedAt, cust.UpdatedBy, cust.ID)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.EditCustomerById() 3 : %w", err)
	}

	return nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
