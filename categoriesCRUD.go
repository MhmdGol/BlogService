package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createNewCategory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var category Category
	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		panic(err)
	}

	db.Create(&category)
	fmt.Fprintf(w, "Category inserted successfully!")
}

func readAllCategories(w http.ResponseWriter, r *http.Request) {
	var allCategories []Category

	db.Find(&allCategories)
	json.NewEncoder(w).Encode(allCategories)
}

// func readACategory(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key, _ := strconv.Atoi(vars["id"])

// 	var category Category
// 	db.First(&category, key)

// 	if category.ID == 0 {
// 		fmt.Fprintf(w, "Category not found!")
// 	} else {
// 		json.NewEncoder(w).Encode(category)
// 	}
// }

func updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updatedCategory Category
	json.Unmarshal(reqBody, &updatedCategory)

	var foundCategory Category
	db.First(&foundCategory, key)

	if foundCategory.ID == 0 {
		fmt.Fprintf(w, "Category not found!")
	} else {
		foundCategory.Name = updatedCategory.Name
		db.Save(&foundCategory)
		fmt.Fprintf(w, "Category updated successfully!")
	}
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	var category Category
	db.First(&category, key)

	if category.ID == 0 {
		fmt.Fprintf(w, "Category not found!")
	} else {
		db.Delete(&category)
		fmt.Fprintf(w, "Category deleted successfully!")
	}
}
