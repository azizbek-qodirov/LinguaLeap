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
	protected.Use(middleware.IsAdminMiddleware())

	protected.GET("mydata", h.GetMyData)

	lesson := protected.Group("/lesson")
	{
		lesson.POST("/", h.LessonCreate)
		lesson.GET("/:id", h.LessonGet)
		lesson.PUT("/:id", h.LessonUpdate)
		lesson.DELETE("/:id", h.LessonDelete)
		protected.GET("/lessons", h.LessonGetAll)
	}

	exercise := protected.Group("/exercise")
	{
		exercise.POST("/", h.ExerciseCreate)
		exercise.GET("/:id", h.ExerciseGet)
		exercise.PUT("/:id", h.ExerciseUpdate)
		exercise.DELETE("/:id", h.ExerciseDelete)
		protected.GET("/exercises", h.ExerciseGetAll)
	}

	return router
}
