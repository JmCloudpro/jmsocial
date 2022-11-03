package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CreatePublication Creates a publication

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Err(w, http.StatusUnauthorized, err)
		return
	}
	requestbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		answers.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publication
	if err := json.Unmarshal(requestbody, &publication); err != nil {
		answers.Err(w, http.StatusBadRequest, err)
		return
	}
	publication.CreatorID = userId

	db, err := db.Conn()
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewPubliRepo(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		answers.Err(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusCreated, publication)
}

// SearchPublications Searches all publicarions
func SearchPublications(w http.ResponseWriter, r *http.Request) {

	SearchAll

}

// SearchPublicarion searches for a specific publication
func SearchPublicarion(w http.ResponseWriter, r *http.Request) {

}

// UpdatePublication Updates a publicarion
func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

// DeletePublication deletes a publicarion
func DeletePublication(w http.ResponseWriter, r *http.Request) {

}
