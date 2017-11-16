/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/sqlGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"log"
)

func CreateUsersTable() {
	_, err := db.Exec("CREATE TABLE user()")
	if err != nil {
		fmt.Printf("CreateUsersTable >> %v\n", err)
	}
}

func AddUser(user User) {
	_, err := db.Exec("INSERT INTO user VALUES( ?, ?, ?, ?, ? )", user.Id, user.Name, user.Age, user.Active, user.Happy)
	if err != nil {
		fmt.Printf("AddUser >> %v\n", err)
	}
}

func GetAllUsers() []User {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserById(id int) User {
	user := User{}
	err := db.QueryRow("SELECT * FROM user WHERE id=?", id).Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
	if err != nil {
		fmt.Printf("GetIdById >> %v\n", err)
	}
	return user
}

func GetAllUsersById(id int) []User {
	rows, err := db.Query("SELECT * FROM user WHERE id=?", id)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserByName(name string) User {
	user := User{}
	err := db.QueryRow("SELECT * FROM user WHERE name=?", name).Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
	if err != nil {
		fmt.Printf("GetNameById >> %v\n", err)
	}
	return user
}

func GetAllUsersByName(name string) []User {
	rows, err := db.Query("SELECT * FROM user WHERE name=?", name)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserByAge(age int) User {
	user := User{}
	err := db.QueryRow("SELECT * FROM user WHERE age=?", age).Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
	if err != nil {
		fmt.Printf("GetAgeById >> %v\n", err)
	}
	return user
}

func GetAllUsersByAge(age int) []User {
	rows, err := db.Query("SELECT * FROM user WHERE age=?", age)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserByActive(active bool) User {
	user := User{}
	err := db.QueryRow("SELECT * FROM user WHERE active=?", active).Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
	if err != nil {
		fmt.Printf("GetActiveById >> %v\n", err)
	}
	return user
}

func GetAllUsersByActive(active bool) []User {
	rows, err := db.Query("SELECT * FROM user WHERE active=?", active)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}

func GetUserByHappy(happy bool) User {
	user := User{}
	err := db.QueryRow("SELECT * FROM user WHERE happy=?", happy).Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
	if err != nil {
		fmt.Printf("GetHappyById >> %v\n", err)
	}
	return user
}

func GetAllUsersByHappy(happy bool) []User {
	rows, err := db.Query("SELECT * FROM user WHERE happy=?", happy)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Active, &user.Happy)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		users = append(users, user)
	}
	return users
}
