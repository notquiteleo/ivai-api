package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

   type User struct {
       ID    int    `json:"id"`
       Name  string `json:"name"`
       Email string `json:"email"`
   }

   var (
       DB *sqlx.DB
   )

   func main() {
       // 初始化数据库连接
       err := initDB()
       if err != nil {
           panic(err)
       }

       r := chi.NewRouter()
       r.Get("/users", handleUsers)

       http.ListenAndServe(":8080", r)
   }

   func initDB() error {
       var err error
       DB, err = sqlx.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
       if err != nil {
           return err
       }
       return nil
   }

   func handleUsers(w http.ResponseWriter, r *http.Request) {
       var users []User
       err := queryAllUsers(&users)
       if err != nil {
           http.Error(w, err.Error(), http.StatusInternalServerError)
           return
       }

       json.NewEncoder(w).Encode(users)
   }

	 func queryAllUsers(users *[]User) error {
    // Implement the database query to populate the users slice
    // Example query: SELECT * FROM users
    query := "SELECT * FROM users"
    return DB.Select(users, query)
}
