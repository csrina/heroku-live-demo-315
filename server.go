package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	// "github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	// "database/sql"
)

// var db *sql.DB

// func createRouter() *mux.Router {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", serveIndex).Methods("GET")
// 	r.HandleFunc("/cat", serveCat).Methods("GET")
// 	r.HandleFunc("/dance", serveDance).Methods("GET")
// 	r.HandleFunc("/what", serveHack).Methods("GET")
// 	r.HandleFunc("/getName", getName).Methods("GET")

// 	return r
// }

// func getName(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("Select * from names")
// 	if err != nil {
// 		w.Write([]byte("Error getting names"))
// 		return
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var first string
// 		if err := rows.Scan(&first); err != nil {
// 			w.Write([]byte("error"))
// 			return
// 		}
// 	w.Write([]byte("Read from db:"+ first + "\n"))

// 	}
// }

// func serveIndex(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "index.html")
// }

// func serveCat(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "cat.html")
// }

// func serveDance(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "dance.html")
// }

// func serveHack(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "hack.html")
// }
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// db, _= sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// log.Println("database openned")

	// r := createRouter()
	// http.Handle("/", r)
	http.HandleFunc("/", rootHandler)
	log.Println("starting server.")
	log.Println(port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
