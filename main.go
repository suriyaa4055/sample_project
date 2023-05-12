package main

import (
	"example/goAPIs/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func getBookById(id string) (*book, error) {
// 	for i, b := range books {
// 		if b.ID == id {
// 			return &books[i], nil
// 		}
// 	}
// 	return nil, errors.New("book not found")
// }

// func dummy() {
// 	con := db.Connect()
// 	defer con.Close()

// 	// var response model.Response

// 	_, err := con.Query("create table books (id varchar(255),title varchar(255),author varchar(255),quantity int)")

// 	fmt.Println("dummy func")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	fmt.Println("inside main func")
	//dummy()
	router := mux.NewRouter()
	router.HandleFunc("/", db.Hello).Methods("GET")
	router.HandleFunc("/books", db.AllBooks).Methods("GET")
	router.HandleFunc("/books", db.InsertEmployee).Methods("POST")
	// router.GET("/book/:id", bookById)
	// router.POST("/books", createBooks)
	// router.PATCH("/checkout", checkoutBook)
	// router.PATCH("/return", returnBook)
	fmt.Print("server started")
	http.ListenAndServe(":8080", router)

}

// type aggregatedLogger struct {
// 	infoLogger  *log.Logger
// 	warnLogger  *log.Logger
// 	errorLogger *log.Logger
// }

// func (l *aggregatedLogger) info(v ...interface{}) {
// 	l.infoLogger.Println(v...)
// }

// func (l *aggregatedLogger) warn(v ...interface{}) {
// 	l.warnLogger.Println(v...)
// }

// func (l *aggregatedLogger) error(v ...interface{}) {
// 	l.errorLogger.Println(v...)
// }
