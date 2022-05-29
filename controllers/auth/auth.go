package auth

import (
	"fmt"

	db "mysite_server/db"

	"golang.org/x/crypto/bcrypt"
)

// check if admin
func CheckAdmin(pass string) int {
	db.Connect()

	passToCheck := []byte(pass)

	// pass from db
	var p []byte

	sql := "SELECT password FROM user WHERE user_name = Lach"

	rows, err := db.Conn.Query(sql)

	if err != nil {
		return 403
	}

	db.Close()

	for rows.Next() {
		err := rows.Scan(&p)

		if err != nil {
			fmt.Println("oops", err)
			return 403
		}
	}

	var t = []byte("helloThere")
	test, err := bcrypt.GenerateFromPassword(t, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(test)
	}

	// hash the attempt - posted password
	hash, err := bcrypt.GenerateFromPassword(passToCheck, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "something wrong with hash")
		return 500
	}

	// compare the two
	err = bcrypt.CompareHashAndPassword(hash, p)

	switch {
	case err != nil:
		return 500
	default: // if err == nil - passwords matched
		return 200
	}
}
