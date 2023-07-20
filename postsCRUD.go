package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Post struct {
	Title string `json:title`
	Text  string `json:text`
}

func createNewPost(w http.ResponseWriter, r *http.Request) {

	// the query on database for creating a post
	sqlQuery := `INSERT INTO posts (title, text, creation_time, modification_time)
				VALUES ($1, $2, $3, $4)`
	reqBody, _ := ioutil.ReadAll(r.Body)

	var post Post
	json.Unmarshal(reqBody, &post)
	Ctime := time.Now().Format("15:04:05")
	Mtime := time.Time{}.Format("00:00:00")

	_, err = db.Exec(sqlQuery, post.Title, post.Text, Ctime, Mtime)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Post successfully inserted")
}
