package controllers

import (
	"encoding/json"
	"fmt"
	"golang-api/pkg/models"
	"golang-api/pkg/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	Signin    models.User
	secretkey string = "secretkeyjwt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	Signin := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&Signin)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	dbuser, _ := models.ValidateUsername(Signin.Username)

	fmt.Println(dbuser.Username)
	//checks if email is already register or not
	if dbuser.Username != "" {
		var err Error
		err = SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	Signin.Password, err = GeneratehashPassword(Signin.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	utils.ParseBody(r, Signin)
	b := Signin.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var authDetails models.User

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	authUser, _ := models.ValidateLogin(authDetails.Username)

	if authUser.Username == "" {
		var err Error
		fmt.Println(err)
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err Error
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Username)
	if err != nil {
		var err Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token models.Token
	token.Username = authUser.Username

	token.TokenString = validToken
	utils.ParseBody(r, authDetails)

	userdetail, db := models.GetUserById(token.Username)

	userdetail.Token = token.TokenString

	fmt.Println(userdetail.Token)
	db.Save(&userdetail)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email

	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
