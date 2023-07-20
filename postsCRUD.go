package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DataPost struct {
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Cats  []string `json:"cats"`
}

type PageSize struct {
	Size int `json:"size"`
}

func createNewPost(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var datapost DataPost
	err := json.Unmarshal(reqBody, &datapost)
	if err != nil {
		panic(err)
	}

	datapost.Cats = deDuplicate(datapost.Cats)
	if len(datapost.Cats) > 6 {
		panic("At most 6 categories are allowed")
	}

	// here we make a post object to be inserted to the database.
	// we fill the title and text by the information we get
	// in the for loop part, we check the cat list
	// if the category is in tables we get the referene to it
	// (it prevents duplication) or it is not present
	// then we make a new cat object and insert it to the categories too
	var post Post
	post.Title = datapost.Title
	post.Text = datapost.Text

	for _, item := range datapost.Cats {
		var findCat Category
		db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			post.Categories = append(post.Categories, &Category{Name: item})
		} else {
			post.Categories = append(post.Categories, &findCat)
		}
	}

	db.Create(&post)
	fmt.Fprintf(w, "Posted Successfully!")
}

func deDuplicate(slice []string) []string {
	uniqueMap := make(map[string]bool)
	deduplicatedSlice := []string{}

	for _, item := range slice {
		if !uniqueMap[item] {
			deduplicatedSlice = append(deduplicatedSlice, item)
			uniqueMap[item] = true
		}
	}

	return deduplicatedSlice
}

func readAllPosts(w http.ResponseWriter, r *http.Request) {
	var allPosts []Post

	db.Preload("Categories").Find(&allPosts)
	json.NewEncoder(w).Encode(allPosts)
}

func readPostByPaging(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["id"])
	if page < 1 {
		panic("paging starts at 1")
	}

	var pageSize PageSize
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &pageSize)

	// sorting posts by their update_at
	// then reading a portion of them
	var posts []Post
	db.Order("updated_at desc").Offset((page - 1) * pageSize.Size).Limit(pageSize.Size).Find(&posts)

	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// placing currently present post to postToUpdate
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	var postToUpdate Post
	db.Preload("Categories").Where("id = ?", key).First(&postToUpdate)

	if postToUpdate.ID == 0 {
		panic("Post doesnt exist!")
	}

	// given information to update the post
	reqBody, _ := ioutil.ReadAll(r.Body)

	var datapost DataPost
	err := json.Unmarshal(reqBody, &datapost)
	if err != nil {
		panic(err)
	}

	datapost.Cats = deDuplicate(datapost.Cats)
	if len(datapost.Cats) > 6 {
		panic("At most 6 categories are allowed")
	}

	// updating by newly given informations
	postToUpdate.Title = datapost.Title
	postToUpdate.Text = datapost.Text
	for _, item := range datapost.Cats {
		var findCat Category
		db.Where("name = ?", item).First(&findCat)

		if findCat.ID == 0 {
			postToUpdate.Categories = append(postToUpdate.Categories, &Category{Name: item})
		} else {
			postToUpdate.Categories = append(postToUpdate.Categories, &findCat)
		}
	}

	// go to database
	db.Save(&postToUpdate)
	fmt.Fprintf(w, "Post updated successfully!")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	var post Post
	db.First(&post, key)

	if post.ID == 0 {
		fmt.Fprintf(w, "Post not found!")
	} else {
		db.Delete(&post)
		fmt.Fprintf(w, "Post deleted successfully!")
	}
}
