package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
	"strconv"
)

type SchoolHandler struct {
	AUsecase domains.SchoolUseCase
}

func NewSchoolHandler(us domains.SchoolUseCase) *SchoolHandler {
	return &SchoolHandler{us}
}

func (s *SchoolHandler) FetchSchool(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	ctx := c.Request.Context()

	school, err := s.AUsecase.Fetch(ctx, int64(limit), 0)
	if err != nil {
		response := helpers.APIResponse("Error to get schools", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(school, "lasldlasda")

	response := helpers.APIResponse("List of school", http.StatusOK, "success", school)
	c.JSON(http.StatusOK, response)
}

func (s *SchoolHandler) GetSchoolById(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	school, err := s.AUsecase.GetByID(ctx, id)
	if err != nil {
		response := helpers.APIResponse("Error to get schools", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(school, "lasldlasda")

	response := helpers.APIResponse("List of school", http.StatusOK, "success", school)
	c.JSON(http.StatusOK, response)
}