package controller

import "os"

func (s *Server) InitializedRoutes() {
	v1 := s.Router.Group(os.Getenv("BASE_URL"))
	{
		v1.GET("/status", s.StatusServer)

		// produk
		v1.POST("/produk", s.CreatedProdukController)
		v1.GET("/produk", s.GetProdukController)
		v1.PUT("/produk/:id", s.UpdateProdukController)
		v1.DELETE("/produk/:id", s.DeleteProdukController)

	}
}
