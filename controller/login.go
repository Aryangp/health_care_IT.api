package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Aryangp/goRest/database"
	"github.com/Aryangp/goRest/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	connection := database.GetDatabase()
	// defer Closedatabase(connection)

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err error
		err = errors.New("Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	var dbuser model.User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//checks if email is already register or not
	if dbuser.Email != "" {
		var err error
		err = errors.New("Email already register")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//insert user details in database
	connection.Create(&user)

	err2 := json.NewEncoder(w).Encode(user)

	if err2 != nil {
		log.Fatal(err2)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("error occure while sending data"))
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	// defer CloseDatabase(connection)

	var authDetails model.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err error
		err = errors.New("Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser model.User
	connection.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
		var err error
		err = errors.New("Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err error
		err = errors.New("Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Role, authUser.Email)
	if err != nil {
		var err error
		err = errors.New("Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token model.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}

type Message struct {
	MessageString string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("welcome to the health app")
	var message Message
	message.MessageString = "welcome to the health app"
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("error occure while sending data"))
	}
}
