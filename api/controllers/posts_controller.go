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

func (server *Server) CreateSkill(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill := models.Skill{}
	err = json.Unmarshal(body, &skill)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill.Prepare()
	err = skill.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != skill.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	skillCreated, err := skill.SaveSkill(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, skillCreated.ID))
	responses.JSON(w, http.StatusCreated, skillCreated)
}

func (server *Server) GetSkills(w http.ResponseWriter, r *http.Request) {

	skill := models.Skill{}

	skills, err := skill.FindAllSkills(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skills)
}

func (server *Server) GetMySkills(w http.ResponseWriter, r *http.Request) {

	skill := models.Skill{}
	uid, err := auth.ExtractTokenID(r)
	skills, err := skill.FindAllMySkills(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skills)
}

func (server *Server) GetSkill(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	skill := models.Skill{}

	skillReceived, err := skill.FindSkillByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skillReceived)
}

func (server *Server) UpdateSkill(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the skill id is valid
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

	// Check if the skill exist
	skill := models.Skill{}
	err = server.DB.Debug().Model(models.Skill{}).Where("id = ?", pid).Take(&skill).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("skill not found"))
		return
	}

	// If a user attempt to update a skill not belonging to him
	if uid != skill.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data skilled
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	skillUpdate := models.Skill{}
	err = json.Unmarshal(body, &skillUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != skillUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	skillUpdate.Prepare()
	err = skillUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	skillUpdate.ID = skill.ID //this is important to tell the model the skill id to update, the other update field are set above

	skillUpdated, err := skillUpdate.UpdateASkill(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, skillUpdated)
}

func (server *Server) DeleteSkill(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid skill id given to us?
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

	// Check if the skill exist
	skill := models.Skill{}
	err = server.DB.Debug().Model(models.Skill{}).Where("id = ?", pid).Take(&skill).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this skill?
	if uid != skill.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = skill.DeleteASkill(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
