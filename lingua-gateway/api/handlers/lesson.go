package handlers

import (
	"net/http"
	"strconv"

	pb "gateway-service/genprotos"

	"github.com/gin-gonic/gin"
)

// LessonCreate handles the creation of a new lesson.
// @Summary Create Lesson
// @Description Creates a new lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param lesson body pb.LessonCReqForSwagger true "Lesson data"
// @Success 200 {object} string "Lesson created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /lesson [post]
func (h *HTTPHandler) LessonCreate(c *gin.Context) {
	var req pb.LessonCReqGRes
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	_, err := h.Lesson.CreateLesson(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Lesson created!!!")
}

// LessonGet handles getting a lesson by it's UUID.
// @Summary Get lesson
// @Description Get a lesson by it's UUID
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "Leesson ID"
// @Success 200 {object} pb.LessonCReqGRes
// @Failure 400 {object} string "Invalid lesson ID"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /lesson/{id} [get]
func (h *HTTPHandler) LessonGet(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	res, err := h.Lesson.GetLessonByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get lesson", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// LessonUpdate handles updating an existing lesson.
// @Summary Update lesson
// @Description Update an existing lesson
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Param lesson body pb.LessonCReqForSwagger true "Updated lesson data"
// @Success 200 {object} string "Lesson updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Lesson not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /lesson/{id} [put]
func (h *HTTPHandler) LessonUpdate(c *gin.Context) {
	id := c.Param("id")
	var req pb.LessonUReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	_, err := h.Lesson.UpdateLesson(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update lesson ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Lesson updated!!!")
}

// LessonDelete handles deleting a lesson by it's UUID.
// @Summary Delete lesson
// @Description Delete a lesson by it's UUID
// @Tags Lesson
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} string "Lesson deleted"
// @Failure 400 {object} string "Invalid lesson  ID"
// @Failure 404 {object} string "Lesson  not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /lesson/{id} [delete]
func (h *HTTPHandler) LessonDelete(c *gin.Context) {
	id := &pb.ByID{Id: c.Param("id")}
	_, err := h.Lesson.DeleteLesson(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete lesson ", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Lesson deleted!!!")
}

// LessonGetAll handles getting all lessons.
// @Summary Get all lessons
// @Description Get all lessons
// @Tags Lesson
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param lang_1 query string false "lang_1"
// @Param lang_2 query string false "lang_2"
// @Param level query string false "level"
// @Param order query integer false "order"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} pb.LessonGARes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /lessons [get]
func (h *HTTPHandler) LessonGetAll(c *gin.Context) {
	var limit, offset, order int
	var err error
	name := c.Query("name")
	lang1 := c.Query("lang_1")
	lang2 := c.Query("lang_2")
	level := c.Query("level")

	orderStr := c.Query("order")
	if orderStr == "" {
		order = 0
	} else {
		order, err = strconv.Atoi(orderStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter"})
			return
		}
	}

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

	res, err := h.Lesson.GetAllLessons(c, &pb.LessonGAReq{
		Name:   name,
		Lang_1: lang1,
		Lang_2: lang2,
		Level:  level,
		Order:  int32(order),
		Pagination: &pb.Pagination{
			Limit:  int64(limit),
			Offset: int64(offset),
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get lessons", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
