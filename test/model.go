package main

import "database/sql"

var db *sql.DB

//go:generate go run ../main/main.go User
type User struct {
	Id     int
	Name   string
	Age    int
	Active bool
	Happy  bool
}

//go:generate go run ../main/main.go Address
type Address struct {
	Id     int
	UserId int
	Street string
	City   string
	State  string
	zip    string
}

//go:generate go run ../main/main.go CreditCard
type CreditCard struct {
	Id             int
	UserId         int
	Number         string
	ExpirationDate string
	SecurityCode   string
	NameOnCard     string
	Typ            string
}
