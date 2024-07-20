package handlers

import (
	"net/http"

	pb "gateway-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// StartTest godoc
// @Summary Start a test
// @Description Start a test by checking the answers of the provided quiz requests against the exercises of the specified lesson
// @Tags quiz
// @Accept json
// @Produce json
// @Param YourAnswers body pb.TestCheckReqForSwagger true "Test Check Request"
// @Success 200 {object} pb.TestResultRes
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /start-test [post]
func (h *HTTPHandler) StartTest(c *gin.Context) {
	var req pb.TestCheckReq
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = claims.(jwt.MapClaims)["user_id"].(string)

	res, err := h.Quiz.StartTest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
