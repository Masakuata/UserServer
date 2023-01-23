package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type User struct {
	Id       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty" validate:"required" binding:"required"`
	Password string `json:"password,omitempty" validate:"required" binding:"required"`
}

var dbConnString, _ = os.LookupEnv("DBCONNSTRING")
var conn, _ = sql.Open("mysql", dbConnString)

func Register(email, password string) (int, User) {

	if IsRegistered(email) {
		return http.StatusConflict, User{}
	}

	var query string = "INSERT INTO User.User (email, password) VALUES (?, ?)"
	exec, err := conn.Exec(query, email, password)
	if err != nil {
		return http.StatusInternalServerError, User{}
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, User{}
	}
	var newUser User = User{Id: id, Email: email, Password: password}
	return http.StatusCreated, newUser
}

func IsRegistered(email string) bool {
	var query string = "SELECT COUNT(*) AS TOTAL FROM User.User WHERE email = ?"
	result := conn.QueryRow(query, email)
	var total uint8 = 0
	err := result.Scan(&total)
	if err != nil {
		return false
	}
	return total == 1
}

func Login(email, password string) bool {
	var query string = "SELECT COUNT(*) AS TOTAL FROM User.User WHERE email = ? AND password = ? AND status = 1"
	result := conn.QueryRow(query, email, password)
	var total uint8 = 0
	err := result.Scan(&total)
	if err != nil {
		return false
	}
	return total == 1
}
