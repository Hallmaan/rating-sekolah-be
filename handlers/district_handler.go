package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
	"strconv"
)

type DistrictHandler struct {
	AUsecase domains.DistrictUseCase
}

func NewDistrictHandler(us domains.DistrictUseCase) *DistrictHandler {
	return &DistrictHandler{us}
}

func (s *DistrictHandler) FetchDistrict(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	ctx := c.Request.Context()

	fmt.Println(limit, "limit data")

	district, err := s.AUsecase.Fetch(ctx, int64(limit), 0)
	if err != nil {
		response := helpers.APIResponse("Error to get school district", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of school district", http.StatusOK, "success", district)
	c.JSON(http.StatusOK, response)
}

func (s *DistrictHandler) GetDistrictById(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	province, err := s.AUsecase.GetByID(ctx, id)
	if err != nil {
		response := helpers.APIResponse("Error to get school district", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("List of school district", http.StatusOK, "success", province)
	c.JSON(http.StatusOK, response)
}