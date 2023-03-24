package interfaces

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/entity"
	"github.com/holizDV/dts-go-chapter2-challenge2/app/domain/service"
	"github.com/holizDV/dts-go-chapter2-challenge2/pkg/helper"
)

type BookController interface {
	CreateAt(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
}

type controller struct {
	service service.BookService
}

func NewBookController(service service.BookService) *controller {
	return &controller{service}
}

func (h *controller) CreateAt(c *gin.Context) {
	var bookRequest entity.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorArr := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field %v, condition: %v", strings.ToLower(e.Field()), e.ActualTag())
			errorArr = append(errorArr, errorMsg)
		}

		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Errors:  errorArr,
		})
		return
	}
	book, err := h.service.CreateAt(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Errors:  err,
		})
		return
	}

	bookResponse := convertToCreatedResponse(book)
	c.JSON(http.StatusOK, entity.HttpResponse{
		Code:    http.StatusOK,
		Success: true,
		Data:    bookResponse,
	})
}

func (h *controller) Update(c *gin.Context) {
	var bookRequest entity.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorArr := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field %v, condition: %v", strings.ToLower(e.Field()), e.ActualTag())
			errorArr = append(errorArr, errorMsg)
		}

		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Errors:  errorArr,
		})
		return
	}
	bookId := c.Params.ByName("bookID")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	book, err := h.service.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Errors:  err,
		})
		return
	}

	bookResponse := convertToResponse(book)
	c.JSON(http.StatusOK, entity.HttpResponse{
		Code:    http.StatusOK,
		Success: true,
		Data:    bookResponse,
	})
}

func (h *controller) Delete(c *gin.Context) {

	bookId := c.Params.ByName("bookID")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	book, err := h.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Errors:  err,
			Code:    http.StatusBadRequest,
		})
		return
	}
	bookResponse := convertToDeleteResponse(book)
	c.JSON(http.StatusOK, entity.HttpResponse{
		Code:    http.StatusOK,
		Success: true,
		Data:    bookResponse,
	})
}

func (h *controller) FindAll(c *gin.Context) {
	var bookResponse []entity.BookResponse

	books, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Data:    bookResponse,
		})
		return
	}

	for _, book := range books {
		response := convertToResponse(book)
		bookResponse = append(bookResponse, response)
	}

	c.JSON(http.StatusOK, entity.HttpResponse{
		Code:    http.StatusOK,
		Success: true,
		Data:    bookResponse,
	})

}

func (h *controller) FindByID(c *gin.Context) {
	var bookResponse entity.BookResponse

	bookId := c.Params.ByName("bookID")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	book, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, entity.HttpResponse{
			Success: false,
			Code:    http.StatusBadRequest,
			Data:    bookResponse,
		})
		return
	}

	c.JSON(http.StatusOK, entity.HttpResponse{
		Code:    http.StatusOK,
		Success: true,
		Data:    book,
	})

}

func convertToResponse(b entity.Book) entity.BookResponse {
	return entity.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}
}

func convertToCreatedResponse(b entity.Book) entity.BookCreatedResponse {
	return entity.BookCreatedResponse{
		Title:       b.Title,
		Author:      b.Author,
		Description: b.Description,
	}
}

func convertToDeleteResponse(b entity.Book) entity.BookDeleteResponse {
	return entity.BookDeleteResponse{
		Id: b.Id,
	}
}
