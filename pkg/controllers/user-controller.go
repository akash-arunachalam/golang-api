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

func SetError(err models.Token, message string) models.Token {
	err.TokenString = ""
	err.Message = message
	return err
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	

	Signin := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&Signin)
	if err != nil {
		var err models.Token
		err = SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	dbuser, _ := models.ValidateUsername(Signin.Username)

	fmt.Println(dbuser.Username)
	//checks if email is already register or not
	if dbuser.Username != "" {
		var err models.Token
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
	var token models.Token
	token.Message = "User Created Successfully"

	token.TokenString = b.Token
	token.Role = b.Role
	token.Branch = b.Branch

	res, _ := json.Marshal(token)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w)

	var authDetails models.User

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err models.Token
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	authUser, _ := models.ValidateLogin(authDetails.Username)

	if authUser.Username == "" {
		var err models.Token
		fmt.Println(err)
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err models.Token
		err = SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(authUser.Username)

	if err != nil {
		var err models.Token
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(err)
		return
	}

	var token models.Token
	token.Message = "Login Successfully"

	token.TokenString = validToken
	token.Role = authUser.Role
	token.Branch = authUser.Branch
	utils.ParseBody(r, authDetails)

	userdetail, db := models.GetUserByName(authUser.Username)

	userdetail.Token = token.TokenString
	userdetail.Role = authUser.Role
	userdetail.Branch = authUser.Branch

	db.Save(&userdetail)

	json.NewEncoder(w).Encode(token)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	var authDetails models.User

	errs := json.NewDecoder(r.Body).Decode(&authDetails)
	if errs != nil {
		var err models.Token
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	utils.ParseBody(r, authDetails)

	userdetail, db := models.GetUserById(ID)

	userdetail.Username = authDetails.Username
	userdetail.Role = authDetails.Role
	userdetail.Branch = authDetails.Branch

	if authDetails.Password != "" {
		userdetail.Password, err = GeneratehashPassword(authDetails.Password) //authDetails.Password
		if err != nil {
			log.Fatalln("error in password hash")
		}
	}

	db.Save(&userdetail)

	var token models.Token
	token.Message = "Updated Successfully"

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

	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		//fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err models.Token
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err models.Token
			err = SetError(err, "Your Token has been expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			/* if claims["role"] == "Admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" { */

			r.Header.Set("Role", "user")
			handler.ServeHTTP(w, r)
			return
			//}
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var reserr models.Token
		reserr = SetError(reserr, "Not Authorized")
		json.NewEncoder(w).Encode(err)
	}
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
