package config

import (
	"github.com/gin-gonic/gin"
	"github.com/holizDV/dts-go-chapter2-challenge2/app/interfaces"
)

func NewRouter(h interfaces.BookController) *gin.Engine {

	router := gin.Default()
	serviceName := "/product"
	bookPoint := "/books"

	v1 := router.Group(serviceName + "/v1")

	v1.GET(bookPoint, h.FindAll)
	v1.GET(bookPoint+"/:bookID", h.FindByID)
	v1.POST(bookPoint, h.CreateAt)
	v1.PUT(bookPoint+"/:bookID", h.Update)
	v1.DELETE(bookPoint+"/:bookID", h.Delete)

	return router
}
