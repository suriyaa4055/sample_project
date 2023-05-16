package db

import (
	"encoding/json"
	"example/goAPIs/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	var employee model.Book
	var response model.Response
	var arrEmployee []model.Book

	con := Connect()
	defer con.Close()

	rows, err := con.Query("SELECT id, title, author FROM books")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&employee.ID, &employee.Title, &employee.Author)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrEmployee = append(arrEmployee, employee)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrEmployee

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// func InsertEmployee1(w http.ResponseWriter, r *http.Request) {
// 	var response model.Response
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	con := Connect()
// 	defer con.Close()
// 	var tag model.Book
// 	json.Unmarshal(reqBody, &tag)
// 	w.Header().Set("Content-Type", "application/json")

// 	// err := r.ParseMultipartForm(4096)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	id := tag.ID
// 	title := tag.Title
// 	author := tag.Author
// 	quantity := tag.Quantity

// 	_, err := con.Exec("INSERT INTO books(id, title, author,quantity) VALUES(?, ?, ?, ?)", id, title, author, quantity)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	response.Status = 200
// 	response.Message = "Insert data successfully"
// 	fmt.Print("Insert data to database")
// 	//response.Data  tag

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

func InsertEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Book
	var response model.Response
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &employee)
	//var arrEmployee []model.Book

	con := Connect()
	defer con.Close()

	rows, err := con.Query("insert into books(id, title, author, quantity) VALUES(?, ?, ?, ?)", employee.ID, employee.Title, employee.Author, employee.Quantity)

	if err != nil {
		log.Print(err)
	}
	err = rows.Scan(&employee.ID, &employee.Title, &employee.Author)
	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")
	response.Data = append(response.Data, employee)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	var response model.Response

	response.Status = 200
	response.Message = "Hi User now can look at books using /books endpoint just updated on 5/16/23"
	fmt.Print("table created in database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
