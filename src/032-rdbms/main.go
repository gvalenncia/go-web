package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB
var err error

func main()  {
	db, err = sql.Open(
		"mysql",
		"root:root@tcp(localhost:3306)/amigos")
	check(err)
	defer db.Close()

	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/insert", insertAmigo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func insertAmigo(w http.ResponseWriter, req *http.Request) {

	stmt, err := db.Prepare("INSERT INTO `tbl_amigo` (`first_name`) VALUES ('jose');")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()

	fmt.Fprintf(w,string(n))
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT first_name from tbl_amigo;")
	check(err)

	var firstName string
	for rows.Next() {
		err = rows.Scan(&firstName)
		fmt.Fprintf(w, firstName)
	}
}

func check(e error) {
	if err != nil {
		fmt.Println(e)
	}
}