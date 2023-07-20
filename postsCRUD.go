package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"
// )

// type Post struct {
// 	Title      string   `json:"title"`
// 	Text       string   `json:"text"`
// 	Categories []string `json:"cats"`
// }

// func createNewPost(w http.ResponseWriter, r *http.Request) {

// 	// the query on database for creating a post
// 	sqlQuery := `INSERT INTO posts (title, text, creation_time, modification_time)
// 					VALUES ($1, $2, $3, $4)`
// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var post Post
// 	json.Unmarshal(reqBody, &post)

// 	if len(post.Categories) > 6 {
// 		panic(errors.New("At most 6 categories are allowed"))
// 	}

// 	Ctime := time.Now().Format("15:04:05")
// 	Mtime := time.Time{}.Format("00:00:00")

// 	_, err = db.Exec(sqlQuery, post.Title, post.Text, Ctime, Mtime)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// here we extract category ids
// 	cats := []int{}

// 	for _, cat := range post.Categories {
// 		_, err = db.Exec()
// 	}

// 	// json.NewEncoder(w).Encode(post)
// 	fmt.Fprintf(w, "Post successfully inserted")

// }

// func listOfCatIds(cats []string) []int {
// 	sqlQuery := `
// 		SELECT id
// 		FROM categories
// 		WHERE name=$1
// 	`
// 	for _, cat := range cats {
// 		var category Category

// 		rows, err := db.Query(sqlQuery, cat)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer rows.Close()

// 		rows.Next()
// 		err = rows.Scan(&category.Name)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }
