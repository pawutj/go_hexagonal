package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {

	customers := []Customer{}
	query := "select * from customers"

	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil

}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {

	customer := Customer{}
	err := r.db.Get(&customer, "select * from customers where id =?", id)

	if err != nil {
		return nil, err
	}

	return &customer, nil

}
