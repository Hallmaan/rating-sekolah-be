package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
	"strconv"
)

type ProvinceHandler struct {
	AUsecase domains.ProvinceUseCase
}

func NewProvinceHandler(us domains.ProvinceUseCase) *ProvinceHandler {
	return &ProvinceHandler{us}
}

func (s *ProvinceHandler) FetchProvince(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	ctx := c.Request.Context()

	province, err := s.AUsecase.Fetch(ctx, int64(limit), 0)
	if err != nil {
		response := helpers.APIResponse("Error to get school province", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of school province", http.StatusOK, "success", province)
	c.JSON(http.StatusOK, response)
}

func (s *ProvinceHandler) GetProvinceById(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	province, err := s.AUsecase.GetByID(ctx, id)
	if err != nil {
		response := helpers.APIResponse("Error to get school province", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of school province", http.StatusOK, "success", province)
	c.JSON(http.StatusOK, response)
}