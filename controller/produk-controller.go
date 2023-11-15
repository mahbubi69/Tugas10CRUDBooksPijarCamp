package controller

import (
	"net/http"
	"strconv"
	"tugas-10-crud-books-pijar-camp/models"
	"tugas-10-crud-books-pijar-camp/response"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreatedProdukController(c *gin.Context) {
	produk := models.Produk{}

	if err := c.BindJSON(&produk); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	errCreatedProduk := produk.CreatedProduk(s.DB)

	if errCreatedProduk != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, errCreatedProduk)
		return
	}

	response.GenericJsonResponse(c, http.StatusOK, "Berhasil")

}

func (s *Server) GetProdukController(c *gin.Context) {
	produk := models.Produk{}

	getAllProduk, count, err := produk.GetProduk(s.DB)

	if err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	response.GetJsonResponse(c, uint32(count), http.StatusOK, "Berhasil", getAllProduk)

}

func (s *Server) UpdateProdukController(c *gin.Context) {
	produk := models.Produk{}

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := c.BindJSON(&produk); err != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	errUpdateProduk := produk.EditProduk(s.DB, uint32(id))

	if errUpdateProduk != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, errUpdateProduk)
		return
	}

	response.GenericJsonResponse(c, http.StatusOK, "Berhasil")
}

func (s *Server) DeleteProdukController(c *gin.Context) {
	produk := models.Produk{}

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	errDeleteProduk := produk.DeleteProduk(s.DB, uint32(id))

	if errDeleteProduk != nil {
		response.ErrorResponse(c, http.StatusUnprocessableEntity, errDeleteProduk)
		return
	}

	response.GenericJsonResponse(c, http.StatusOK, "Berhasil")
}
