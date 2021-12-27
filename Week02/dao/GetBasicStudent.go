package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	no     int //student number
	name   string
	gender string
	age    int
	dNo    int // department number
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql",
		"bnswt:123456Aa*@tcp(localhost)/db_lab")
	if err != nil {
		fmt.Printf("can not open database!\n")
		panic(err)
	}
}

func GetBasicStudentByNo(no int) (Student, error) {
	var student Student
	row := Db.QueryRow("select * from students1 where no = ?", no)
	err := row.Scan(&student.no, &student.name)
	if err != nil {
		return student, fmt.Errorf("row not found in database: %v", err)
	}
	return student, nil
}
