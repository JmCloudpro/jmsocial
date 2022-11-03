package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is responsable to authenticate the users
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
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
	userindb, err := repository.SearchEmail(user.Email)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPasswd(userindb.Passwd, user.Passwd); err != nil {
		w.Write([]byte("Use or Password invalid!\n"))
		//answers.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CrateToken(userindb.ID)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return

	}

	w.Write([]byte(token))
	//answers.JSON(w, http.StatusFound, token)

}
