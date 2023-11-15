package controller

import "os"

func (s *Server) InitializedRoutes() {
	v1 := s.Router.Group(os.Getenv("BASE_URL"))
	{
		v1.GET("/status", s.StatusServer)
	}
}
