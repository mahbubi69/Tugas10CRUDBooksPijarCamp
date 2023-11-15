package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"tugas-10-crud-books-pijar-camp/models"
	"tugas-10-crud-books-pijar-camp/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) InitializedConnectToDb(DbDriver, DbUser, DbPassword, DbHost, DbPort, DbName string) {
	var err error

	if DbDriver == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)

		s.DB, err = gorm.Open(DbDriver, dsn)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Printf("not connected to the %v database", DbDriver)
	}

	s.DB.AutoMigrate(
		models.Produk{},
	)

	gin.SetMode(gin.ReleaseMode)
	s.Router = gin.Default()
	s.InitializedRoutes()

}

func (s *Server) RunServer(addr string) {
	fmt.Println("Listen Of Port Server : " + os.Getenv("PORT"))
	handler := s.Router
	log.Fatal(http.ListenAndServe(addr, handler))
}

// status on server
func (s *Server) StatusServer(c *gin.Context) {
	response.GenericJsonResponse(c, http.StatusOK, "server produk aktif 🗿🗿")
}
