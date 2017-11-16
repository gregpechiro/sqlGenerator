/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/sqlGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"log"
)

func CreateCreditCardsTable() {
	_, err := db.Exec("CREATE TABLE credit-card()")
	if err != nil {
		fmt.Printf("CreateCreditCardsTable >> %v\n", err)
	}
}

func AddCreditCard(creditCard CreditCard) {
	_, err := db.Exec("INSERT INTO user VALUES( ?, ?, ?, ?, ?, ?, ? )", creditCard.Id, creditCard.UserId, creditCard.Number, creditCard.ExpirationDate, creditCard.SecurityCode, creditCard.NameOnCard, creditCard.Typ)
	if err != nil {
		fmt.Printf("AddCreditCard >> %v\n", err)
	}
}

func GetAllCreditCards() []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card")
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardById(id int) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE id=?", id).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetIdById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsById(id int) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE id=?", id)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardByUserId(userId int) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE user-id=?", userId).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetUserIdById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsByUserId(userId int) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE user-id=?", userId)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardByNumber(number string) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE number=?", number).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetNumberById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsByNumber(number string) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE number=?", number)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardByExpirationDate(expirationDate string) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE expiration-date=?", expirationDate).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetExpirationDateById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsByExpirationDate(expirationDate string) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE expiration-date=?", expirationDate)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardBySecurityCode(securityCode string) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE security-code=?", securityCode).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetSecurityCodeById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsBySecurityCode(securityCode string) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE security-code=?", securityCode)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardByNameOnCard(nameOnCard string) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE name-on-card=?", nameOnCard).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetNameOnCardById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsByNameOnCard(nameOnCard string) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE name-on-card=?", nameOnCard)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}

func GetCreditCardByTyp(typ string) CreditCard {
	creditCard := CreditCard{}
	err := db.QueryRow("SELECT * FROM credit-card WHERE typ=?", typ).Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
	if err != nil {
		fmt.Printf("GetTypById >> %v\n", err)
	}
	return creditCard
}

func GetAllCreditCardsByTyp(typ string) []CreditCard {
	rows, err := db.Query("SELECT * FROM credit-card WHERE typ=?", typ)
	if err != nil {
		log.Fatalf("Query: %v\n", err)
	}
	var creditCards []CreditCard
	for rows.Next() {
		var creditCard CreditCard
		err = rows.Scan(&creditCard.Id, &creditCard.UserId, &creditCard.Number, &creditCard.ExpirationDate, &creditCard.SecurityCode, &creditCard.NameOnCard, &creditCard.Typ)
		if err != nil {
			log.Fatalf("Scan: %v", err)
		}
		creditCards = append(creditCards, creditCard)
	}
	return creditCards
}
