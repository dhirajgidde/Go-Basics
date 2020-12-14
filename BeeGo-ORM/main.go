package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int    `json:"id"` //default field
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Marks   int    `json:"marks"`
}

var jsonString []byte

func init() {
	orm.RegisterModel(new(User))                                                    //Register A ORM
	orm.RegisterDriver("mysql", orm.DRMySQL)                                        //Register A Driver
	orm.RegisterDataBase("default", "mysql", "root:dev@tcp(127.0.0.1:3306)/dhiraj") //Register A DataBase
}

func fetchDetails() {
	o := orm.NewOrm()
	user := User{Id: 8}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("No Result Found")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key is found")
	} else {
		fmt.Println(user)

	}
	Bytearr, err := json.Marshal(user)
	jsonString = Bytearr
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(string(Bytearr))
	}
}

func insertQuery() {
	o := orm.NewOrm()
	var user User
	user.Id = 8
	user.Name = "Tata"
	user.Subject = "Marathi"
	user.Marks = 450

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
		fmt.Println("Record Inserted")

	} else {
		panic(err.Error())
	}

}

func insertMultiple() {
	o := orm.NewOrm()
	//var user User
	user := []User{
		{
			Id:      9,
			Name:    "Roshan",
			Subject: "Geogrphy",
			Marks:   86,
		},
		{
			Id:      10,
			Name:    "Mogambo",
			Subject: "Aljebra",
			Marks:   96,
		},
	}
	inserted, err := o.InsertMulti(2, user)
	if err == nil {
		fmt.Println(inserted)
		fmt.Println("Record inserted Successfully")

	}
}

func updateQuery() {
	o := orm.NewOrm()
	user := User{Id: 1}
	if o.Read(&user) == nil {
		user.Subject = "English"
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
			fmt.Println("Record Updated")
		}
	}

}

func deleteQuery() {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: 1}); err == nil {
		fmt.Println(num)
		fmt.Println("Record Deleted")
	}
}

func indexElement(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(jsonString))
}

func connectServer() {
	http.HandleFunc("/", indexElement)
	http.ListenAndServe(":9085", nil)
}

// func applyingFilter() {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("user")
// 	res := qs.Filter("name__contains", "Ta")
// 	fmt.Println(res)
// }

func main() {
	fetchDetails()
	//insertQuery()
	//deleteQuery()
	//updateQuery()
	//insertMultiple()
	//applyingFilter()

	connectServer()
}
