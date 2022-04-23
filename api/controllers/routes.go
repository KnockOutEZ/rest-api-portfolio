package controllers

import "github.com/KnockOutEZ/rest-api-portfolio/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Skills routes
	s.Router.HandleFunc("/skills", middlewares.SetMiddlewareJSON(s.CreateSkill)).Methods("POST")
	s.Router.HandleFunc("/myskills", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMySkills))).Methods("GET")
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareJSON(s.GetSkill)).Methods("GET")
	//public route start
	s.Router.HandleFunc("/{key}/myskills", middlewares.SetMiddlewareJSON(s.GoGetAllSkills)).Methods("GET")
	s.Router.HandleFunc("/{key}/myskills/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneSkills)).Methods("GET")
	//public route end
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSkill))).Methods("PUT")
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSkill)).Methods("DELETE")


	//SkillAreas routes
	s.Router.HandleFunc("/skillareas", middlewares.SetMiddlewareJSON(s.CreateSkillArea)).Methods("POST")
	s.Router.HandleFunc("/myskillareas", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMySkillAreas))).Methods("GET")
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareJSON(s.GetSkillAreas)).Methods("GET")
	//public route start
	s.Router.HandleFunc("/{key}/myskillareas", middlewares.SetMiddlewareJSON(s.GoGetAllSkillAreas)).Methods("GET")
	s.Router.HandleFunc("/{key}/myskillareas/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneSkillAreas)).Methods("GET")
	//public route end
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSkillArea))).Methods("PUT")
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSkillArea)).Methods("DELETE")

	//routes for admin
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/skills", middlewares.SetMiddlewareJSON(s.GetSkills)).Methods("GET")
	s.Router.HandleFunc("/skillareas", middlewares.SetMiddlewareJSON(s.GetSkillAreas)).Methods("GET")
}
