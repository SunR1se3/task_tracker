package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"task_tracker/constants"
	"task_tracker/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(data *domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	sql := fmt.Sprintf("INSERT INTO %s(id, login, password, firstname, middlename, lastname, is_active, account_disable_time, system_role, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", constants.UserTable)
	_, err = tx.Exec(sql, data.Id, data.Login, data.Password, data.Firstname, data.Middlename, data.Lastname, data.IsActive, data.AccountDisableTime, data.SystemRole, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, department := range data.Departments {
		sql := fmt.Sprintf("INSERT INTO %s(user_id, department_id) VALUES($1, $2)", constants.UserDepartmentTable)
		_, err := tx.Exec(sql, data.Id, department)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, position := range data.Positions {
		sql := fmt.Sprintf("INSERT INTO %s(user_id, position_id) VALUES($1, $2)", constants.UserPositionTable)
		_, err := tx.Exec(sql, data.Id, position)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, specialization := range data.Specializations {
		sql := fmt.Sprintf("INSERT INTO %s(user_id, specialization_id) VALUES($1, $2)", constants.UserSpecializationTable)
		_, err := tx.Exec(sql, data.Id, specialization)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	return err
}

func (r *UserRepository) GetUserById(id uuid.UUID) (*domain.User, error) {
	data := new(domain.User)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", constants.UserTable)
	err := r.db.Get(&data, sql, id)
	return data, err
}

func (r *UserRepository) GetUserDTOById(id uuid.UUID) (*domain.UserDTO, error) {
	data := new(domain.UserDTO)
	sql := fmt.Sprintf("SELECT u.id, u.login, u.firstname, u.middlename, u.lastname, u.is_active, u.account_disable_time, u.created_at,"+
		"CASE WHEN system_role = 1 THEN 'ROLE_USER' ELSE 'ROLE_ADMIN' END AS system_role "+
		"FROM %s u "+
		"WHERE u.id = $1", constants.UserTable)
	fmt.Println(id)
	err := r.db.Get(data, sql, id)
	return data, err
}
