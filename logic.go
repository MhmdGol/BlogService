package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var secretKey = []byte("Blogloglog")

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_service"
)

type Category struct {
	gorm.Model
	Name  string  `json:"name"`
	Posts []*Post `gorm:"many2many:post_categories;"`
}

type Post struct {
	gorm.Model
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	Categories []*Category `json:"cats" gorm:"many2many:post_categories;"`
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection to the database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//=====================================

	err = db.AutoMigrate(&Category{}, &Post{})
	if err != nil {
		panic(err)
	}

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/login", loginHandler).Methods("POST")

	myRouter.HandleFunc("/post/create", createNewPost).Methods("POST")
	myRouter.HandleFunc("/post/read", readAllPosts).Methods("GET")
	myRouter.HandleFunc("/post/read/{id}", readPostByPaging).Methods("GET")
	myRouter.HandleFunc("/post/update/{id}", updatePost).Methods("PUT")
	myRouter.HandleFunc("/post/delete/{id}", deletePost).Methods("DELETE")

	myRouter.HandleFunc("/category/create", createNewCategory).Methods("POST")
	myRouter.HandleFunc("/category/read", readAllCategories).Methods("GET")
	// myRouter.HandleFunc("/category/read/{id}", readACategory).Methods("GET")
	myRouter.HandleFunc("/category/update/{id}", updateCategory).Methods("PUT")
	myRouter.HandleFunc("/category/delete/{id}", deleteCategory).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

type User struct {
	ID       int
	Username string `json:"user"`
	Password string `json:"pass"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// ------ Hardcoded user ------
	allowedUser := User{
		ID:       1,
		Username: "Mhmd",
		Password: "1234",
	}
	// ----------------------------

	// login check
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)

	if user.Username != allowedUser.Username || user.Password != allowedUser.Password {
		fmt.Fprintf(w, "Not Authenticated!")
		return
	}

	// here we generate a token for the user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic(err)
	}

	// give the token to the user
	w.Write([]byte(tokenString))
}

func checkAuthentication(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
