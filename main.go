package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pawutj/go_hexagonal/repository"
)

func main() {
	url := ""
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)

}
