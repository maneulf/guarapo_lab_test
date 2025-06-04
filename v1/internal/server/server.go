package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maneulf/guarapo_lab_test/internal/handlers"
	"github.com/maneulf/guarapo_lab_test/internal/middlewares"
	"github.com/maneulf/guarapo_lab_test/internal/ports"
	"github.com/maneulf/guarapo_lab_test/internal/repositories"
	aprouter "github.com/maneulf/guarapo_lab_test/internal/routers"
	"github.com/maneulf/guarapo_lab_test/internal/services"
)

// Server ...
type Server struct {
	Eng *gin.Engine

	srvHTTP  *http.Server
	srvHTTPS *http.Server

	pathCertHTTPS string
	pathKeyHTTPS  string
}

func New() *Server {

	serverConf, err := ConfigFromEnv()

	if err != nil {
		log.Printf("Could not read enviroment configs, Err: %s", err)
	}

	eng := gin.New()
	eng.Use(gin.Logger())
	eng.Use(gin.Recovery())

	base := &handlers.Base{
		Usernames: map[string]string{},
	}

	//Dependencies injection
	var tasksRepository ports.TasksRepository
	if serverConf.PERSISTENCE_TYPE == "sqlite" {
		tasksRepository = repositories.NewSQLiteRepository("tasks.db")
		log.Println("Executing with sqlite persistence")
	} else if serverConf.PERSISTENCE_TYPE == "inmemory" {
		tasksRepository = repositories.NewMemTasksRepository()
		log.Println("Executing with inmemory persistence")
	} else {
		log.Println("persistence type is unknown, aborting...")
		return nil
	}

	tasksService := services.NewTasksService(tasksRepository)
	handlers := handlers.New(base, tasksService)
	middlewares := middlewares.New(base)

	r := aprouter.Router{
		Eng: eng,
		H:   handlers,
		M:   middlewares,
	}

	r.InitRouters()

	return &Server{
		Eng: eng,
		srvHTTP: &http.Server{
			Addr:           serverConf.AddressHTTP,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		srvHTTPS: &http.Server{
			Addr:           serverConf.AddressHTTPS,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		pathCertHTTPS: serverConf.PathCertHTTPS,
		pathKeyHTTPS:  serverConf.PathKeyHTTPS,
	}

}

// Run ...
func (s *Server) Run() {
	// http
	go func() {
		log.Printf("http server listen on :%s", s.srvHTTP.Addr)

		if e := s.srvHTTP.ListenAndServe(); e != nil && !errors.Is(e, http.ErrServerClosed) {
			log.Fatalf("http server not run, error: %s", e.Error())
		}
	}()
	if len(s.pathCertHTTPS) != 0 && len(s.pathKeyHTTPS) != 0 { // run https only if find path certs
		// https
		go func() {
			log.Printf("https server listen on :%s", s.srvHTTPS.Addr)

			if e := s.srvHTTPS.ListenAndServeTLS(s.pathCertHTTPS, s.pathKeyHTTPS); e != nil && !errors.Is(e, http.ErrServerClosed) {
				log.Fatalf("https server not run, error: %s", e.Error())
			}
		}()
	} else {
		log.Println("https server not run, cert and/or key no find on enviroment")
	}
}

func (s *Server) Shutdown(ctx context.Context) {
	if e := s.srvHTTP.Shutdown(ctx); e != nil {
		log.Printf("http server shutdown, error: %s", e.Error())
	} else {
		log.Println("http server shutdown")
	}
	if e := s.srvHTTPS.Shutdown(ctx); e != nil {
		log.Printf("https server shutdown, error: %s", e.Error())
	} else {
		log.Println("https server shutdown")
	}
}
