package controller

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB
// var err error

// const DNS = "root:aryangupta05@tcp(127.0.0.1:3306)/aryan?charset=utf8mb4&parseTime=True&loc=Local"

// type User struct {
// 	gorm.Model
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	Email     string `json:"email"`
// }

// func IntializeMigration() {
// 	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect to database!")
// 	}
// 	DB.AutoMigrate(&User{})
// }

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var users []User
// 	DB.Find(&users)
// 	json.NewEncoder(w).Encode(users)
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(user)
// }

// func CreateUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Create(&user)
// 	json.NewEncoder(w).Encode(user)
// }

// func UpdateUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.First(&user, params["id"])
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Save(&user)
// 	json.NewEncoder(w).Encode(user)
// }

// func DeleteUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User
// 	DB.First(&user, params["id"])
// 	json.NewDecoder(r.Body).Decode(&user)
// 	DB.Delete(&user)
// 	json.NewEncoder(w).Encode("User Deleted")
// }
