package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/base_ms/configs"
	"github.com/gin-gonic/gin"
)

type Server struct {
	eng *gin.Engine

	srvHTTP  *http.Server
	srvHTTPS *http.Server

	pathCertHTTPS string
	pathKeyHTTPS  string
}

func (s *Server) InitRouter(path string, h func(c gin.Context)) {
}

func (s *Server) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.eng.GET(relativePath, handlers...)
}

func (s *Server) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return s.eng.POST(relativePath, handlers...)
}

func (s *Server) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.eng.Group(relativePath, handlers...)
}

func New() *Server {
	serverConf := configs.ConfigFromEnv("")

	if len(serverConf.Service.PathCertHTTPS) > 0 {
		gin.DisableConsoleColor()
		f, e := os.Create(serverConf.Service.PathFileLogs)
		if e != nil {
			log.Fatalf("%s", e.Error())

		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	eng := gin.New()
	eng.Use(gin.Logger())
	eng.Use(gin.Recovery())

	return &Server{
		eng: eng,
		srvHTTP: &http.Server{
			Addr:           serverConf.Service.AddressHTTP,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.Service.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.Service.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		srvHTTPS: &http.Server{
			Addr:           serverConf.Service.AddressHTTPS,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.Service.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.Service.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		pathCertHTTPS: serverConf.Service.PathCertHTTPS,
		pathKeyHTTPS:  serverConf.Service.PathKeyHTTPS,
	}
}

func (s *Server) Run() {
	go func() {
		log.Printf("http server listen on: %s", s.srvHTTP.Addr)

		if e := s.srvHTTP.ListenAndServe(); e != nil && !errors.Is(e, http.ErrServerClosed) {
			log.Fatalf("http server nor run, error: %s", e.Error())
		}
	}()

	if len(s.pathCertHTTPS) != 0 && len(s.pathKeyHTTPS) != 0 {
		go func() {
			log.Printf("https server listen on: %s", s.srvHTTPS.Addr)

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
