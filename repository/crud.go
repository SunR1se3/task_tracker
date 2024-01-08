package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CRUDRepository struct {
	db *sqlx.DB
}

func NewCRUDRepository(db *sqlx.DB) *CRUDRepository {
	return &CRUDRepository{db: db}
}

func (r *CRUDRepository) Create(fields []string, values []any, tableName string) error {
	preparedFields := strings.Join(fields, ",")
	preparedValues := ""
	for i, _ := range values {
		if i != len(values)-1 {
			preparedValues += fmt.Sprintf("$%d, ", i+1)
		} else {
			preparedValues += fmt.Sprintf("$%d", i+1)
		}

	}
	//preparedValues = strings.Join(strings.Split(preparedValues, " "), ",")
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", tableName, preparedFields, preparedValues)
	_, err := r.db.Exec(sql, values...)
	return err
}
