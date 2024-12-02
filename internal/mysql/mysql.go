package mysqlconnector

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var TableName = "users"

type Tag struct {
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

func Insert(firstName string, lastName string, email string, ticketNumber int) {

	SqlConfig := SqlConfig{
		Password:     "test@test",
		UserName:     "root",
		MysqlIP:      "127.0.0.1",
		MysqlPort:    3306,
		DatabaseName: "db",
		TableName:    "users",
	}

	tableInfo := Tag{
		FirstName: `json:"firsname"`,
		Lastname:  `json:"lastname"`,
		Email:     `json:"email"`,
		TableName: TableName,
	}
	connectioString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", SqlConfig.UserName, SqlConfig.Password, SqlConfig.MysqlIP, SqlConfig.MysqlPort, SqlConfig.DatabaseName)
	db, err := sql.Open("mysql", connectioString)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var insertQuery = fmt.Sprintf("insert into %v(firsname, lastname, email, ticketnumber) values ('%v', '%v', '%v', %v)", tableInfo.TableName, firstName, lastName, email, ticketNumber)
	fmt.Println("Debug: from Insert func", insertQuery)
	insert, err := db.Query(insertQuery)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(insert.Columns())
	defer insert.Close()

	selectQuery, err := db.Query("select * from users")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(selectQuery.Columns())
	for selectQuery.Next() {
		var tag Tag

		// for each row, scan the result into our tag composite object

		err = selectQuery.Scan(&tag.FirstName, &tag.Lastname, &tag.Email, &tag.TicketNumber)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(tag.FirstName, tag.Lastname, tag.Email, tag.TicketNumber)
		defer insert.Close()
	}
}
