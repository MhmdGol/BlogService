package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// func createNewCategory(w http.ResponseWriter, r *http.Request) {
// 	sqlQuery := `INSERT INTO categories (name) VALUES ($1)`

// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var category Category
// 	json.Unmarshal(reqBody, &category)

// 	_, err = db.Exec(sqlQuery, category.Name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprintf(w, "Category successfully inserted")
// }

// func readAllCategories(w http.ResponseWriter, r *http.Request) {
// 	sqlQuery := `SELECT name FROM categories`

// 	rows, err := db.Query(sqlQuery)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var categories []Category
// 	for rows.Next() {
// 		var category Category
// 		err = rows.Scan(&category.Name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		categories = append(categories, category)
// 	}
// 	json.NewEncoder(w).Encode(categories)
// }

// func readACategory(w http.ResponseWriter, r *http.Request) {
// 	sqlQuery := `SELECT name FROM categories WHERE id=$1`

// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])
// 	rows, err := db.Query(sqlQuery, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	var category Category
// 	rows.Next()
// 	err = rows.Scan(&category.Name)
// 	if err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(w).Encode(category)
// }

// func updateCategory(w http.ResponseWriter, r *http.Request) {
// 	sqlQuery := `
// 		UPDATE categories
// 		SET name=$1
// 		WHERE id=$2
// 	`
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var category Category
// 	json.Unmarshal(reqBody, &category)

// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	_, err := db.Exec(sqlQuery, category.Name, key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintf(w, "Categories successfully updated")
// }

// func deleteCategory(w http.ResponseWriter, r *http.Request) {
// 	sqlQuery := `
// 		DELETE FROM categories WHERE id=$1
// 	`
// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	_, err := db.Exec(sqlQuery, key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Fprintf(w, "Categories successfully deleted")
// }
