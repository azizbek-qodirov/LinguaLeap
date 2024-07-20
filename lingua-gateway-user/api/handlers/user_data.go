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

	id := claims.(jwt.MapClaims)["user_id"].(string)

	data, err := s.UserData.GetUserData(c, &pb.ByID{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get user data", "details": err.Error()})
	}
	user, err := getUserDetails(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't fetch user details", "details": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"user_data": data, "user": user})
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

// @Summary Get Leaderboard
// @Description Get the leaderboard sorted by XP
// @Tags leaderboard
// @Produce json
// @Success 200 {object} []pb.LeadboardRes
// @Failure 500 {object} string
// @Security BearerAuth
// @Router /leaderboard [get]
func (h *HTTPHandler) GetLeadBoard(c *gin.Context) {
	type UserRes struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	type LeadBoardRes struct {
		Level             string
		NativeLang        string
		Xp                int64
		DailyStreak       int32
		PlayedGamesCount  int64
		WinningPercentage float32
		User              *UserRes
	}
	leadboard, err := h.UserData.GetLeadBoard(c, &pb.Void{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ress := []*LeadBoardRes{}
	res := LeadBoardRes{}
	for _, v := range leadboard.Users {
		user, err := getUserDetails(v.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res.Level = v.Level
		res.NativeLang = v.NativeLang
		res.Xp = v.Xp
		res.DailyStreak = v.DailyStreak
		res.PlayedGamesCount = v.PlayedGamesCount
		res.WinningPercentage = v.WinningPercentage
		res.User = &UserRes{
			Username: user.Username,
			Email:    user.Email,
		}
		ress = append(ress, &res)
	}
	c.JSON(http.StatusOK, ress)
}
