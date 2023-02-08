package main

import "github.com/gin-gonic/gin"

type Server struct {
	Engine *gin.Engine
}

// TODO: Set up endpoints
// TODO: This should eventually be set up to connect to a database and maybe use custom logging
func (s *Server) Init(dbUrl string) {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello world"})
	})
	s.Engine = engine
}

func (s *Server) Run(port string) error {
	return s.Engine.Run("localhost:" + port)
}
func NewServer(dbUrl string) Server {
	var s Server
	s.Init(dbUrl)
	return s
}
