package controllers

import "github.com/KnockOutEZ/rest-api-portfolio/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET","OPTIONS")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST","OPTIONS")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST","OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUser)).Methods("PUT","OPTIONS")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE","OPTIONS")

	//Skills routes
	s.Router.HandleFunc("/skills", middlewares.SetMiddlewareJSON(s.CreateSkill)).Methods("POST")
	s.Router.HandleFunc("/myskills", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMySkills))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareJSON(s.GetSkill)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myskills", middlewares.SetMiddlewareJSON(s.GoGetAllSkills)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myskills/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneSkills)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSkill))).Methods("PUT")
	s.Router.HandleFunc("/skills/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSkill)).Methods("DELETE")

	//Projects routes
	s.Router.HandleFunc("/projects", middlewares.SetMiddlewareJSON(s.CreateProject)).Methods("POST")
	s.Router.HandleFunc("/myprojects", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMyProjects))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/projects/{id}", middlewares.SetMiddlewareJSON(s.GetProject)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myprojects", middlewares.SetMiddlewareJSON(s.GoGetAllProjects)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myprojects/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneProjects)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/projects/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProject))).Methods("PUT")
	s.Router.HandleFunc("/projects/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProject)).Methods("DELETE")

	//SkillAreas routes
	s.Router.HandleFunc("/skillareas", middlewares.SetMiddlewareJSON(s.CreateSkillArea)).Methods("POST")
	s.Router.HandleFunc("/myskillareas", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMySkillAreas))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareJSON(s.GetSkillAreas)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myskillareas", middlewares.SetMiddlewareJSON(s.GoGetAllSkillAreas)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myskillareas/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneSkillAreas)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSkillArea))).Methods("PUT")
	s.Router.HandleFunc("/skillareas/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSkillArea)).Methods("DELETE")

	//Education routes
	s.Router.HandleFunc("/education", middlewares.SetMiddlewareJSON(s.CreateEducation)).Methods("POST")
	s.Router.HandleFunc("/myeducation", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMyEducations))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/education/{id}", middlewares.SetMiddlewareJSON(s.GetEducation)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myeducation", middlewares.SetMiddlewareJSON(s.GoGetAllEducations)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myeducation/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneEducations)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/education/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateEducation))).Methods("PUT")
	s.Router.HandleFunc("/education/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteEducation)).Methods("DELETE")

	//ProfessionalExp routes
	s.Router.HandleFunc("/professional", middlewares.SetMiddlewareJSON(s.CreateProfessionalExp)).Methods("POST")
	s.Router.HandleFunc("/myprofessional", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMyProfessionalExps))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/professional/{id}", middlewares.SetMiddlewareJSON(s.GetProfessionalExp)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myprofessional", middlewares.SetMiddlewareJSON(s.GoGetAllProfessionalExps)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myprofessional/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneProfessionalExps)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/professional/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProfessionalExp))).Methods("PUT")
	s.Router.HandleFunc("/professional/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProfessionalExp)).Methods("DELETE")

	//Hobbys routes
	s.Router.HandleFunc("/hobbies", middlewares.SetMiddlewareJSON(s.CreateHobby)).Methods("POST")
	s.Router.HandleFunc("/myhobbies", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMyHobbys))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/hobbies/{id}", middlewares.SetMiddlewareJSON(s.GetHobby)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/myhobbies", middlewares.SetMiddlewareJSON(s.GoGetAllHobbys)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/myhobbies/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneHobbys)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/hobbies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateHobby))).Methods("PUT")
	s.Router.HandleFunc("/hobbies/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteHobby)).Methods("DELETE")

	//Socials routes
	s.Router.HandleFunc("/socials", middlewares.SetMiddlewareJSON(s.CreateSocials)).Methods("POST")
	s.Router.HandleFunc("/mysocials", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetMySocialss))).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/socials/{id}", middlewares.SetMiddlewareJSON(s.GetSocials)).Methods("GET","OPTIONS")
	//public route start
	s.Router.HandleFunc("/{key}/mysocials", middlewares.SetMiddlewareJSON(s.GoGetAllSocialss)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/{key}/mysocials/{id}", middlewares.SetMiddlewareJSON(s.GoGetOneSocialss)).Methods("GET","OPTIONS")
	//public route end
	s.Router.HandleFunc("/socials/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateSocials))).Methods("PUT")
	s.Router.HandleFunc("/socials/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteSocials)).Methods("DELETE")

	//routes for admin
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/skills", middlewares.SetMiddlewareJSON(s.GetSkills)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/professional", middlewares.SetMiddlewareJSON(s.GetProfessionalExp)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/projects", middlewares.SetMiddlewareJSON(s.GetProjects)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/education", middlewares.SetMiddlewareJSON(s.GetEducations)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/hobbies", middlewares.SetMiddlewareJSON(s.GetHobbys)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/skillareas", middlewares.SetMiddlewareJSON(s.GetSkillAreas)).Methods("GET","OPTIONS")
	s.Router.HandleFunc("/socials", middlewares.SetMiddlewareJSON(s.GetSocials)).Methods("GET","OPTIONS")
}
