package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/rasul07/books_api_gateway/api/models"
	"github.com/rasul07/books_api_gateway/genproto/book"

	"github.com/gin-gonic/gin"
	"github.com/rasul07/books_api_gateway/pkg/util"
)

// Create Book godoc
// @ID create-book
// @Router /v1/book [POST]
// @Summary create book
// @Description Create Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.BookCreate true "book"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBook(c *gin.Context) {
	var newBook models.BookCreate

	if err := c.BindJSON(&newBook); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.BookService().Create(
		context.Background(),
		&book.BookCreate{
			Name:        newBook.Name,
			Author:      newBook.Author,
			Category:    newBook.Category,
			Description: newBook.Description,
			Pages:       newBook.Pages,
			Year:        newBook.Year,
		},
	)

	if !handleError(h.log, c, err, "error while creating book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Book godoc
// @ID get-book
// @Router /v1/book/{book_id} [GET]
// @Summary get book
// @Description Get Book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} models.ResponseModel{data=book.GetAllBooksResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBookById(c *gin.Context) {
	book_id := c.Param("book_id")

	if !util.IsValidUUID(book_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "book id is not valid", errors.New("book id is not valid"))
		return
	}

	resp, err := h.services.BookService().GetBookById(
		context.Background(),
		&book.GetBookByIdRequest{
			BookId: book_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Get All Books godoc
// @ID get-all-books
// @Router /v1/book [GET]
// @Summary get all books
// @Description Get All Books
// @Tags book
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} models.ResponseModel{data=models.BookList} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBooks(c *gin.Context) {

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	page, err := h.ParseQueryParam(c, "page", "1")
	if err != nil {
		return
	}

	resp, err := h.services.BookService().GetBooks(
		context.Background(),
		&book.GetAllBooksRequest{
			Limit: int32(limit),
			Page:  int32(page),
		},
	)

	if !handleError(h.log, c, err, "error while getting all books") {
		return
	}

	c.JSON(http.StatusOK, resp)

}

// Update Book godoc
// @ID update-book
// @Router /v1/book [PUT]
// @Summary update book
// @Description Update Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.Book true "book"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var updateBook models.Book

	if err := c.BindJSON(&updateBook); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.BookService().Update(
		context.Background(),
		&book.Book{
			ID:          updateBook.ID,
			Name:        updateBook.Name,
			Author:      updateBook.Author,
			Category:    updateBook.Category,
			Description: updateBook.Description,
			Pages:       updateBook.Pages,
			Year:        updateBook.Year,
		},
	)

	if !handleError(h.log, c, err, "error while updating book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Book godoc
// @ID delete-book
// @Router /v1/book/{book_id} [DELETE]
// @Summary delete book
// @Description Delete book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "Book ID"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Delete(c *gin.Context) {

	bookId := c.Param("book_id")

	resp, err := h.services.BookService().Delete(
		context.Background(),
		&book.Book{
			ID: bookId,
		},
	)

	if !handleError(h.log, c, err, "error while deleting book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}
