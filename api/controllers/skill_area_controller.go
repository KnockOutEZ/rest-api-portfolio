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

func (server *Server) CreateSkillArea(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill_area := models.SkillArea{}
	err = json.Unmarshal(body, &skill_area)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill_area.Prepare()
	err = skill_area.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != skill_area.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	skill_areaCreated, err := skill_area.SaveSkillArea(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, skill_areaCreated.ID))
	responses.JSON(w, http.StatusCreated, skill_areaCreated)
}

func (server *Server) GoGetAllSkillAreas(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	skill_area := models.SkillArea{}
	skill_areas, err := skill_area.GoFindAllMySkillAreas(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areas)
}

func (server *Server) GoGetOneSkillAreas(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	pid, err := strconv.ParseUint(key["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	skill_area := models.SkillArea{}

	skill_areaReceived, err := skill_area.GoFindSkillAreaByID(server.DB, pid,uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areaReceived)
}

func (server *Server) GetSkillAreas(w http.ResponseWriter, r *http.Request) {

	skill_area := models.SkillArea{}

	skill_areas, err := skill_area.FindAllSkillAreas(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areas)
}

func (server *Server) GetMySkillAreas(w http.ResponseWriter, r *http.Request) {
	skill_area := models.SkillArea{}
	uid, err := auth.ExtractTokenID(r)
	skill_areas, err := skill_area.FindAllMySkillAreas(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areas)
}



func (server *Server) GetSkillArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	skill_area := models.SkillArea{}

	skill_areaReceived, err := skill_area.FindSkillAreaByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areaReceived)
}

func (server *Server) UpdateSkillArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the skill_area id is valid
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

	// Check if the skill_area exist
	skill_area := models.SkillArea{}
	err = server.DB.Debug().Model(models.SkillArea{}).Where("id = ?", pid).Take(&skill_area).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("skill_area not found"))
		return
	}

	// If a user attempt to update a skill_area not belonging to him
	if uid != skill_area.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data skill_areaed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	skill_areaUpdate := models.SkillArea{}
	err = json.Unmarshal(body, &skill_areaUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != skill_areaUpdate.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	skill_areaUpdate.Prepare()
	err = skill_areaUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	skill_areaUpdate.ID = skill_area.ID //this is important to tell the model the skill_area id to update, the other update field are set above

	skill_areaUpdated, err := skill_areaUpdate.UpdateASkillArea(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, skill_areaUpdated)
}

func (server *Server) DeleteSkillArea(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid skill_area id given to us?
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

	// Check if the skill_area exist
	skill_area := models.SkillArea{}
	err = server.DB.Debug().Model(models.SkillArea{}).Where("id = ?", pid).Take(&skill_area).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this skill_area?
	if uid != skill_area.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = skill_area.DeleteASkillArea(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
