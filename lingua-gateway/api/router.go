package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gateway-service/api/docs"
	"gateway-service/api/handlers"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connL, ConnP *grpc.ClientConn) *gin.Engine {
	h := handlers.NewHandler(connL, ConnP)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// protected := router.Group("/", middleware.JWTMiddleware())

	lesson := router.Group("/lesson")
	{
		lesson.POST("/", h.LessonCreate)
		lesson.GET("/:id", h.LessonGet)
		lesson.PUT("/:id", h.LessonUpdate)
		lesson.DELETE("/:id", h.LessonDelete)
		router.GET("/lessons", h.LessonGetAll)
	}

	exercise := router.Group("/exercise")
	{
		exercise.POST("/", h.ExerciseCreate)
		exercise.GET("/:id", h.ExerciseGet)
		exercise.PUT("/:id", h.ExerciseUpdate)
		exercise.DELETE("/:id", h.ExerciseDelete)
		router.GET("/exercises", h.ExerciseGetAll)
	}

	vocabulary := router.Group("/vocabulary")
	{
		vocabulary.POST("/:id", h.AddToVocabulary)
		vocabulary.DELETE("/:id", h.DeleteFromVocabulary)
		router.GET("/vocabularies", h.GetVocabularies)
	}

	return router
}
