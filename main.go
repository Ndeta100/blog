package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/controllers"
	"myapp/models"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	controllers.Sessions = make(map[string]*models.User)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//pwd := os.Getenv("PASSWORD")

	// Never use _, := db.Open(), release resources with db.Close
	//models.Db, models.Err = sql.Open("mysql", "root:"+pwd+"@tcp(127.0.0.1:3306)/dbblog")
	models.Db, models.Err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(models.Err)
	}
	models.Err = models.Db.Ping()
	if models.Err != nil {
		log.Fatal(models.Err)
	}
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/dashboard", controllers.Dashboard)
	http.HandleFunc("/", controllers.Index)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	log.Println(fmt.Sprintf("Your app is running on port %s.", os.Getenv("PORT")))
	log.Println(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
