package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	pb "gateway-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type VocabularyUpdateMessage struct {
	ExerciseID string `json:"exercise_id"`
	Action     string `json:"action"`
}

func (h *HTTPHandler) publishToRabbitMQ(msg VocabularyUpdateMessage) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = h.RabbitMQ.Publish(
		"progress_updates", // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		},
	)
	return err
}

// AddToVocabulary handles adding an exercise to the vocabulary.
// @Summary Add to Vocabulary
// @Description Adds an exercise to the vocabulary
// @Tags Vocabulary
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Success 200 {object} string "Exercise added to vocabulary"
// @Failure 400 {object} string "Invalid exercise ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vocabulary/{id} [post]
func (h *HTTPHandler) AddToVocabulary(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	msg := VocabularyUpdateMessage{
		ExerciseID: id.Id,
		Action:     "add",
	}
	if err := h.publishToRabbitMQ(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't publish message to RabbitMQ", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Exercise added to vocabulary!!!")
}

// DeleteFromVocabulary handles deleting an exercise from the vocabulary.
// @Summary Delete from Vocabulary
// @Description Deletes an exercise from the vocabulary
// @Tags Vocabulary
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Success 204 {object} string "Exercise deleted from vocabulary"
// @Failure 400 {object} string "Invalid exercise ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vocabulary/{id} [delete]
func (h *HTTPHandler) DeleteFromVocabulary(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	_, err := h.Vocabulary.DeleteFromVocabulary(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete exercise from vocabulary", "details": err.Error()})
		return
	}

	msg := VocabularyUpdateMessage{
		ExerciseID: id.Id,
		Action:     "delete",
	}
	if err := h.publishToRabbitMQ(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't publish message to RabbitMQ", "details": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "Exercise deleted from vocabulary!!!")
}

// GetVocabularies handles getting vocabularies.
// @Summary Get Vocabularies
// @Description Gets vocabularies
// @Tags Vocabulary
// @Accept json
// @Produce json
// @Param lesson_id query string false "Lesson ID"
// @Param type query string false "Exercise type"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.VocabulariesGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vocabularies [get]
func (h *HTTPHandler) GetVocabularies(c *gin.Context) {
	var limit, offset int
	var err error

	lessonID := c.Query("lesson_id")
	exerciseType := c.Query("type")
	limitStr := c.Query("limit")
	if limitStr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	res, err := h.Vocabulary.GetVocabularies(c, &pb.VocabulariesGAReq{
		LessonId: lessonID,
		Type:     exerciseType,
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	fmt.Println(lessonID, exerciseType, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get vocabularies", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
