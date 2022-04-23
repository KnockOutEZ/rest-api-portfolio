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

func (server *Server) CreateProject(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	project := models.Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	project.Prepare()
	err = project.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != project.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	projectCreated, err := project.SaveProject(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, projectCreated.ID))
	responses.JSON(w, http.StatusCreated, projectCreated)
}

func (server *Server) GoGetAllProjects(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	project := models.Project{}
	projects, err := project.GoFindAllMyProjects(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, projects)
}

func (server *Server) GoGetOneProjects(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)
	uid, err := strconv.ParseUint(key["key"], 10, 64)
	pid, err := strconv.ParseUint(key["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	project := models.Project{}

	projectReceived, err := project.GoFindProjectByID(server.DB, pid,uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, projectReceived)
}

func (server *Server) GetProjects(w http.ResponseWriter, r *http.Request) {

	project := models.Project{}

	projects, err := project.FindAllProjects(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, projects)
}

func (server *Server) GetMyProjects(w http.ResponseWriter, r *http.Request) {
	project := models.Project{}
	uid, err := auth.ExtractTokenID(r)
	projects, err := project.FindAllMyProjects(server.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, projects)
}



func (server *Server) GetProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	project := models.Project{}

	projectReceived, err := project.FindProjectByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, projectReceived)
}

func (server *Server) UpdateProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the project id is valid
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

	// Check if the project exist
	project := models.Project{}
	err = server.DB.Debug().Model(models.Project{}).Where("id = ?", pid).Take(&project).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("project not found"))
		return
	}

	// If a user attempt to update a project not belonging to him
	if uid != project.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data projected
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	projectUpdate := models.Project{}
	err = json.Unmarshal(body, &projectUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != projectUpdate.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	projectUpdate.Prepare()
	err = projectUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	projectUpdate.ID = project.ID //this is important to tell the model the project id to update, the other update field are set above

	projectUpdated, err := projectUpdate.UpdateAProject(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, projectUpdated)
}

func (server *Server) DeleteProject(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid project id given to us?
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

	// Check if the project exist
	project := models.Project{}
	err = server.DB.Debug().Model(models.Project{}).Where("id = ?", pid).Take(&project).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this project?
	if uid != project.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = project.DeleteAProject(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
