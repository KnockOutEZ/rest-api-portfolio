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

func (server *Server) CreateProfessionalExp(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	professional_exp := models.ProfessionalExp{}
	err = json.Unmarshal(body, &professional_exp)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	professional_exp.Prepare()
	err = professional_exp.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != professional_exp.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	professional_expCreated, err := professional_exp.SaveProfessionalExp(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, professional_expCreated.ID))
	responses.JSON(w, http.StatusCreated, professional_expCreated)
}

func (server *Server) GoGetAllProfessionalExps(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	professional_exp := models.ProfessionalExp{}
	professional_exps, err := professional_exp.GoFindAllMyProfessionalExps(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, professional_exps)
}

func (server *Server) GoGetOneProfessionalExps(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	pid, err := strconv.ParseUint(key["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	professional_exp := models.ProfessionalExp{}

	professional_expReceived, err := professional_exp.GoFindProfessionalExpByID(server.DB, pid,uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, professional_expReceived)
}

func (server *Server) GetProfessionalExps(w http.ResponseWriter, r *http.Request) {

	professional_exp := models.ProfessionalExp{}

	professional_exps, err := professional_exp.FindAllProfessionalExps(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, professional_exps)
}

func (server *Server) GetMyProfessionalExps(w http.ResponseWriter, r *http.Request) {
	professional_exp := models.ProfessionalExp{}
	uid, err := auth.ExtractTokenID(r)
	professional_exps, err := professional_exp.FindAllMyProfessionalExps(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, professional_exps)
}



func (server *Server) GetProfessionalExp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	professional_exp := models.ProfessionalExp{}

	professional_expReceived, err := professional_exp.FindProfessionalExpByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, professional_expReceived)
}

func (server *Server) UpdateProfessionalExp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the professional_exp id is valid
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

	// Check if the professional_exp exist
	professional_exp := models.ProfessionalExp{}
	err = server.DB.Debug().Model(models.ProfessionalExp{}).Where("id = ?", pid).Take(&professional_exp).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("professional_exp not found"))
		return
	}

	// If a user attempt to update a professional_exp not belonging to him
	if uid != professional_exp.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data professional_exped
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	professional_expUpdate := models.ProfessionalExp{}
	err = json.Unmarshal(body, &professional_expUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != professional_expUpdate.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	professional_expUpdate.Prepare()
	err = professional_expUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professional_expUpdate.ID = professional_exp.ID //this is important to tell the model the professional_exp id to update, the other update field are set above

	professional_expUpdated, err := professional_expUpdate.UpdateAProfessionalExp(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, professional_expUpdated)
}

func (server *Server) DeleteProfessionalExp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid professional_exp id given to us?
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

	// Check if the professional_exp exist
	professional_exp := models.ProfessionalExp{}
	err = server.DB.Debug().Model(models.ProfessionalExp{}).Where("id = ?", pid).Take(&professional_exp).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this professional_exp?
	if uid != professional_exp.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = professional_exp.DeleteAProfessionalExp(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
