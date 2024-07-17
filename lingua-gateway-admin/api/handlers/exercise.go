package handlers

import (
	"net/http"
	"strconv"

	pb "gateway-service/genprotos"

	"github.com/gin-gonic/gin"
)

// ExerciseCreate handles the creation of a new exercise.
// @Summary Create Exercise
// @Description Creates a new exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param exercise body pb.ExerciseCReqUReqForSwagger true "Exercise data"
// @Success 201 {object} string "Exercise created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /exercise [post]
func (h *HTTPHandler) ExerciseCreate(c *gin.Context) {
	var req pb.ExerciseCReqUReqForSwagger
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	_, err := h.Exercise.CreateExercise(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "Exercise created!!!")
}

// ExerciseGet handles getting an exercise by its ID.
// @Summary Get exercise
// @Description Get an exercise by its ID
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Success 200 {object} pb.ExerciseGResUReq
// @Failure 400 {object} string "Invalid exercise ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /exercise/{id} [get]
func (h *HTTPHandler) ExerciseGet(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	res, err := h.Exercise.GetExerciseByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get exercise", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ExerciseUpdate handles updating an existing exercise.
// @Summary Update exercise
// @Description Update an existing exercise
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Param exercise body pb.ExerciseCReqUReqForSwagger true "Updated exercise data"
// @Success 200 {object} string "Exercise updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Exercise not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /exercise/{id} [put]
func (h *HTTPHandler) ExerciseUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pb.ExerciseGResUReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	_, err := h.Exercise.UpdateExercise(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update exercise", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Exercise updated!!!")
}

// ExerciseDelete handles deleting an exercise by its ID.
// @Summary Delete exercise
// @Description Delete an exercise by its ID
// @Tags Exercise
// @Accept json
// @Produce json
// @Param id path string true "Exercise ID"
// @Success 204 {object} string "Exercise deleted"
// @Failure 400 {object} string "Invalid exercise ID"
// @Failure 404 {object} string "Exercise not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /exercise/{id} [delete]
func (h *HTTPHandler) ExerciseDelete(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	_, err := h.Exercise.DeleteExercise(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete exercise", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Exercise deleted!!!")
}

// ExerciseGetAll handles getting all exercises.
// @Summary Get all exercises
// @Description Get all exercises
// @Tags Exercise
// @Accept json
// @Produce json
// @Param lesson_id query string false "Lesson ID"
// @Param type query string false "Exercise type"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.ExerciseGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /exercises [get]
func (h *HTTPHandler) ExerciseGetAll(c *gin.Context) {
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
	res, err := h.Exercise.GetAllExercises(c, &pb.ExerciseGAReq{
		LessonId: lessonID,
		Type:     exerciseType,
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get exercises", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
