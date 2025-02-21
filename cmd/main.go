package main

import (
	"log"
	"net/http"

	_ "sample-project/docs"
	"sample-project/handlers"
	"sample-project/internal/usecases/subject"
	"sample-project/internal/usecases/user"
	"sample-project/pkg/database"
	"sample-project/repositories"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	dsn := "postgresql://postgres:123456@127.0.0.1:5432/go_sample_project"

	db, err := database.NewPostgresDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	subjectRepo := repositories.NewSubjectRepository(db)

	// Initialize use cases
	userUseCase := user.NewUserUseCase(userRepo)
	classroomUseCase := subject.NewSubjectUseCase(subjectRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userUseCase)
	subjectHandler := handlers.NewSubjectHandler(classroomUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/subjects", subjectHandler.CreateSubject).Methods("POST")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
