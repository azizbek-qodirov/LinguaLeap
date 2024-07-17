package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gateway-service/api/docs"
	"gateway-service/api/handlers"
	"gateway-service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connL, ConnP *grpc.ClientConn) *gin.Engine {
	h := handlers.NewHandler(connL, ConnP)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.Use(middleware.IsUserMiddleware())

	router.GET("mydata", h.GetMyData)

	lesson := protected.Group("/lesson")
	{
		lesson.GET("/:id", h.LessonGet)
		protected.GET("/lessons", h.LessonGetAll)
	}

	exercise := protected.Group("/exercise")
	{
		exercise.GET("/:id", h.ExerciseGet)
		protected.GET("/exercises", h.ExerciseGetAll)
	}

	vocabulary := protected.Group("/vocabulary")
	{
		vocabulary.POST("/:id", h.AddToVocabulary)
		vocabulary.DELETE("/:id", h.DeleteFromVocabulary)
		protected.GET("/vocabularies", h.GetVocabularies)
	}

	return router
}
