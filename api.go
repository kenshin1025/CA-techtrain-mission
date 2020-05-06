package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB

type userJSON struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)

		var user userJSON
		json.Unmarshal(buf.Bytes(), &user)
		
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "POST hello! %v\n", user)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT * FROM user") //
		if err != nil {
			panic(err.Error())
		}

		columns, err := rows.Columns() // カラム名を取得
		if err != nil {
			panic(err.Error())
		}
		fmt.Fprintf(w, "Hello get%v\n", columns)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		fmt.Fprintf(w, "Hello put\n")
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func main() {
	var err error
	fmt.Printf("Starting server at 'localhost:8080'\n")
	db, err = sql.Open("mysql", "root:@/ca_mission")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	http.HandleFunc("/", handler)
	http.HandleFunc("/user/create", create)
	http.HandleFunc("/user/get", get)
	http.HandleFunc("/user/update", update)
	http.ListenAndServe(":8080", nil)
}
