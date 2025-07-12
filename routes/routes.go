package routes

import (
	"golang_API/handlers"
	"golang_API/middleware"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(r *mux.Router) {
	// Problem routes - require authentication
	r.HandleFunc("/problemEntry/reportProblem", middleware.AuthMiddleware(handlers.ReportProblemHandler)).Methods("POST")
	r.HandleFunc("/problemEntry/solveProblem", middleware.AuthMiddleware(handlers.SolveProblemHandler)).Methods("PUT")
	r.HandleFunc("/problemEntry/problems", middleware.AuthMiddleware(handlers.GetProblemsHandler)).Methods("GET")
	r.HandleFunc("/problemEntry/problem/{id}", middleware.AuthMiddleware(handlers.GetProblemByIDHandler)).Methods("GET")
	r.HandleFunc("/problemEntry/problem/{id}", middleware.AuthMiddleware(handlers.UpdateProblemHandler)).Methods("PUT")
	r.HandleFunc("/problemEntry/problem/{id}", middleware.AuthMiddleware(handlers.DeleteProblemHandler)).Methods("DELETE")
	r.HandleFunc("/problemEntry/problem/{id}/reset-solution", middleware.AuthMiddleware(handlers.ResetSolutionHandler)).Methods("PUT")
	r.HandleFunc("/problemEntry/problem/{id}/update-solution", middleware.AuthMiddleware(handlers.UpdateSolutionHandler)).Methods("PUT")
	r.HandleFunc("/problemEntry/problem/{id}/update-problem", middleware.AuthMiddleware(handlers.UpdateProblemANDSolveProblemHandler)).Methods("PUT")
	r.HandleFunc("/problemEntry/problem/{id}/delete-solution", middleware.AuthMiddleware(handlers.DeleteProblemandSolveProblemHandler)).Methods("DELETE")

	// Branch office routes - require authentication
	r.HandleFunc("/branchEntry/branchOffice", middleware.AuthMiddleware(handlers.AddBranchOfficeHandler)).Methods("POST")
	r.HandleFunc("/branchEntry/branchOffice/{ip_phone}", middleware.AuthMiddleware(handlers.UpdateBranchOfficeHandler)).Methods("PUT")
	r.HandleFunc("/branchEntry/branchOffice/{ip_phone}", middleware.AuthMiddleware(handlers.DeleteBranchOfficeHandler)).Methods("DELETE")
	r.HandleFunc("/branchEntry/branchOffices", middleware.AuthMiddleware(handlers.GetBranchOfficesHandler)).Methods("GET")

	// User routes - require admin access
	r.HandleFunc("/userEntry/user", middleware.AdminMiddleware(handlers.AddUserHandler)).Methods("POST")
	r.HandleFunc("/userEntry/users", middleware.AdminMiddleware(handlers.GetUsersHandler)).Methods("GET")

	// Program routes - require authentication
	r.HandleFunc("/programEntry/program", middleware.AuthMiddleware(handlers.AddProgramHandler)).Methods("POST")
	r.HandleFunc("/programEntry/programs", middleware.AuthMiddleware(handlers.GetProgramsHandler)).Methods("GET")
	r.HandleFunc("/programEntry/program/{id}", middleware.AuthMiddleware(handlers.UpdateProgramHandler)).Methods("PUT")
	r.HandleFunc("/programEntry/program/{id}", middleware.AuthMiddleware(handlers.DeleteProgramHandler)).Methods("DELETE")

	// Delete all data routes - require admin access
	r.HandleFunc("/problemEntry/deleteAllProblems", middleware.AdminMiddleware(handlers.DeleteAllProblemsHandler)).Methods("DELETE")
	r.HandleFunc("/branchEntry/deleteAllBranchOffices", middleware.AdminMiddleware(handlers.DeleteAllBranchOfficesHandler)).Methods("DELETE")
	r.HandleFunc("/programEntry/deleteAllPrograms", middleware.AdminMiddleware(handlers.DeleteAllProgramsHandler)).Methods("DELETE")

	// Dashboard routes - require authentication
	r.HandleFunc("/dashboardEntry/dashboard", middleware.AuthMiddleware(handlers.GetDashboardDataHandler)).Methods("GET")

	// Health check - public access
	r.HandleFunc("/healthEntry/health", handlers.HealthCheckHandler).Methods("GET")
}

// RegisterAuthRoutes registers all authentication-related routes
func RegisterAuthRoutes(r *mux.Router) {
	// Authentication routes - public access
	r.HandleFunc("/authEntry/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/authEntry/registerUser", handlers.RegisterHandler("user")).Methods("POST")
	r.HandleFunc("/authEntry/registerAdmin", middleware.AdminMiddleware(handlers.RegisterHandler("admin"))).Methods("POST")
	r.HandleFunc("/authEntry/updateUser", middleware.AdminMiddleware(handlers.UpdateUserHandler)).Methods("PUT")
	r.HandleFunc("/authEntry/deleteUser", middleware.AdminMiddleware(handlers.DeleteUserHandler)).Methods("DELETE")
	r.HandleFunc("/authEntry/logout", handlers.LogoutHandler).Methods("POST")
}
