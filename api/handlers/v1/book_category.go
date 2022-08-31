package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rasul07/books_api_gateway/api/models"
	"github.com/rasul07/books_api_gateway/genproto/category"

	"github.com/gin-gonic/gin"
	"github.com/rasul07/books_api_gateway/pkg/util"
)

// Create Category godoc
// @ID create-category
// @Router /v1/category [POST]
// @Summary create category
// @Description Create Category
// @Tags category
// @Accept json
// @Produce json
// @Param category body models.CategoryCreate true "category"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Create(c *gin.Context) {
	var newCategory models.CategoryCreate

	if err := c.BindJSON(&newCategory); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.CategoryService().Create(
		context.Background(),
		&category.CategoryCreate{
			Name: newCategory.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating book category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Category godoc
// @ID get-category
// @Router /v1/category/{category_id} [GET]
// @Summary get category
// @Description Get Category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseModel{data=models.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCategoryById(c *gin.Context) {
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "category id is not valid", errors.New("category id is not valid"))
		return
	}

	resp, err := h.services.CategoryService().GetCategoryById(
		context.Background(),
		&category.GetCategoryByIdRequest{
			BookId: category_id,
		},
	)

	fmt.Println("Resp =====> ", resp)

	if !handleError(h.log, c, err, "error while getting category") {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Get All Categories godoc
// @ID get-all-categories
// @Router /v1/category [GET]
// @Summary get all categories
// @Description Get All Categories
// @Tags category
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} models.ResponseModel{data=category.GetAllCategoriesResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCategories(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	page, err := h.ParseQueryParam(c, "page", "1")
	if err != nil {
		return
	}

	resp, err := h.services.CategoryService().GetCategories(
		context.Background(),
		&category.GetAllCategoriesRequest{
			Limit: int32(limit),
			Page:  int32(page),
		},
	)

	if !handleError(h.log, c, err, "error while getting all categories") {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Update Category godoc
// @ID update-category
// @Router /v1/category [PUT]
// @Summary update category
// @Description Update Category
// @Tags category
// @Accept json
// @Produce json
// @Param book body models.Category true "category"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		updateCategory models.Category
	)

	if err := c.BindJSON(&updateCategory); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.CategoryService().Update(
		context.Background(),
		&category.Category{
			ID:   updateCategory.ID,
			Name: updateCategory.Name,
		},
	)

	if !handleError(h.log, c, err, "error while updating category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Category godoc
// @ID delete-category
// @Router /v1/category/{category_id} [DELETE]
// @Summary delete category
// @Description Delete category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "Category ID"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("category_id")

	resp, err := h.services.CategoryService().Delete(
		context.Background(),
		&category.Category{
			ID: categoryId,
		},
	)

	if !handleError(h.log, c, err, "error while deleting category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}
