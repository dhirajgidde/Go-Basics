package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	no      int64
	name    string
	subject string
	marks   int64
}

func insertQuery(db *sql.DB) {
	no := 5
	name := "Shubham"
	subject := "Geo"
	marks := 100
	insert, err := db.Query("INSERT INTO student values(?,?,?,?)", no, name, subject, marks)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Reocord Inserted Successfully")
	}

	defer insert.Close()
}

func deleteQuery(db *sql.DB) {
	delete, err := db.Query("DELETE FROM student where no=1")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Reocord deleted Successfully")
	}
	defer delete.Close()
}

func fetchRecord(db *sql.DB) {
	fetch, err := db.Query("SELECT * FROM student")
	// no:=5
	// fetch, err := db.Query("SELECT * FROM student where no=?",no)
	if err != nil {
		panic(err.Error())
	}
	stud := student{}
	res := []student{}

	for fetch.Next() {
		var no, mark int64
		var name, subject string
		err := fetch.Scan(&no, &name, &subject, &mark)
		if err != nil {
			panic(err.Error())
		}
		stud.no = no
		stud.name = name
		stud.subject = subject
		stud.marks = mark
		res = append(res, stud)
	}

	fmt.Println(res)
}

func main() {
	fmt.Println("GO My-Sql Tutorial....")

	dbDriver := "mysql"
	dbUser := "dhiraj"
	dbPass := "dhiraj"
	dbDataBase := "dhiraj"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbDataBase)
	defer db.Close()

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connected Succefully")
	}

	insertQuery(db)
	//deleteQuery(db)
	//fetchRecord(db)
}
