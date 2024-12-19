package mysqlconnector

// package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var TableName = "users"

type Tabelinfo struct {
	FirstName    string `json:"firsname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	TicketNumber int    `json:"ticketnumber"`
	TableName    string
}

type SqlConfig struct {
	Password     string
	UserName     string
	MysqlIP      string
	MysqlPort    int
	DatabaseName string
	TableName    string
}

type UsersInfoTable struct {
	Firstnames   []string
	LastName     []string
	Email        []string
	TicketNumber []int
}

func MakeConnectionToDB() *sql.DB {
	SqlConfig := SqlConfig{
		Password:     "test@test",
		UserName:     "root",
		MysqlIP:      "127.0.0.1",
		MysqlPort:    3306,
		DatabaseName: "db",
		TableName:    "users",
	}
	connectioString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", SqlConfig.UserName, SqlConfig.Password, SqlConfig.MysqlIP, SqlConfig.MysqlPort, SqlConfig.DatabaseName)
	db, err := sql.Open("mysql", connectioString)

	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db // Make correct return db.

}

func Insert(firstName string, lastName string, email string, ticketNumber int) {

	tableInfo := Tabelinfo{
		FirstName: `json:"firsname"`,
		Lastname:  `json:"lastname"`,
		Email:     `json:"email"`,
		TableName: TableName,
	}
	db := MakeConnectionToDB()
	defer db.Close()

	var insertQuery = fmt.Sprintf("insert into %v(firsname, lastname, email, ticketnumber) values ('%v', '%v', '%v', %v)", tableInfo.TableName, firstName, lastName, email, ticketNumber)
	insert, err := db.Query(insertQuery)

	if err != nil {
		panic(err.Error())

	}

	fmt.Println(insert.Columns())
	defer insert.Close()

}

func SelectQury() UsersInfoTable {
	var firstNameList []string
	var lastNameList []string
	var emailNameList []string
	var ticketNumberList []int
	// var usersInfosTable Tabelinfo
	var usersTable Tabelinfo

	db := MakeConnectionToDB()
	selectQuery, err := db.Query("select * from users") // For example: db.Query("select * from users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for selectQuery.Next() {
		err = selectQuery.Scan(&usersTable.FirstName, &usersTable.Lastname, &usersTable.Email, &usersTable.TicketNumber)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println("Value from database: ", usersTable.FirstName, usersTable.Lastname, usersTable.Email, usersTable.TicketNumber)
		firstNameList = append(firstNameList, usersTable.FirstName)
		lastNameList = append(lastNameList, usersTable.Lastname)
		emailNameList = append(emailNameList, usersTable.Email)
		ticketNumberList = append(ticketNumberList, usersTable.TicketNumber)
		defer db.Close()
	}
	UsersInfoTable := UsersInfoTable{
		Firstnames:   firstNameList,
		LastName:     lastNameList,
		Email:        emailNameList,
		TicketNumber: ticketNumberList,
	}
	fmt.Println("Len count is: ", len(UsersInfoTable.Firstnames))
	return UsersInfoTable
}
