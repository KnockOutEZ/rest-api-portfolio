package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/KnockOutEZ/rest-api-portfolio/api/auth"
	"github.com/KnockOutEZ/rest-api-portfolio/api/models"
	"github.com/KnockOutEZ/rest-api-portfolio/api/responses"
	"github.com/KnockOutEZ/rest-api-portfolio/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) CreateEducation(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	education := models.Education{}
	err = json.Unmarshal(body, &education)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	education.Prepare()
	err = education.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != education.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	educationCreated, err := education.SaveEducation(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, educationCreated.ID))
	responses.JSON(w, http.StatusCreated, educationCreated)
}

func (server *Server) GoGetAllEducations(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	education := models.Education{}
	educations, err := education.GoFindAllMyEducations(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, educations)
}

func (server *Server) GoGetOneEducations(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	pid, err := strconv.ParseUint(key["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	education := models.Education{}

	educationReceived, err := education.GoFindEducationByID(server.DB, pid,uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, educationReceived)
}

func (server *Server) GetEducations(w http.ResponseWriter, r *http.Request) {

	education := models.Education{}

	educations, err := education.FindAllEducations(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, educations)
}

func (server *Server) GetMyEducations(w http.ResponseWriter, r *http.Request) {
	education := models.Education{}
	uid, err := auth.ExtractTokenID(r)
	educations, err := education.FindAllMyEducations(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, educations)
}



func (server *Server) GetEducation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	education := models.Education{}

	educationReceived, err := education.FindEducationByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, educationReceived)
}

func (server *Server) UpdateEducation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the education id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the education exist
	education := models.Education{}
	err = server.DB.Debug().Model(models.Education{}).Where("id = ?", pid).Take(&education).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("education not found"))
		return
	}

	// If a user attempt to update a education not belonging to him
	if uid != education.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data educationed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	educationUpdate := models.Education{}
	err = json.Unmarshal(body, &educationUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != educationUpdate.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	educationUpdate.Prepare()
	err = educationUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	educationUpdate.ID = education.ID //this is important to tell the model the education id to update, the other update field are set above

	educationUpdated, err := educationUpdate.UpdateAEducation(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, educationUpdated)
}

func (server *Server) DeleteEducation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid education id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the education exist
	education := models.Education{}
	err = server.DB.Debug().Model(models.Education{}).Where("id = ?", pid).Take(&education).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this education?
	if uid != education.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = education.DeleteAEducation(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
