package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pawutj/go_hexagonal/repository"
	"github.com/spf13/viper"
)

func main() {
	url := fmt.Sprintf("%v:%v@tcp:%v", viper.GetString("db.username"), viper.GetString("db.password"))
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)

	customers, err := customerRepository.GetAll()

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
