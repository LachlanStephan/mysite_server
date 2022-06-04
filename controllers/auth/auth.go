package auth

import (
	"fmt"

	db "mysite_server/db"

	"golang.org/x/crypto/bcrypt"
)

// check if admin
func CheckAdmin(pass string) int {
	db.Connect()

	fmt.Println("do we reach?", pass)

	passToCheck := []byte(pass)

	var p []byte

	sql := "SELECT password FROM user WHERE user_name = ?"

	rows, err := db.Conn.Query(sql, "LachTwo")

	if err != nil {
		fmt.Println(err)
		return 403
	}

	fmt.Println(rows)

	db.Close()

	fmt.Println("do we reach? 2")

	for rows.Next() {
		err := rows.Scan(&p)

		if err != nil {
			fmt.Println("scan err", err)
			return 403
		}
	}

	var t = []byte("helloThere")
	test, err := bcrypt.GenerateFromPassword(t, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("generate err", test)
	}

	fmt.Println(test)

	// hash the attempt - posted password
	hash, err := bcrypt.GenerateFromPassword(passToCheck, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "something wrong with hash")
		return 500
	}

	fmt.Println(hash, p)

	// compare the two
	err = bcrypt.CompareHashAndPassword(hash, p)

	switch {
	case err != nil:
		fmt.Println("compare shitted")
		return 500
	default: // if err == nil - passwords matched
		fmt.Println("worked")
		return 200
	}
}
