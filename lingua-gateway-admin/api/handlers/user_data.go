package handlers

import (
	"encoding/json"
	"fmt"
	pb "gateway-service/genprotos"
	"gateway-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MyDataResponse struct {
	UserData *pb.UserDataGRes `json:"user_data"`
	User     *models.User     `json:"user"`
}

// GetMyData retrieves user data.
// @Summary Get User Data
// @Description Retrieves user data
// @Tags UserData
// @Accept json
// @Produce json
// @Success 200 {object} MyDataResponse
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /mydata [get]
func (s *HTTPHandler) GetMyData(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id := &pb.ByID{Id: claims.(jwt.MapClaims)["user_id"].(string)}

	data, err := s.UserData.GetUserData(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get user data", "details": err.Error()})
		return
	}
	user, err := getUserDetails(id.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch user details", "details": err.Error()})
		return
	}

	res := MyDataResponse{
		UserData: data,
		User:     user,
	}
	c.JSON(http.StatusOK, res)
}

func getUserDetails(userID string) (*models.User, error) {
	url := fmt.Sprintf("http://localhost:8088/user/%s", userID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user service returned non-200 status: %s", resp.Status)
	}
	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user details: %w", err)
	}
	return &user, nil
}
