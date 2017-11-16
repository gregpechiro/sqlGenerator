/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/sqlGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"log"
)

func CreateAddresssTable() {
	_, err := db.Exec("CREATE TABLE address()")
	if err != nil {
		fmt.Printf("CreateAddresssTable >> %v\n", err)
	}
}

func AddAddress(address Address) {
	_, err := db.Exec("INSERT INTO user VALUES( ?, ?, ?, ?, ?, ? )", address.Id, address.UserId, address.Street, address.City, address.State, address.zip)
	if err != nil {
		fmt.Printf("AddAddress >> %v\n", err)
	}
}

func GetAllAddresss() []Address {
	rows, err := db.Query("SELECT * FROM address")
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressById(id int) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE id=?", id).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetIdById >> %v\n", err)
	}
	return address
}

func GetAllAddresssById(id int) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE id=?", id)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressByUserId(userId int) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE user-id=?", userId).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetUserIdById >> %v\n", err)
	}
	return address
}

func GetAllAddresssByUserId(userId int) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE user-id=?", userId)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressByStreet(street string) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE street=?", street).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetStreetById >> %v\n", err)
	}
	return address
}

func GetAllAddresssByStreet(street string) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE street=?", street)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressByCity(city string) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE city=?", city).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetCityById >> %v\n", err)
	}
	return address
}

func GetAllAddresssByCity(city string) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE city=?", city)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressByState(state string) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE state=?", state).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetStateById >> %v\n", err)
	}
	return address
}

func GetAllAddresssByState(state string) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE state=?", state)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}

func GetAddressByzip(zip string) Address {
	address := Address{}
	err := db.QueryRow("SELECT * FROM address WHERE zip=?", zip).Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
	if err != nil {
		fmt.Printf("GetzipById >> %v\n", err)
	}
	return address
}

func GetAllAddresssByzip(zip string) []Address {
	rows, err := db.Query("SELECT * FROM address WHERE zip=?", zip)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var addresss []Address
	for rows.Next() {
		var address Address
		err = rows.Scan(&address.Id, &address.UserId, &address.Street, &address.City, &address.State, &address.zip)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		addresss = append(addresss, address)
	}
	return addresss
}
