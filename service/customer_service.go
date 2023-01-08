package service

import (
	"log"

	"github.com/pawutj/go_hexagonal/repository"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{customerRepository: customerRepository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	customersResponse := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{CustomerId: customer.CustomerID, Name: customer.Name}
		customersResponse = append(customersResponse, customerResponse)
	}

	return customersResponse, nil

}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	customerResponse := CustomerResponse{CustomerId: customer.CustomerID, Name: customer.Name}

	return &customerResponse, nil
}
