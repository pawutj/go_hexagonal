package repository

type Customer struct {
	CustomerID int    `db:"customer_id"`
	Name       string `db:"name"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	// * for return nil
	GetById(int id) (*Customer, error)
}
