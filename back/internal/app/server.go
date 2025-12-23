package app

import (
	"back/internal/config"
	"back/internal/endpoints/candidate_resume"
	"back/internal/middlewares"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	cfg    *config.Config
	db     *gorm.DB
	router *mux.Router
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	s := &Server{
		cfg:    cfg,
		db:     db,
		router: mux.NewRouter(),
	}
	s.routers()
	return s
}

func (s *Server) routers() {
	protected := s.router.PathPrefix("/api/v1").Subrouter()
	protected.Use(middlewares.AuthMiddleware(s.cfg.SecretToken))

	//Обработка резюме
	candidateResumeRepo := candidate_resume.NewRepository(s.db)
	candidateResumeService := candidate_resume.NewService(candidateResumeRepo, s.cfg)
	candidateResumeHandler := candidate_resume.NewHandler(candidateResumeService)

	protected.HandleFunc("/resume", candidateResumeHandler.CheckCandidate).Methods("POST")
}

func (s *Server) Start() {
	addr := s.cfg.BackPort
	logrus.Info("Server is running on ", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatal(err)
	}
}
