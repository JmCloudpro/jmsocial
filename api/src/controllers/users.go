package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser calls the metlhod repository.Create and creates the user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	//create a var using model, and use json Unmarshal do convert the realall in struct.

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepo(db)

	user.ID, err = repository.Create(user)

	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return

	}

	//w.Write([]byte(fmt.Sprintf("User Inserted: %d", userID)))

	answers.JSON(w, http.StatusCreated, user)

}

// SearchUsers - searches all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameornick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repository.NewUsersRepo(db)
	users, err := repo.Search(nameornick)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, users)
}

// SearchUser - search for on specific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)

	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepo(db)
	user, err := repo.SearchID(userID)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, user)
}

// UpdateUser - update an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenuserID, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}

	if tokenuserID != userID {
		answers.Err(w, http.StatusForbidden, errors.New("Not allowed change another user"))
		return

	}

	requestbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		answers.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(requestbody, &user); err != nil {

		answers.Err(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare("update"); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepo(db)
	if err = repo.UpdateUser(userID, user); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}
	answers.JSON(w, http.StatusOK, "User Updated")
}

// DeleteUser - deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userID, err := strconv.ParseUint(param["userID"], 10, 64)

	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	tokenuserID, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println(tokenuserID, userID)
	if tokenuserID != userID {
		answers.Err(w, http.StatusForbidden, errors.New("Not allowed Delete another user"))
		return

	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepo(db)
	if err = repo.DeleteUser(userID); err != nil {

		answers.Err(w, http.StatusInternalServerError, err)
		return

	}

	answers.JSON(w, http.StatusOK, "user deleted")

}

// Followuser allows the loged user follow others users.
func Followuser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	followid, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := strconv.ParseUint(param["userID"], 10, 64)

	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}

	if followid == userID {
		answers.Err(w, http.StatusForbidden, errors.New("Not allowed follow yourself"))
		return

	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepo(db)
	if err = repository.Follow(userID, followid); err != nil {

		answers.Err(w, http.StatusInternalServerError, err)
		return

	}
	answers.JSON(w, http.StatusOK, "User followed")

}

// Unollowuser allows the loged user unfollow others users.
func Unollowuser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	followid, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}

	if followid == userID {
		answers.Err(w, http.StatusForbidden, errors.New("Not unallowed follow yourself"))
		return

	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepo(db)
	if err = repository.Unfollow(userID, followid); err != nil {

		answers.Err(w, http.StatusInternalServerError, err)
		return

	}
	answers.JSON(w, http.StatusOK, "User unfollowed")

}

// SearchFollowers show all followers of an user.
func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repository.NewUsersRepo(db)
	followers, err := repo.SearchFollowers(userID)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, followers)

}

// Update password  - updates user password
func UpdatePawword(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	userID, err := strconv.ParseUint(param["userID"], 10, 64)
	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return

	}

	tokenuserID, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}
	fmt.Println(tokenuserID, userID)
	if tokenuserID != userID {
		answers.Err(w, http.StatusForbidden, errors.New("Not allowed Update a diferent user"))
		return

	}
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password

	if err = json.Unmarshal(bodyRequest, &password); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	//defer db.Close()

	repo := repository.NewUsersRepo(db)

	dbcurrent, err := repo.SearchPasswd(userID)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPasswd(dbcurrent, password.Passwd); err != nil {

		answers.Err(w, http.StatusInternalServerError, errors.New("Password not match"))
		return
	}
	passhash, err := security.Hash(password.Newpasswd)
	if err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = repo.UpdatePasswd(userID, string(passhash)); err != nil {

		answers.Err(w, http.StatusInternalServerError, err)
		return

	}

	answers.JSON(w, http.StatusOK, "Password Updated")

}
